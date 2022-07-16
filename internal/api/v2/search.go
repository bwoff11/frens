package v2

import (
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type SearchRequest struct {
	AccountID         string `json:"account_id"`         // If provided, statuses returned will be authored only by this account
	MaxID             string `json:"max_id"`             // Return results older than this id
	MinID             string `json:"min_id"`             // Return results immediately newer than this id
	Type              string `json:"type"`               // Enum(accounts, hashtags, statuses)
	ExcludeUnreviewed bool   `json:"exclude_unreviewed"` // Filter out unreviewed tags? Defaults to false. Use true when trying to find trending tags.
	Resolve           bool   `json:"resolve"`            // Attempt WebFinger lookup. Defaults to false.
	Limit             int    `json:"limit"`              // Maximum number of results to load, per type. Defaults to 20. Max 40.
	Offset            int    `json:"offset"`             // Offset in search results. Used for pagination. Defaults to 0.
	Following         bool   `json:"following"`          // Only include accounts that the user is following. Defaults to false.
	Q                 string `json:"q"`                  // Search query.
}

type SearchResponse struct {
	Accounts []models.Account `json:"accounts"`
	Statuses []models.Status  `json:"statuses"`
	Hashtags []models.Hashtag `json:"hashtags"`
}

func Search(c *fiber.Ctx) error {
	var req SearchRequest = SearchRequest{
		AccountID:         c.Query("account_id"),
		MaxID:             c.Query("max_id"),
		MinID:             c.Query("min_id"),
		Type:              c.Query("type"),
		ExcludeUnreviewed: c.Query("exclude_unreviewed") == "true",
		Resolve:           c.Query("resolve") == "true",
		//Limit:             c.Query("limit"),
		//Offset:            c.Query("offset"),
		Following: c.Query("following") == "true",
		Q:         c.Query("q"),
	}
	// validate here
	// req.validate() blah blah blah

	var resp SearchResponse

	SearchAccounts(req, &resp)
	SearchStatuses(req, &resp)
	SearchHashtags(req, &resp)

	return c.Status(200).JSON(resp)
}

func SearchAccounts(req SearchRequest, resp *SearchResponse) {
	var foundAccounts []models.Account
	db.Postgres.Where("username LIKE = ?", "%"+req.Q+"%").Find(&foundAccounts)
	resp.Accounts = foundAccounts
}

func SearchStatuses(req SearchRequest, resp *SearchResponse) {
	var foundStatuses []models.Status
	db.Postgres.Where("content LIKE = ?", "%"+req.Q+"%").Find(&foundStatuses)
	resp.Statuses = foundStatuses
}

func SearchHashtags(req SearchRequest, resp *SearchResponse) {
	var foundHashtags []models.Hashtag
	db.Postgres.Where("id LIKE = ?", "%"+req.Q+"%").Find(&foundHashtags)
	resp.Hashtags = foundHashtags
}
