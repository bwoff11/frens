package client

import "github.com/gofiber/fiber/v2"

func getScheduledStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "getScheduledStatus",
	})
}

func getScheduledStatuses(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "createScheduledStatus",
	})
}

func scheduleStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "scheduleStatus",
	})
}

func deleteScheduledStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "deleteScheduledStatus",
	})
}
