package v1

import (
	"github.com/gofiber/fiber/v2"
)

func getFilters(c *fiber.Ctx) error {
	return c.Status(200).SendString("Not implemented")
}

func getFilter(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func createFilter(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func updateFilter(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func deleteFilter(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
