package client

import "github.com/gofiber/fiber/v2"

func GetSelfFavourites(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Favourites",
	})
}
