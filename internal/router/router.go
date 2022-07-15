package router

import (
	"encoding/json"
	"time"

	"github.com/bwoff11/frens/internal/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:                      false,
		ServerHeader:                 "Frens",
		StrictRouting:                false,
		CaseSensitive:                false,
		Immutable:                    false,
		UnescapePath:                 false,
		ETag:                         false,
		BodyLimit:                    4 * 1024 * 1024,
		Concurrency:                  256 * 1024,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  (60 * time.Second),
		WriteTimeout:                 (60 * time.Second),
		IdleTimeout:                  (60 * time.Second),
		ReadBufferSize:               4096,
		WriteBufferSize:              4096,
		CompressedFileSuffix:         ".gz",
		ProxyHeader:                  "",
		GETOnly:                      false,
		ErrorHandler:                 fiber.DefaultErrorHandler,
		DisableKeepalive:             false,
		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		DisableStartupMessage:        false,
		AppName:                      "Frens",
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		DisablePreParseMultipartForm: false,
		StreamRequestBody:            false,
		EnablePrintRoutes:            false,
		Network:                      fiber.NetworkTCP4,
		JSONEncoder:                  json.Marshal,
		JSONDecoder:                  json.Unmarshal,
	})

	addStaticRoutes(app)
	addMiddleware(app)
	api.AddRoutes(app)

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
