package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bwoff11/frens/internal/config"
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/test"
	"github.com/gofiber/fiber/v2"
)

func TestCreateAccount(t *testing.T) {
	tests := []test.APITest{
		{
			Description: "Valid request",
			Method:      "POST",
			Route:       "/api/v1/accounts",
			ReqBody: CreateAppRequestBody{
				ClientName:   "Test App",
				RedirectURIs: "https://example.com/callback",
				Scopes:       "read",
				Website:      "https://example.com",
			},
			ExpectedCode: http.StatusCreated,
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

func TestVerifyCredentials(t *testing.T) {
	t.Skip("Not implemented")
}

func TestUpdateCredentials(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetUserFeaturedTags(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetRelationships(t *testing.T) {
	t.Skip("Not implemented")
}

func TestSearchUsers(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetUserInfo(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetUserStatuses(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetUserFollowers(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetUserFollowing(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetUserLists(t *testing.T) {
	t.Skip("Not implemented")
}

func TestFollowUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestUnfollowUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestBlockUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestUnblockUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestMuteUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestUnmuteUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestPinUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestUnpinUser(t *testing.T) {
	t.Skip("Not implemented")
}

func TestNoteUser(t *testing.T) {
	t.Skip("Not implemented")
}
