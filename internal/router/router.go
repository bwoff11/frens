package router

import (
	v1 "github.com/bwoff11/frens/internal/api/v1"
	"github.com/bwoff11/frens/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Create() *fiber.App {
	app := fiber.New(config.Router)

	addStaticRoutes(app)
	addMiddleware(app)
	v1.AddRoutes(app)

	return app
}

func Start() {
	app := Create()
	app.Listen(":4000")
}

func addStaticRoutes(app *fiber.App) {
	app.Static("/", "./public")
	app.Static("/favicon.ico", "./public/favicon.ico")
}

func addMiddleware(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Use(logger.New(logger.Config{
		//Format:       "method=${method}, uri=${uri}, status=${status}, latency=${latency}, bytes=${bytes}",
		//TimeFormat:   "2006-01-02 15:04:05",
		//TimeZone:     "UTC",
		//TimeInterval: time.Second,
		//Output:       os.Stdout,
	}))
}
