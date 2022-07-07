package main

import (
	v1 "github.com/bwoff11/frens/internal/api/v1"
	"github.com/bwoff11/frens/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db.Connect()
	api := fiber.New()
	api.Use(logger.New())
	api.Use(recover.New())

	v1.AddRoutes(api)

	api.Listen(":8080")
}
