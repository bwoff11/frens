package client

import "github.com/gofiber/fiber/v2"

func getLoginPage(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "not implemented",
	})
}
