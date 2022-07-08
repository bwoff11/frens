package accounts

import "github.com/gofiber/fiber/v2"

func GetDomainBlocks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DomainBlocks",
	})
}

func BlockDomain(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "BlockDomain",
	})
}

func UnblockDomain(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnblockDomain",
	})
}
