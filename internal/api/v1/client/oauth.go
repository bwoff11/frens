package client

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthorizeUserRequest struct {
	ResponseType string `json:"response_type" form:"response_type"` // Should be set equal to code.
	ClientID     string `json:"client_id" form:"client_id`          // Client ID, obtained during app registration.
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri`    // Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect URIs declared during app registration.
	Scope        string `json:"scope" form:"scope"`                 // List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read.
	ForceLogin   bool   `json:"force_login" form:"force_login`      // Added in 2.6.0. Forces the user to re-login, which is necessary for authorizing with multiple accounts from the same instance.
}

// Displays an authorization form to the user. If approved, it will create and return an authorization code, then redirect to the desired redirect_uri, or show the authorization code if urn:ietf:wg:oauth:2.0:oob was requested. The authorization code can be used while requesting a token to obtain access to user-level methods.
func authorizeUser(c *fiber.Ctx) error {
	req := AuthorizeUserRequest{
		ResponseType: c.Query("response_type"),
		ClientID:     c.Query("client_id"),
		RedirectURI:  c.Query("redirect_uri"),
		Scope:        c.Query("scope"),
		ForceLogin:   c.Query("force_login") == "true",
	}

	log.Println(req.ClientID)

	redirect_uri := c.Query("redirect_uri")
	return c.Redirect(redirect_uri + "?code=12345")
}

type GetTokenRequest struct {
	GrantType    string `json:"grant_type" form:"grant_type"`      // Must be set equal to authorization_code.
	ClientID     string `json:"client_id" form:"client_id`         // Client ID, obtained during app registration.
	ClientSecret string `json:"client_secret" form:"client_secret` // Client secret, obtained during app registration.
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri`   // Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect URIs declared during app registration.
	Scope        string `json:"scope" form:"scope"`                // List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read.
	Code         string `json:"code" form:"code"`                  // Authorization code obtained from the authorization server.
}

// Returns an access token, to be used during API calls that are not public.
func getToken(c *fiber.Ctx) error {
	/*req := GetTokenRequest{
		GrantType:    c.Query("grant_type"),
		ClientID:     c.Query("client_id"),
		ClientSecret: c.Query("client_secret"),
		RedirectURI:  c.Query("redirect_uri"),
		Scope:        c.Query("scope"),
		Code:         c.Query("code"),
	}*/

	// do something

	newToken := Token{
		AccessToken: "12345",
		TokenType:   "Bearer",
		Scope:       "read write follow push",
		CreatedAt:   uint64(time.Now().Unix()),
	}
	return c.Status(fiber.StatusOK).JSON(newToken)
}

// Revoke an access token to make it no longer valid for use.
func revokeToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
