package client

import "github.com/gofiber/fiber/v2"

func viewPoll(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "viewPoll",
	})
}

func voteOnPoll(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "voteOnPoll",
	})
}
