package v1

import (
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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

func GetPublicTimeline(c *fiber.Ctx) error {
	//var reqBody PublicTimelineRequestBody
	//if err := c.QueryParser(&reqBody); err != nil {
	//	return c.Status(400).JSON(map[string]string{
	//		"error": "Invalid request body",
	//	})
	//}

	var resp []db.Status
	if err := db.Postgres.Preload("Account").Order("id desc").Limit(20).Find(&resp).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get public timeline")
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

func GetHashtagTimeline(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

type HomeTimelineRequestBody struct {
	MaxID   string `json:"max_id"`   // Return results older than this id
	SinceID string `json:"since_id"` // Return results newer than this id
	MinID   string `json:"min_id"`   // Return results immediately newer than this id
	Limit   int    `json:"limit"`    // Maximum number of results to return. Defaults to 20.
	Local   bool   `json:"local"`    // Show only local statuses? Defaults to false.
}

// View statuses from followed users.
func GetHomeTimeline(c *fiber.Ctx) error {
	id, err := utils.GetAccountID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Request body CAN be null, so I need to think about how to handle that.
	//var reqBody HomeTimelineRequestBody
	//if err := c.BodyParser(&reqBody); err != nil {
	//	return c.Status(400).JSON(map[string]string{
	//		"error": "Invalid request body",
	//	})
	//}

	// Find relationships
	var relationships []db.Relationship
	if err := db.Postgres.Where("source_account_id = ? AND following = true", id).Find(&relationships).Error; err != nil {
		logrus.Errorf("Failed to get relationships for account %d", id)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get relationships")
	}

	// Find statuses
	var resp []db.Status
	if err := db.Postgres.Preload("Account").Order("id desc").Limit(20).Find(&resp).Error; err != nil {
		logrus.Errorf("Failed to get statuses for account %d", id)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get home timeline")
	}
	return c.Status(200).JSON(resp)
}

func GetListTimeline(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
