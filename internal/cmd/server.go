package cmd

import (
	"thor/internal/routes"
	"thor/internal/services"

	"github.com/gofiber/fiber/v2"
)

func startFiber(svcs *services.Services) error {

	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	routes.Init(app, svcs)

	return app.Listen(svcs.Configs.ServerAddress)
}
