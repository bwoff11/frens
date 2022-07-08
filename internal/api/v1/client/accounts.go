package client

import "github.com/gofiber/fiber/v2"

func GetAccounts(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Accounts",
	})
}

func GetAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Account",
	})
}

func AccountAction(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AccountAction",
	})
}

func ApproveAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "ApproveAccount",
	})
}

func RejectAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "RejectAccount",
	})
}

func EnableAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "EnableAccount",
	})
}

func UnsilenceAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnsilenceAccount",
	})
}

func UnsuspendAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnsuspendAccount",
	})
}
