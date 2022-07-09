package client

import (
	"time"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

func getUserFeaturedTags(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

type CreateAccountRequestBody struct {
	Username  string `json:"username" form:"username"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Agreement bool   `json:"agreement" form:"agreement"`
	Reason    string `json:"reason" form:"reason"`
}

func createAccount(c *fiber.Ctx) error {
	var reqBody CreateAccountRequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	var newAccount = &models.Account{
		Username:     reqBody.Username,
		Acct:         reqBody.Username,
		DisplayName:  reqBody.Username,
		Note:         "",
		Avatar:       "",
		AvatarStatic: "",
		Header:       "",
		HeaderStatic: "",
		Locked:       false,
		//Emojis: "",
		Discoverable:  true,
		CreatedAt:     time.Now(),
		LastStatusAt:  time.Now(),
		Bot:           false,
		Suspended:     false,
		MuteExpiresAt: time.Now(),
		Password:      reqBody.Password, // salt this
	}

	if err := db.DB.Create(newAccount).Error; err != nil {
		return handleDatabaseError(c, err)
	} else {
		return c.Status(fiber.StatusCreated).SendString("Account created")
	}
}

func verifyCredentials(c *fiber.Ctx) error {
	return c.Status(200).SendString("Not implemented")
}

func updateCredentials(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getRelationships(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func searchUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getUserInfo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getUserStatuses(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getUserFollowers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getUserFollowing(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getUserLists(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func followUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unfollowUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func blockUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unblockUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func muteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unmuteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func pinUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func unpinUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func noteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
