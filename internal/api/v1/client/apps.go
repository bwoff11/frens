package client

import (
	"github.com/gofiber/fiber/v2"
)

type CreateAppRequest struct {
	ClientName   string `json:"client_name"`
	RedirectURIs string `json:"redirect_uris"`
	Scopes       string `json:"scopes"`
	Website      string `json:"website"`
}

type Application struct {
	Name         string `json:"name"`
	Website      string `json:"website"`
	VapidKey     string `json:"vapid_key"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Create a new application to obtain OAuth2 credentials.
func createApp(c *fiber.Ctx) error {
	var req CreateAppRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(422).JSON(map[string]string{
			"message": "Validation failed: Redirect URI must be an absolute URI.",
		})
	}

	var resp Application
	resp.Name = req.ClientName
	resp.Website = req.Website
	resp.ClientID = "client_id"
	resp.ClientSecret = "supersecret"
	resp.VapidKey = "vapid_key"
	return c.Status(200).JSON(resp)
}

// Confirm that the app's OAuth2 credentials work.
func verifyAppCredentials(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
