package routes

import (
	"github.com/gofiber/fiber/v2"
)

func createTorrent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func listTorrents(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
