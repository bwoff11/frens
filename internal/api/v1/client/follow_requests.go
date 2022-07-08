package client

import "github.com/gofiber/fiber/v2"

func GetFollowRequests(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FollowRequests",
	})
}

func AcceptFollowRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AcceptFollowRequest",
	})
}

func RejectFollowRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "RejectFollowRequest",
	})
}
