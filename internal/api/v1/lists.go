package v1

import "github.com/gofiber/fiber/v2"

func getLists(c *fiber.Ctx) error {
	return c.Status(200).SendString("Not implemented")
}

func getList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func createList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func updateList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func deleteList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getListAccounts(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func addAccountToList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func removeAccountFromList(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
