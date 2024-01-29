package routes

import (
	"bytes"
	"context"
	"fmt"
	"thor/internal/services"
	"thor/internal/web"

	"github.com/a-h/templ"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	c.SendStatus(fiber.StatusInternalServerError)
	fmt.Printf("Route Error [%s]: %v\n", c.Path(), err)
	return sendPage(c, web.ErrorPage())
}

func notFoundHandler(c *fiber.Ctx) error {
	c.SendStatus(fiber.StatusNotFound)
	return sendPage(c, web.NotFoundPage())
}

func sendPage(c *fiber.Ctx, page templ.Component) error {
	c.Set("Content-Type", "text/html")
	return page.Render(context.Background(), c)
}

func sendPageWS(c *websocket.Conn, page templ.Component) error {
	buf := new(bytes.Buffer)
	page.Render(context.Background(), buf)
	return c.WriteMessage(websocket.TextMessage, buf.Bytes())
}

type routeHandler func(svcs *services.Services, c *fiber.Ctx) error

func bindCtx(svcs *services.Services, handler routeHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(svcs, c)
	}
}

type wsRouteHandler func(svcs *services.Services, c *websocket.Conn)

func bindWSCtx(svcs *services.Services, handler wsRouteHandler) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		handler(svcs, c)
	})
}

func wsUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
