package client

import "github.com/gofiber/fiber/v2"

func getInstance(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Instance",
	})
}

func getPeers(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Peers",
	})
}

func getInstanceActivity(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "InstanceActivity",
	})
}
