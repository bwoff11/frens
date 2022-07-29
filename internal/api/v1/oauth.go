package v1

import (
	"encoding/binary"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/bwoff11/frens/internal/db"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type GetLoginPageRequest struct {
	ResponseType string `json:"response_type" form:"response_type"` // Should be set equal to code.
	ClientID     string `json:"client_id" form:"client_id`          // Client ID, obtained during app registration.
	RedirectURI  string `json:"redirect_uri" form:"redirect_uri`    // Set a URI to redirect the user to. If this parameter is set to urn:ietf:wg:oauth:2.0:oob then the authorization code will be shown instead. Must match one of the redirect URIs declared during app registration.
	Scope        string `json:"scope" form:"scope"`                 // List of requested OAuth scopes, separated by spaces (or by pluses, if using query parameters). Must be a subset of scopes declared during app registration. If not provided, defaults to read.
	ForceLogin   string `json:"force_login" form:"force_login`      // Added in 2.6.0. Forces the user to re-login, which is necessary for authorizing with multiple accounts from the same instance.
}

// Returns a login page for the user.
func GetLoginPage(c *fiber.Ctx) error {

	// Convert query parameters to struct
	req := GetLoginPageRequest{
		ResponseType: c.Query("response_type"),
		ClientID:     c.Query("client_id"),
		RedirectURI:  c.Query("redirect_uri"),
		Scope:        c.Query("scope"),
		ForceLogin:   c.Query("force_login"),
	}

	// TODO: Add IDP provider options POSSIBLY (e.g. Google, Facebook, etc.)

	// Return login form
	return c.Render("./public/login.html", req)
}

// After the user provides a username and password, it is checked against the db.
// If the user is valid, a code is generated that maps to the user's account.
// The code is stored then sent to the end user to request a token.
func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	// Client should encrypt the password before sending it.
	// As a temporary solution, we will just hash it here
	password = db.Sha256(password)

	// Find matching account in db
	var account db.Account
	if err := db.Postgres.Where("username = ? AND password = ?", username, password).First(&account).Error; err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Generate authorization code
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	codeRunes := make([]rune, 32)
	for i := range codeRunes {
		codeRunes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	codeString := string(codeRunes)
	codeBytes := []byte(codeString)

	// Store code to id map in cache
	accountIDBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(accountIDBytes, uint64(account.ID))
	db.Badger.Update(func(txn *badger.Txn) error {
		if err := txn.Set(codeBytes, accountIDBytes); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})

	// Return redirect URL
	return c.Redirect("http://localhost:4002/settings/instances/add?code=" + codeString)
}

func LoginStyle(c *fiber.Ctx) error {
	return c.Status(200).SendFile("./public/login_style.css")
}

func SignupPage(c *fiber.Ctx) error {
	return c.Status(200).SendFile("./public/signup.html")
}

func SignupStyle(c *fiber.Ctx) error {
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
func GetToken(c *fiber.Ctx) error {
	var reqBody GetTokenRequest
	if err := c.BodyParser(&reqBody); err != nil {
		return err
	}

	//validate here

	if reqBody.Code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	} else {
		log.Println("Code:", reqBody.Code)
	}

	// Lookup code in cache
	var accountIDBytes = make([]byte, 8)
	db.Badger.View(func(txn *badger.Txn) error {
		res, err := txn.Get([]byte(reqBody.Code))
		if err != nil {
			return err
		}
		if res == nil {
			return errors.New("Code not found")
		}
		res.ValueCopy(accountIDBytes)
		return nil
	})
	accountID := binary.LittleEndian.Uint64(accountIDBytes)
	if accountID == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"account_id": accountID,
		"client_id":  reqBody.ClientID,
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
	log.Println("Token:", t)
	return c.Status(fiber.StatusOK).JSON(resp)
}

// Revoke an access token to make it no longer valid for use.
func RevokeToken(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
