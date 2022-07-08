package client

import "github.com/gofiber/fiber/v2"

func GetFeaturedTags(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FeaturedTags",
	})
}

func CreateFeaturedTag(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFeaturedTag",
	})
}

func DeleteFeaturedTag(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFeaturedTag",
	})
}

func GetFeaturedTagSuggestions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FeaturedTagSuggestions",
	})
}
