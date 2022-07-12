package client

import (
	"github.com/bwoff11/frens/internal/db"
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
	if err := c.QueryParser(&reqBody); err != nil {
		return c.Status(400).JSON(map[string]string{
			"error": "Invalid request body",
		})
	}

	var resp []models.Status
	if err := db.Postgres.Order("id desc").Limit(reqBody.Limit).Find(&resp).Error; err != nil {
		return handleDatabaseError(c, err)
	}
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

	// if config show self is true, include own posts

	return nil
}

func getListTimeline(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
