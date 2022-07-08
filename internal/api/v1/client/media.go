package client

import "github.com/gofiber/fiber/v2"

func attachMedia(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "attachMedia",
	})
}

func getMedia(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "getMedia",
	})
}

func updateMedia(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "updateMedia",
	})
}
