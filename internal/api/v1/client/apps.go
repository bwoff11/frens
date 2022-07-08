package client

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type CreateAppRequest struct {
	ClientName   string `json:"client_name"`
	RedirectURIs string `json:"redirect_uris"`
	//Scopes       string `json:"scopes"`
	//Website      string `json:"website"`
}

type Application struct {
	Name         string `json:"name"`
	Website      string `json:"website"`
	VapidKey     string `json:"vapid_key"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func getAppOptions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AppOptions",
	})
}

// Create a new application to obtain OAuth2 credentials.
func createApp(c *fiber.Ctx) error {
	var req CreateAppRequest
	log.Println(string(c.Body()))
	if err := c.BodyParser(&req); err != nil {
		return c.Status(422).JSON(map[string]string{
			"message": "Validation failed: Redirect URI must be an absolute URI.",
		})
	}

	var resp Application
	return c.Status(200).JSON(resp)
}

// Confirm that the app's OAuth2 credentials work.
func verifyAppCredentials(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Verified",
	})
}
