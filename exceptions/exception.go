package exceptions

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

var (
	ErrCardNotFound                  = errors.New("card not found")
	ErrDistributorAudienceRequired   = errors.New("user must be a distributor")
	ErrorCardAlreadyLinked           = errors.New("card already linked to holder")
	ErrHolderMappingAlreadyExists    = errors.New("holder mapping already exists")
	ErrorAudienceForbidden           = errors.New("operation not allowed for user")
	ErrCardCustomerNotFound          = errors.New("card has not been purchased by a customer")
	ErrCardDistributorNotFound       = errors.New("card has not been purchased by a distributor")
	ErrOperationNotAllowed           = errors.New("sorry, you are not allowed to perform this operation")
	ErrNoCardsAvailableInBatch       = errors.New("sorry, this batch has no available cards")
	ErrBatchQuantityMappingsTooLarge = errors.New("sorry, this batch doesn't have enough cards to meet the quantity requirement")
)

func GetErrorCode(err error) int {
	statusCode := fiber.StatusInternalServerError
	if errors.Is(err, ErrCardNotFound) {
		statusCode = fiber.StatusNotFound
	}

	if errors.Is(err, ErrorCardAlreadyLinked) {
		statusCode = fiber.StatusUnprocessableEntity
	}

	if errors.Is(err, ErrorAudienceForbidden) {
		statusCode = fiber.StatusForbidden
	}

	if errors.Is(err, ErrOperationNotAllowed) {
		statusCode = fiber.StatusForbidden
	}

	if errors.Is(err, ErrHolderMappingAlreadyExists) {
		statusCode = fiber.StatusUnprocessableEntity
	}

	if errors.Is(err, ErrCardDistributorNotFound) {
		statusCode = fiber.StatusUnprocessableEntity
	}

	if errors.Is(err, ErrCardCustomerNotFound) {
		statusCode = fiber.StatusUnprocessableEntity
	}

	return statusCode
}

func GetValidatorErrors(err error) []string {
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)
	errorMessages := make([]string, 0)
	for _, fieldError := range validationErrors {
		log.Println("fieldError.Error()", fieldError.Error())
		log.Println("fieldError.Field()", fieldError.Field())
		log.Println("fieldError.Tag()", fieldError.Tag())
		log.Println("fieldError.Type()", fieldError.Type())
		log.Println("fieldError.Value()", fieldError.Value())
		log.Println("fieldError.Type()", fieldError.Type())
		log.Println("fieldError.ActualTag()", fieldError.ActualTag())
		log.Println("fieldError.Kind()", fieldError.Kind())
		log.Println("fieldError.Namespace()", fieldError.Namespace())
		log.Println("fieldError.StructField()", fieldError.StructField())
		log.Println("fieldError.Param()", fieldError.Param())
		//errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", fieldError.Tag(), fieldError.Value()))
		errorMessages = append(errorMessages, fieldError.Error())
	}

	return errorMessages
}
