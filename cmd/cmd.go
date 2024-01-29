package cmd

import (
	"thor/internal/config"
	"thor/internal/routes"
	"thor/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command, args []string) error {
	cfg, err := config.GetServerConfigs()
	if err != nil {
		return err
	}

	svcs, err := services.NewServices(cfg)
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	routes.Init(app, svcs)

	err = app.Listen(cfg.ServerAddress)
	if err != nil {
		return err
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "thor",
	Short: "start thor server",
	RunE:  Start,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
