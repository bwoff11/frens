package client

import "github.com/gofiber/fiber/v2"

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
