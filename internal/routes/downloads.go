package routes

import (
	"fmt"
	"thor/internal/services"
	"thor/internal/web"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func initDownloadsRoutes(gp fiber.Router, svcs *services.Services) {
	gp.Post("/", bindCtx(svcs, createDownload))
	gp.Get("/ws", wsUpgrade, bindWSCtx(svcs, wsDownloadsHandler))
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

func createDownload(svcs *services.Services, c *fiber.Ctx) error {
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

	err := svcs.Download.AddDownload(req.Hash)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}
