package routes

import (
	"context"
	"fmt"
	"thor/internal/web"

	"github.com/a-h/templ"
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

func healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func sendPage(c *fiber.Ctx, page templ.Component) error {
	c.Set("Content-Type", "text/html")
	return page.Render(context.Background(), c)
}
