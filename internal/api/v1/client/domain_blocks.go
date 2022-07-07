package client

import "github.com/gofiber/fiber/v2"

func getDomainBlocks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DomainBlocks",
	})
}

func blockDomain(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "BlockDomain",
	})
}

func unblockDomain(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnblockDomain",
	})
}
