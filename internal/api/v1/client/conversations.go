package client

import "github.com/gofiber/fiber/v2"

func getSelfConversations(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func deleteConversation(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func markConversationRead(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
