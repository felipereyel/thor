package routes

import (
	"fmt"
	"thor/internal/services"
	"thor/internal/web"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func wsUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func Init(app *fiber.App, svcs *services.Services) error {
	app.Use(cors.New())
	app.Use("/healthz", healthzHandler)

	app.Get("/", bindCtx(svcs, homeHandler))
	app.Post("/downloads", bindCtx(svcs, createTorrent))
	app.Get("/downloads/ws", wsUpgrade, bindWSCtx(svcs, wsDownloadsHandler))

	app.Use(notFoundHandler)

	return nil
}

func homeHandler(svcs *services.Services, c *fiber.Ctx) error {
	return sendPage(c, web.HomePage())
}

func wsDownloadsHandler(svcs *services.Services, c *websocket.Conn) {
	for {
		torrs := svcs.Download.ListDownloads()
		component := web.DownloadList(torrs)

		if err := sendPageWS(c, component); err != nil {
			c.Close()
			return
		}

		time.Sleep(2 * time.Second)
	}
}

func createTorrent(svcs *services.Services, c *fiber.Ctx) error {
	type request struct {
		Hash string `json:"hash"`
	}

	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if req.Hash == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := svcs.Download.AddDownload(req.Hash)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}
