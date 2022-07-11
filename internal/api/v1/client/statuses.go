package client

import (
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

func addStatusRoutes(app *fiber.App) {
	sts := app.Group("/api/v1/statuses")            // Statuses
	sts.Post("/", createStatus)                     // /api/v1/statuses POST
	sts.Get("/:id", getStatus)                      // /api/v1/statuses/:id GET
	sts.Delete("/:id", deleteStatus)                // /api/v1/statuses/:id DELETE
	sts.Get("/:id/context", getStatusContext)       // /api/v1/statuses/:id/context GET
	sts.Get("/:id/favorited_by", getFavoritedBy)    // /api/v1/statuses/:id/favorited_by GET
	sts.Post("/:id/favourite", favouriteStatus)     // /api/v1/statuses/:id/favourite POST
	sts.Post("/:id/unfavourite", unfavouriteStatus) // /api/v1/statuses/:id/unfavourite POST
	sts.Post("/:id/reblog", reblogStatus)           // /api/v1/statuses/:id/reblog POST
	sts.Post("/:id/unreblog", unreblogStatus)       // /api/v1/statuses/:id/unreblog POST
	sts.Post("/:id/bookmark", bookmarkStatus)       // /api/v1/statuses/:id/bookmark POST
	sts.Post("/:id/unbookmark", unbookmarkStatus)   // /api/v1/statuses/:id/unbookmark POST
	sts.Post("/:id/mute", muteStatus)               // /api/v1/statuses/:id/mute POST
	sts.Post("/:id/unmute", unmuteStatus)           // /api/v1/statuses/:id/unmute POST
	sts.Post("/:id/pin", pinStatus)                 // /api/v1/statuses/:id/pin POST
	sts.Post("/:id/unpin", unpinStatus)             // /api/v1/statuses/:id/unpin POST
}

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

// View statuses above and below this status in the thread.
func getStatusContext(c *fiber.Ctx) error {
	//id := c.Params("id")

	var context models.Context
	// to be implemented
	return c.Status(200).JSON(&context)
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
