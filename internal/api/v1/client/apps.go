package client

import (
	"log"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateAppRequest struct {
	ClientName   string `json:"client_name"`   // A name for your application
	RedirectURIs string `json:"redirect_uris"` // Where the user should be redirected after authorization. To display the authorization code to the user instead of redirecting to a web page, use urn:ietf:wg:oauth:2.0:oob in this parameter.
	Scopes       string `json:"scopes"`        // Space separated list of scopes. If none is provided, defaults to read.
	Website      string `json:"website"`       // A URL to the homepage of your app
}

// Create a new application to obtain OAuth2 credentials.
func createApp(c *fiber.Ctx) error {

	// Parse the request body
	var req CreateAppRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(422).JSON(map[string]string{
			"error": "Invalid request body",
		})
	}

	// Validate here
	// Be sure to add defaul values for scope (read) at some point

	// Create a full application object
	newApp := models.Application{
		Name:         req.ClientName,
		Website:      req.Website,
		RedirectURIs: req.RedirectURIs,
		ClientID:     uuid.New().String()[:32],
		ClientSecret: uuid.New().String()[:32],
		Scopes:       req.Scopes,
	}

	// Store full application in database
	if err := db.DB.Create(&newApp).Error; err != nil {
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}

	// Create a oauth2 compatible application
	newOAuthApp := models.OAuthApplication{
		ID:     newApp.ClientID,
		Secret: newApp.ClientSecret,
		Domain: req.Website,
		UserID: "",
	}

	// Store oauth2 application in database
	if err := db.DB.Create(&newOAuthApp).Error; err != nil {
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}

	// Create a version of the application that is safe to return to the client
	newAppResponse := models.ApplicationResponse{
		ID:           newApp.ClientID,
		Name:         newApp.Name,
		Website:      newApp.Website,
		RedirectURIs: newApp.RedirectURIs,
		ClientID:     newApp.ClientID,
		ClientSecret: newApp.ClientSecret,
	}

	// Return the application to the client
	return c.Status(200).JSON(newAppResponse)
}

// Confirm that the app's OAuth2 credentials work.
func verifyAppCredentials(c *fiber.Ctx) error {
	token := c.Get("Bearer")
	log.Println("Token:", token)
	return nil
}
