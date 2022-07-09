package client

import (
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type PublicTimelineRequestBody struct {
	Local     bool   `json:"local"`      // Show only local statuses? Defaults to false.
	Remote    bool   `json:"remote"`     // Show only remote statuses? Defaults to false.
	OnlyMedia bool   `json:"only_media"` // Show only statuses with media attached? Defaults to false.
	MaxID     string `json:"max_id"`     // Return results older than this id
	SinceID   string `json:"since_id"`   // Return results newer than this id
	MinID     string `json:"min_id"`     // Return results immediately newer than this id
	Limit     int    `json:"limit"`      // Maximum number of results to return. Defaults to 20.
}

func getPublicTimeline(c *fiber.Ctx) error {
	var reqBody PublicTimelineRequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		// check back on this later
		//return err
	}

	var resp []models.Status
	return c.Status(200).JSON(resp)
}

type HashtagTimelineRequestBody struct {
	Local     bool   `json:"local"`      // Show only local statuses? Defaults to false.
	OnlyMedia bool   `json:"only_media"` // Show only statuses with media attached? Defaults to false.
	MaxID     string `json:"max_id"`     // Return results older than this id
	SinceID   string `json:"since_id"`   // Return results newer than this id
	MinID     string `json:"min_id"`     // Return results immediately newer than this id
	Limit     int    `json:"limit"`      // Maximum number of results to return. Defaults to 20.
}

func getHashtagTimeline(c *fiber.Ctx) error {
	var resp []models.Status
	return c.Status(200).JSON(resp)
}

type HomeTimelineRequestBody struct {
	MaxID   string `json:"max_id"`   // Return results older than this id
	SinceID string `json:"since_id"` // Return results newer than this id
	MinID   string `json:"min_id"`   // Return results immediately newer than this id
	Limit   int    `json:"limit"`    // Maximum number of results to return. Defaults to 20.
	Local   bool   `json:"local"`    // Show only local statuses? Defaults to false.
}

// View statuses from followed users.
func getHomeTimeline(c *fiber.Ctx) error {
	var reqBody HomeTimelineRequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.SendStatus(400)
	}

	// TODO: get ID from token
	var resp []models.Status
	return c.Status(200).JSON(resp)
}

func getListTimeline(c *fiber.Ctx) error {
	var resp []models.Status
	return c.Status(200).JSON(resp)
}
