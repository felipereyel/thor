package routes

import (
	"thor/internal/services"
	"thor/internal/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init(app *fiber.App, svcs *services.Services) error {
	app.Use(cors.New())
	app.Get("/", homeHandler)
	app.Use("/healthz", healthzHandler)

	downloadsGroup := app.Group("/downloads")
	initDownloadsRoutes(downloadsGroup, svcs)

	app.Use(notFoundHandler)
	return nil
}

func homeHandler(c *fiber.Ctx) error {
	return sendPage(c, web.HomePage())
}
