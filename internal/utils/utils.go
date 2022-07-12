package utils

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetAccountID(c *fiber.Ctx) (*uint64, error) {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	accountIDfloat := claims["account_id"].(float64)
	accountIDStr := strconv.FormatFloat(accountIDfloat, 'f', 0, 64)
	if accountIDStr == "" {
		return nil, fmt.Errorf("no account ID found in JWT")
	}
	accountIDuint64, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing account ID: %v", err)
	}
	return &accountIDuint64, nil
}
