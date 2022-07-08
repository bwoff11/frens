package client

import (
	"time"

	"github.com/bwoff11/frens/internal/api/v1/models"
	"github.com/gofiber/fiber/v2"
)

type PublicTimelineRequest struct {
	Local     bool   `json:"local"`      // Show only local statuses? Defaults to false.
	Remote    bool   `json:"remote"`     // Show only remote statuses? Defaults to false.
	OnlyMedia bool   `json:"only_media"` // Show only statuses with media attached? Defaults to false.
	MaxID     string `json:"max_id"`     // Return results older than this id
	SinceID   string `json:"since_id"`   // Return results newer than this id
	MinID     string `json:"min_id"`     // Return results immediately newer than this id
	Limit     int    `json:"limit"`      // Maximum number of results to return. Defaults to 20.
}

func getPublicTimelines(c *fiber.Ctx) error {
	var req PublicTimelineRequest
	if err := c.BodyParser(&req); err != nil {
		// check back on this later
		//return err
	}

	var resp []models.Status
	resp = append(resp, models.Status{
		ID:                 "253",
		CreatedAt:          time.Now(),
		InReplyToID:        nil,
		InReplyToAccountID: nil,
		Sensitive:          false,
		SpoilerText:        "",
		Visibility:         "public",
		Language:           "en",
		URI:                "https://frens.com/status/1",
		URL:                "https://frens.com/status/1",
		RepliesCount:       3,
		ReblogsCount:       5,
		FavouritesCount:    6,
		Favourited:         false,
		Content:            "<p>&quot;I lost my inheritance with one wrong digit on my sort code&quot;</p><p><a href=\"https://www.theguardian.com/money/2019/dec/07/i-lost-my-193000-inheritance-with-one-wrong-digit-on-my-sort-code\" rel=\"nofollow noopener noreferrer\" target=\"_blank\"><span class=\"invisible\">https://www.</span><span class=\"ellipsis\">theguardian.com/money/2019/dec</span><span class=\"invisible\">/07/i-lost-my-193000-inheritance-with-one-wrong-digit-on-my-sort-code</span}</p>",
		Reblog:             nil,
		Application: models.Application{
			Name:    "Frens",
			Website: "https://frens.com",
		},
		Account: models.Account{
			ID:           1,
			Username:     "bwoff11",
			Acct:         "bwoff11",
			DisplayName:  "Brent Wofford",
			Locked:       false,
			Bot:          false,
			Discoverable: true,
		},
		Card: models.Card{
			URL:          "https://example.com/",
			Title:        "Example",
			Description:  "This is an example of a card.",
			Type:         "summary",
			AuthorName:   "Example",
			AuthorURL:    "https://example.com/",
			ProviderName: "Example",
			ProviderURL:  "https://example.com/",
			HTML:         "<p>This is an example of a card.</p>",
			Width:        300,
			Height:       300,
			EmbedURL:     "https://example.com/",
		},
	})

	return c.JSON(resp)
}
