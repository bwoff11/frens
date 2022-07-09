package v2

import "github.com/gofiber/fiber/v2"

func AddRoutes(app *fiber.App) {
	app.Get("/api/v2/search", search)
}
