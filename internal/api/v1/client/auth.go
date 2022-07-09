package client

import "github.com/gofiber/fiber/v2"

func getLoginPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
