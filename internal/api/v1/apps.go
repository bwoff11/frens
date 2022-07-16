package v1

import (
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateAppRequestBody struct {
	ClientName   string `json:"client_name" form:"client_name" validate:"required,max=64"` // A name for your application
	RedirectURIs string `json:"redirect_uris" form:"redirect_uris validate:"required"`     // Where the user should be redirected after authorization. To display the authorization code to the user instead of redirecting to a web page, use urn:ietf:wg:oauth:2.0:oob in this parameter.
	Scopes       string `json:"scopes" form:"scopes"`                                      // Space separated list of scopes. If none is provided, defaults to read.
	Website      string `json:"website" form:"website" validate:"max=64"`                  // A URL to the homepage of your app
}

// Create a new application to obtain OAuth2 credentials.
// The application then makes a POST request to /api/v1/apps to obtain an OAuth2 "app" token.
// Null scope defaults to read.
// See: https://docs.joinmastodon.org/client/token/
func createApp(c *fiber.Ctx) error {
	var req CreateAppRequestBody
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Error parsing request body: " + err.Error(),
		})
	}

	// Check scope for value and default to read if not provided
	if req.Scopes == "" {
		req.Scopes = "read"
	}

	if err := validator.New().Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

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
	if err := db.Postgres.Create(&newApp).Error; err != nil {
		// Add error logging here
		return c.Status(500).JSON(map[string]string{
			"error": "Internal server error",
		})
	}

	// Return the application to the client
	return c.Status(fiber.StatusCreated).JSON(newApp)
}

// Confirm that the app's OAuth2 credentials work.
func verifyAppCredentials(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
