package http

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ceblay/billing-demo/pkg/app"
)

type HttpServer struct {
	app app.Application
}

func NewServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) Run() {
	_app := fiber.New()
	healthRoutes := _app.Group("healthz")
	historyRoutes := _app.Group("v1").Group("billing").Group("histories")

	historyRoutes.Get("/", func(c *fiber.Ctx) error {
		result, err := h.app.Queries.AllBillingHistory.Handle()
		if err != nil {
			c.SendString("An error occurred while getting history")
		}

		return c.JSON(fiber.Map{
			"message": result,
		})
	})

	healthRoutes.Get("/ready", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "HTTP Server is ready",
		})
	})

	healthRoutes.Get("/live", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "HTTP Server is live",
		})
	})

	log.Fatal(_app.Listen(":7000"))
}
