package main

import (
	v1 "github.com/bwoff11/frens/internal/api/v1"
	"github.com/bwoff11/frens/internal/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()
	api := fiber.New()
	v1.AddRoutes(api)

	api.Listen(":8080")
}
