package client

import "github.com/gofiber/fiber/v2"

func GetSuggestions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Suggestions",
	})
}

func DeleteSuggestion(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteSuggestion",
	})
}
