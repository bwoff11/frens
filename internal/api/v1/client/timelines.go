package client

import (
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

func getPublicTimeline(c *fiber.Ctx) error {
	var req PublicTimelineRequest
	if err := c.BodyParser(&req); err != nil {
		// check back on this later
		//return err
	}

	var resp []models.Status
	return c.Status(200).JSON(resp)
}

func getHashtagTimeline(c *fiber.Ctx) error {
	var resp []models.Status
	return c.Status(200).JSON(resp)
}

func getHomeTimeline(c *fiber.Ctx) error {
	var resp []models.Status
	return c.Status(200).JSON(resp)
}

func getListTimeline(c *fiber.Ctx) error {
	var resp []models.Status
	return c.Status(200).JSON(resp)
}
