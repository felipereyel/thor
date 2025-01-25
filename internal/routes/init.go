package routes

import (
	"thor/internal/components"
	"thor/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init(app *fiber.App, svcs *services.Services) error {
	app.Use(cors.New())
	app.Get("/", homeHandler)
	app.Get("/discard", discardHandler)
	app.Use("/statics", staticsHandler)
	app.Use("/healthz", healthzHandler)

	downloadsGroup := app.Group("/downloads")
	initDownloadsRoutes(downloadsGroup, svcs)

	app.Use(notFoundHandler)
	return nil
}

func homeHandler(c *fiber.Ctx) error {
	return sendPage(c, components.HomePage())
}

func healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func discardHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
