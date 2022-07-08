package client

import (
	"github.com/gofiber/fiber/v2"
)

type AuthorizeUserRequest struct {
	ResponseType string `json:"response_type"` // Should be set equal to code.
	ClientID     string `json:"client_id"`     // Client ID, obtained during app registration.
	//RedirectURI  string `json:"redirect_uri"`  // Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect URIs declared during app registration.
	//Scope        string `json:"scope"`         // List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read.
	//ForceLogin   bool   `json:"force_login"`   // Added in 2.6.0. Forces the user to re-login, which is necessary for authorizing with multiple accounts from the same instance.
}

// Displays an authorization form to the user. If approved, it will create and return an authorization code, then redirect to the desired redirect_uri, or show the authorization code if urn:ietf:wg:oauth:2.0:oob was requested. The authorization code can be used while requesting a token to obtain access to user-level methods.
func authorizeUser(c *fiber.Ctx) error {
	var req AuthorizeUserRequest
	if err := c.BodyParser(&req); err != nil {
		//return err
	}

	return c.Redirect("/auth/sign_in")
}

// Returns an access token, to be used during API calls that are not public.
func getToken(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Token",
	})
}

// Revoke an access token to make it no longer valid for use.
func revokeToken(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "RevokeToken",
	})
}
