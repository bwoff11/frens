package client

import "github.com/gofiber/fiber/v2"

func getSuggestions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Suggestions",
	})
}

func deleteSuggestion(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteSuggestion",
	})
}
