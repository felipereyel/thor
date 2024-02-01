package cmd

import (
	"github.com/felipereyel/thor/internal/config"
	"github.com/felipereyel/thor/internal/routes"
	"github.com/felipereyel/thor/internal/services"

	"github.com/gofiber/fiber/v2"
)

func startFiber(cfg *config.ServerConfigs) error {
	svcs, err := services.NewServices(cfg)
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	routes.Init(app, svcs)

	return app.Listen(cfg.ServerAddress)
}
