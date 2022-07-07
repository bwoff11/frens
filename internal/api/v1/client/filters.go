package client

import "github.com/gofiber/fiber/v2"

func getFilters(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Filters",
	})
}

func getFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Filter",
	})
}

func createFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFilter",
	})
}

func updateFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UpdateFilter",
	})
}

func deleteFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFilter",
	})
}
