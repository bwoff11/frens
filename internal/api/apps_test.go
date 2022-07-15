package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/bwoff11/frens/internal/db"
	"github.com/gofiber/fiber/v2"
)

func TestCreateApp(t *testing.T) {
	tests := []struct {
		description  string               // description of the test case
		method       string               // HTTP method
		route        string               // route path to test
		body         CreateAppRequestBody // request body
		expectedCode int                  // expected HTTP status code
		// expectedBody string // expected response body
	}{
		{
			description: "Valid request",
			method:      "POST",
			route:       "/api/v1/apps",
			body: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			expectedCode: http.StatusCreated,
		},
		{
			description: "Missing scopes",
			method:      "POST",
			route:       "/api/v1/apps",
			body: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "",
				Website:      "https://example.com",
			},
			expectedCode: http.StatusCreated,
		},
		{
			description: "Missing website",
			method:      "POST",
			route:       "/api/v1/apps",
			body: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "",
			},
			expectedCode: http.StatusCreated,
		},
		{
			description: "Missing client name",
			method:      "POST",
			route:       "/api/v1/apps",
			body: CreateAppRequestBody{
				ClientName:   "",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			description: "Client name too long",
			method:      "POST",
			route:       "/api/v1/apps",
			body: CreateAppRequestBody{
				ClientName:   "a" + strings.Repeat("b", 512),
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			description: "Website too long",
			method:      "POST",
			route:       "/api/v1/apps",
			body: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com" + strings.Repeat("b", 512),
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			db.Connect()
			jsonBody, err := json.Marshal(test.body)
			if err != nil {
				t.Errorf("Error marshalling JSON: %v", err)
			}
			req, err := http.NewRequest(test.method, test.route, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			app := fiber.New()
			app.Post("/api/v1/apps", createApp)

			resp, _ := app.Test(req)
			if resp.StatusCode != test.expectedCode {
				t.Errorf("Expected status code %d, got %d", test.expectedCode, resp.StatusCode)
			}
		})
	}
}

func TestVerifyAppCredentials(t *testing.T) {
	t.Skip("Not implemented")
}
