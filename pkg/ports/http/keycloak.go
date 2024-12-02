package http

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/ceblay/billing-demo/common"
	"github.com/ceblay/billing-demo/exceptions"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
)

type KeycloakUser struct {
	ID            string `json:"sub,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	Name          string `json:"name,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	Username      string `json:"preferred_username,omitempty"`
	Email         string `json:"email,omitempty"`
	Audience      string `json:"audience"`
	EmailVerified bool   `json:"emailVerified,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	Country       string `json:"country"`
	Approved      bool   `json:"approved"`
	OrgID         string `json:"orgId"`
	RealmAccess   struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
}

func (kc *KeycloakUser) String() string {
	jsonBytes, _ := json.Marshal(kc)
	return string(jsonBytes)
}

func GetUserClaims(ctx *fiber.Ctx) *KeycloakUser {
	claims := ctx.Locals("keycloakUser")
	if claims == nil {
		return nil
	}

	return claims.(*KeycloakUser)
}

func GetAccessToken(ctx *fiber.Ctx) string {
	rawAccessToken, ok := ctx.GetReqHeaders()["Authorization"]
	if !ok {
		return ""
	}

	return rawAccessToken[0]
}

func IsAuthorizedJWT(ctx *fiber.Ctx) error {
	var keycloakURL = fmt.Sprintf("%s/realms/%s", os.Getenv("KC.BASE_URL"), os.Getenv("KC.REALM"))
	var clientID = os.Getenv("KC.CLIENT_ID")

	rawAccessToken, ok := ctx.GetReqHeaders()["Authorization"]
	if !ok {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   time.Duration(6000) * time.Second,
		Transport: tr,
	}

	c := oidc.ClientContext(context.Background(), client)
	provider, err := oidc.NewProvider(c, keycloakURL)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	accessToken := strings.Split(rawAccessToken[0], " ")[1]
	idToken, err := verifier.Verify(c, accessToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	user := new(KeycloakUser)
	if err := idToken.Claims(user); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	ctx.Locals("keycloakUser", user)

	return ctx.Next()
}

func IsCustomerOrAgent(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	userClaims := GetUserClaims(ctx)
	authorizedPersonas := []string{"vendor", "agent"}
	audience := strings.ToLower(userClaims.Audience)
	if !slices.Contains(authorizedPersonas, audience) {
		apiResponse.Success = false
		apiResponse.Errors = []string{exceptions.ErrorAudienceForbidden.Error()}
		return ctx.Status(fiber.StatusForbidden).JSON(apiResponse)
	}

	return ctx.Next()
}

func IsDistributor(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	userClaims := GetUserClaims(ctx)
	authorizedPersonas := []string{"distributor"}
	audience := strings.ToLower(userClaims.Audience)
	log.Println("User audience ::::: ", audience)
	if !slices.Contains(authorizedPersonas, audience) {
		apiResponse.Success = false
		apiResponse.Errors = []string{exceptions.ErrorAudienceForbidden.Error()}
		return ctx.Status(fiber.StatusForbidden).JSON(apiResponse)
	}

	return ctx.Next()
}

func IsRootUser(ctx *fiber.Ctx) error {
	apiResponse := common.NewApiResponse()
	userClaims := GetUserClaims(ctx)
	if strings.ToLower(userClaims.Audience) != "root" {
		apiResponse.Success = false
		apiResponse.Errors = []string{exceptions.ErrorAudienceForbidden.Error()}
		return ctx.Status(fiber.StatusForbidden).JSON(apiResponse)
	}

	return ctx.Next()
}
