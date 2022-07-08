package client

import "github.com/gofiber/fiber/v2"

func GetFilters(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Filters",
	})
}

func GetFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Filter",
	})
}

func CreateFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFilter",
	})
}

func UpdateFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UpdateFilter",
	})
}

func DeleteFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFilter",
	})
}
