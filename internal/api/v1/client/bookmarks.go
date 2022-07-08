package client

import "github.com/gofiber/fiber/v2"

func GetSelfBookmarks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "GetSelfBookmarks",
	})
}
