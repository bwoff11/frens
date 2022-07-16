package v1

import "github.com/gofiber/fiber/v2"

func ReportUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
