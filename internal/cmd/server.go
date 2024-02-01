package cmd

import (
	"github.com/felipereyel/thor/internal/routes"
	"github.com/felipereyel/thor/internal/services"

	"github.com/gofiber/fiber/v2"
)

func startFiber(svcs *services.Services) error {

	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	routes.Init(app, svcs)

	return app.Listen(svcs.Configs.ServerAddress)
}
