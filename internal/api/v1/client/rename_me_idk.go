package client

import "github.com/gofiber/fiber/v2"

func getAccounts(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Accounts",
	})
}

func getAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Account",
	})
}

func accountAction(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AccountAction",
	})
}

func approveAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "ApproveAccount",
	})
}

func rejectAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "RejectAccount",
	})
}

func enableAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "EnableAccount",
	})
}

func unsilenceAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnsilenceAccount",
	})
}

func unsuspendAccount(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnsuspendAccount",
	})
}

func getUserIdentityProofs(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserIdentityProofs",
	})
}
