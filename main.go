package main

import (
	"log"

	v1 "github.com/bwoff11/frens/internal/api/v1"
	"github.com/bwoff11/frens/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db.Connect()
	api := fiber.New()

	api.Use(func(c *fiber.Ctx) error {
		log.Println("Body:", string(c.Body()))
		return c.Next()
	})
	api.Use(logger.New())
	api.Use(recover.New())
	api.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	api.Static("/", "./public")

	v1.AddRoutes(api)

	log.Fatal(api.Listen(":4000"))
}
