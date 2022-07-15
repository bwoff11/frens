package api

import "github.com/gofiber/fiber/v2"

func GetSelfMutes(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Mutes",
	})
}

func GetSelfBlocks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Blocks",
	})
}

func GetSelfEndorsements(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Endorsements",
	})
}

func GetSelfPreferences(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Preferences",
	})
}
