package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/test"
	"github.com/gofiber/fiber/v2"
)

func TestCreateApp(t *testing.T) {
	tests := []test.APITest{
		{
			Description: "Valid request",
			Method:      "POST",
			Route:       "/api/v1/apps",
			ReqBody: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Description: "Missing scopes",
			Method:      "POST",
			Route:       "/api/v1/apps",
			ReqBody: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "",
				Website:      "https://example.com",
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Description: "Missing website",
			Method:      "POST",
			Route:       "/api/v1/apps",
			ReqBody: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "",
			},
			ExpectedCode: http.StatusCreated,
		},
		{
			Description: "Missing client name",
			Method:      "POST",
			Route:       "/api/v1/apps",
			ReqBody: CreateAppRequestBody{
				ClientName:   "",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Description: "Client name too long",
			Method:      "POST",
			Route:       "/api/v1/apps",
			ReqBody: CreateAppRequestBody{
				ClientName:   "a" + strings.Repeat("b", 512),
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Description: "Website too long",
			Method:      "POST",
			Route:       "/api/v1/apps",
			ReqBody: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com" + strings.Repeat("b", 512),
			},
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			db.Connect()
			jsonBody, err := json.Marshal(test.ReqBody)
			if err != nil {
				t.Errorf("Error marshalling JSON: %v", err)
			}
			req, err := http.NewRequest(test.Method, test.Route, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			app := fiber.New(config.Router)
			AddRoutes(app)

			resp, _ := app.Test(req)
			if resp.StatusCode != test.ExpectedCode {
				t.Errorf("Expected status code %d, got %d", test.ExpectedCode, resp.StatusCode)
			}
		})
	}
}

func TestVerifyAppCredentials(t *testing.T) {
	t.Skip("Not implemented")
}
