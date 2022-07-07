package client

import "github.com/gofiber/fiber/v2"

func getFeaturedTags(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FeaturedTags",
	})
}

func createFeaturedTag(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFeaturedTag",
	})
}

func deleteFeaturedTag(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFeaturedTag",
	})
}

func getFeaturedTagSuggestions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FeaturedTagSuggestions",
	})
}
