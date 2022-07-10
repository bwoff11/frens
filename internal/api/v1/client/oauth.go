package client

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func addOAuthRoutes(app *fiber.App) {
	oat := app.Group("/oauth")
	oat.Get("/authorize", getLoginPage)
	oat.Post("/authorize", login)
	oat.Get("/login_style.css", loginStyle)
	oat.Get("/signup.html", signupPage)
	oat.Get("/signup_style.css", signupStyle)
	oat.Post("/token", getToken)
	oat.Post("/revoke", revokeToken)
}

type GetLoginPageRequest struct {
	ResponseType string `json:"response_type" form:"response_type"` // Should be set equal to code.
	ClientID     string `json:"client_id" form:"client_id`          // Client ID, obtained during app registration.
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri`    // Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect URIs declared during app registration.
	Scope        string `json:"scope" form:"scope"`                 // List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read.
	ForceLogin   string `json:"force_login" form:"force_login`      // Added in 2.6.0. Forces the user to re-login, which is necessary for authorizing with multiple accounts from the same instance.
}

// Returns a login page for the user.
func getLoginPage(c *fiber.Ctx) error {
	req := GetLoginPageRequest{
		ResponseType: c.Query("response_type"),
		ClientID:     c.Query("client_id"),
		RedirectURI:  c.Query("redirect_uri"),
		Scope:        c.Query("scope"),
		ForceLogin:   c.Query("force_login"),
	}

	log.Println(req.RedirectURI)

	// Return login form
	return c.Render("./public/login.html", req)
	//return c.Status(200).SendFile("./public/login.html")
}

func login(c *fiber.Ctx) error {
	//req := GetTokenRequest{}

	//return c.Status(200).SendString("Logged in")
	return c.Redirect("http://localhost:4002/settings/instances/add?code=tFwEduJ4Act2PTmZ0osjp175WX-dfrOxO_H7W20-9rU")
}

func loginStyle(c *fiber.Ctx) error {
	return c.Status(200).SendFile("./public/login_style.css")
}

func signupPage(c *fiber.Ctx) error {
	return c.Status(200).SendFile("./public/signup.html")
}

func signupStyle(c *fiber.Ctx) error {
	return c.Status(200).SendFile("./public/signup_style.css")
}

type GetTokenRequest struct {
	GrantType    string `json:"grant_type" form:"grant_type"`      // Must be set equal to authorization_code.
	ClientID     string `json:"client_id" form:"client_id`         // Client ID, obtained during app registration.
	ClientSecret string `json:"client_secret" form:"client_secret` // Client secret, obtained during app registration.
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri`   // Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect URIs declared during app registration.
	Scope        string `json:"scope" form:"scope"`                // List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read.
	Code         string `json:"code" form:"code"`                  // Authorization code obtained from the authorization server.
}

type GetTokenRespone struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	CreatedAt   int64  `json:"created_at"`
}

// The client sends the application data we returned in the application registration form.
// Here, we use that information to generate a access token.
func getToken(c *fiber.Ctx) error {

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	resp := GetTokenRespone{
		AccessToken: t,
		TokenType:   "Bearer",
		Scope:       "read",
		CreatedAt:   1579098983,
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

// Revoke an access token to make it no longer valid for use.
func revokeToken(c *fiber.Ctx) error {

	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
