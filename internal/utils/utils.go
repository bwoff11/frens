package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetAccountID(c *fiber.Ctx) (*int, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idFloat := claims["account_id"].(float64)
	idInt := int(idFloat)

	if idInt == 0 {
		return nil, errors.New("no account ID found in JWT")
	}

	return &idInt, nil
}
