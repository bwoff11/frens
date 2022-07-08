package client

import (
	"github.com/bwoff11/frens/internal/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

func GetFilters(c *fiber.Ctx) error {
	var resp []models.Filter
	return c.Status(200).JSON(resp)
}

func GetFilter(c *fiber.Ctx) error {
	var resp models.Filter
	return c.Status(200).JSON(resp)
}

func CreateFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFilter",
	})
}

func UpdateFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UpdateFilter",
	})
}

func DeleteFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFilter",
	})
}
