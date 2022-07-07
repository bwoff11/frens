package client

import "github.com/gofiber/fiber/v2"

func getFollowRequests(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FollowRequests",
	})
}

func acceptFollowRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AcceptFollowRequest",
	})
}

func rejectFollowRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "RejectFollowRequest",
	})
}
