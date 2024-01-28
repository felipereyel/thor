package cmd

import (
	"thor/internal/config"
	"thor/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command, args []string) {
	cfg := config.GetServerConfigs()

	app := fiber.New(fiber.Config{
		ErrorHandler: routes.ErrorHandler,
	})

	app.Use(cors.New())
	routes.Init(app, cfg)

	if err := app.Listen(cfg.ServerAddress); err != nil {
		panic(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "thor",
	Short: "start thor server",
	Run:   Start,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
