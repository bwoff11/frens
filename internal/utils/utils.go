package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	user := c.Locals("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims), nil
}
