package client

import (
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateStatusBody struct {
	Status string `json:"status" form:"status"` // Text content of the status. If media_ids is provided, this becomes optional. Attaching a poll is optional while status is provided.
	// media ids
	// poll opions
	// poll expires in
}

func createStatus(c *fiber.Ctx) error {
	var reqBody CreateStatusBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	var newStatus = &models.Status{
		Content: reqBody.Status,
		//Visibility:  models.VisibilityPublic,
		Sensitive:   false,
		SpoilerText: "",
		//MediaAttachments:   []models.Attachment{},
		//Application:        models.Application{},
		//Mentions:           []models.Mention{},
		//Tags:               []models.Tag{},
		//Emojis:             []models.Emoji{},
		//ReblogsCount:       0,
		//FavoritesCount:     0,
		//RepliesCount:       0,
		URL:                "",
		InReplyToID:        "",
		InReplyToAccountID: "",
		//Reblog:             nil,
		//Poll:               nil,
		//Card:               nil,
		Language:   "",
		Text:       "",
		Favourited: false,
		Reblogged:  false,
		Muted:      false,
		Bookmarked: false,
		Pinned:     false,
		AccountID:  1,
	}

	// validate here
	if err := db.DB.Create(newStatus).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating status")
	} else {
		return c.Status(200).JSON(newStatus)
	}
}

func getStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var status models.Status
	if err := db.DB.Where("id = ?", id).First(&status).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Status not found")
	} else {
		return c.Status(200).JSON(&status)
	}
}

func deleteStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getStatusContext(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getFavoritedBy(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func favouriteStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unfavouriteStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func reblogStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unreblogStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func bookmarkStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unbookmarkStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func muteStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unmuteStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func pinStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unpinStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
