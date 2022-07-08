package client

import "github.com/gofiber/fiber/v2"

// Displays an authorization form to the user. If approved, it will create and return an authorization code, then redirect to the desired redirect_uri, or show the authorization code if urn:ietf:wg:oauth:2.0:oob was requested. The authorization code can be used while requesting a token to obtain access to user-level methods.
func authorizeUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AuthorizeUser",
	})
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
