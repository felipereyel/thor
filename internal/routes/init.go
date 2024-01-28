package routes

import (
	"thor/internal/config"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App, cfg config.ServerConfigs) error {
	app.Get("/torrents", listTorrents)
	app.Post("/torrents", createTorrent)

	app.Use("/healthz", healthzHandler)
	app.Use(notFoundHandler)

	return nil
}
