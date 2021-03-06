package v1

import (
	"log"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/utils"
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

	accountID, err := utils.GetAccountID(c)
	if err != nil {
		log.Println("Error getting account ID:", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	var newStatus = &db.Status{
		Content: reqBody.Status,
		//Visibility:  db.VisibilityPublic,
		Sensitive:   false,
		SpoilerText: "",
		//MediaAttachments:   []db.Attachment{},
		//Application: ,
		//Mentions:           []db.Mention{},
		//Tags:               []db.Tag{},
		//Emojis:             []db.Emoji{},
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
		AccountID:  *accountID,
	}

	// validate here
	if err := db.Postgres.Create(newStatus).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating status")
	} else {
		return c.Status(200).JSON(newStatus)
	}
}

func getStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var status db.Status
	if err := db.Postgres.Where("id = ?", id).First(&status).Error; err != nil {
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
