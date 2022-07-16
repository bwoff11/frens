package v1

import "github.com/gofiber/fiber/v2"

func GetAnnouncements(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func DismissAnnouncement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func ReactToAnnouncement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func UnreactToAnnouncement(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
