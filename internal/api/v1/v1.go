package v1

import (
	"github.com/bwoff11/frens/internal/api/v1/client"
	"github.com/bwoff11/frens/internal/api/v1/server"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	client.AddRoutes(app)
	server.AddRoutes(app)
}
