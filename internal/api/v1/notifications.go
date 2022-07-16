package v1

import (
	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/bwoff11/frens/internal/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func getSelfNotifications(c *fiber.Ctx) error {
	id, err := utils.GetAccountID(c)
	if err != nil {
		return err
	}

	var notifications []models.Notification
	if err := db.Postgres.Find(&notifications, "account_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.SendStatus(200)
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
	} else {
		return c.JSON(notifications)
	}

}

func getSelfNotification(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func clearSelfNotifications(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func dismissSelfNotification(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
