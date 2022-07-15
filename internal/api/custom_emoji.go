package api

import "github.com/gofiber/fiber/v2"

func getCustomEmojis(c *fiber.Ctx) error {
	return c.Status(200).SendString("Not implemented")
}
