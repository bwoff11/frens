package apps

import "github.com/gofiber/fiber/v2"

func AddRoutes(app *fiber.App) {
	aps := app.Group("/api/v1/apps")
	aps.Options("/", createApp)
	aps.Get("/verify_credentials", verifyAppCredentials)

	oat := app.Group("/api/v1/oauth")
	oat.Get("/authorize", authorizeUser)
	oat.Get("/token", getToken)
	oat.Post("/revoke", revokeToken)
}

// Create a new application to obtain OAuth2 credentials.
func createApp(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "App",
	})
}

// Confirm that the app's OAuth2 credentials work.
func verifyAppCredentials(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Verified",
	})
}
