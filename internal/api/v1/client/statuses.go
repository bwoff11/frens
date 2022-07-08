package client

import "github.com/gofiber/fiber/v2"

func createStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "createStatus",
	})
}

func getStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "getStatus",
	})
}

func deleteStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "deleteStatus",
	})
}

func getStatusContext(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "getStatusContext",
	})
}

func getFavoritedBy(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "getFavoritedBy",
	})
}

func favouriteStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "favouriteStatus",
	})
}

func unfavouriteStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "unfavouriteStatus",
	})
}

func reblogStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "reblogStatus",
	})
}

func unreblogStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "unreblogStatus",
	})
}

func bookmarkStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "bookmarkStatus",
	})
}

func unbookmarkStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "unbookmarkStatus",
	})
}

func muteStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "muteStatus",
	})
}

func unmuteStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "unmuteStatus",
	})
}

func pinStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "pinStatus",
	})
}

func unpinStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "unpinStatus",
	})
}
