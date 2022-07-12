package client

import (
	"log"
	"time"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

func addAccountRoutes(app *fiber.App) {
	acc := app.Group("/api/v1/accounts")                // Accounts
	acc.Post("/", createAccount)                        // /api/v1/accounts POST
	acc.Get("/verify_credentials", verifyCredentials)   // /api/v1/accounts/verify_credentials GET
	acc.Patch("/update_credentials", updateCredentials) // /api/v1/accounts/update_credentials PATCH
	acc.Get("/:id", getUserInfo)                        // /api/v1/accounts/:id GET
	acc.Get("/:id/statuses", getUserStatuses)           // /api/v1/accounts/:id/statuses GET
	acc.Get("/:id/followers", getUserFollowers)         // /api/v1/accounts/:id/followers GET
	acc.Get("/:id/following", getUserFollowing)         // /api/v1/accounts/:id/following GET
	acc.Get("/:id/featured_tags", getUserFeaturedTags)  // /api/v1/accounts/:id/featured_tags GET
	acc.Get("/:id/lists", getUserLists)                 // /api/v1/accounts/:id/lists GET
	// Todo: The following isnt in the right place for some reason
	acc.Get("/:id/identity_proofs", getUserIdentityProofs) // /api/v1/accounts/:id/identity_proofs GET
	acc.Post("/:id/follow", followUser)                    // /api/v1/accounts/:id/follow POST
	acc.Post("/:id/unfollow", unfollowUser)                // /api/v1/accounts/:id/unfollow POST
	acc.Post("/:id/block", blockUser)                      // /api/v1/accounts/:id/block POST
	acc.Post("/:id/unblock", unblockUser)                  // /api/v1/accounts/:id/unblock POST
	acc.Post("/:id/mute", muteUser)                        // /api/v1/accounts/:id/mute POST
	acc.Post("/:id/unmute", unmuteUser)                    // /api/v1/accounts/:id/unmute POST
	acc.Post("/:id/pin", pinUser)                          // /api/v1/accounts/:id/pin POST
	acc.Post("/:id/unpin", unpinUser)                      // /api/v1/accounts/:id/unpin POST
	acc.Post("/:id/note", noteUser)                        // /api/v1/accounts/:id/note POST
	acc.Get("/relationships", getRelationships)            // /api/v1/accounts/relationships GET
	acc.Get("/search", searchUsers)                        // /api/v1/accounts/search GET
}

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

	password := reqBody.Password
	hashedPassword := db.Sha256(password)
	log.Println(password)
	log.Println(hashedPassword)

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
		Password:      db.Sha256(reqBody.Password),
	}

	if err := db.Postgres.Create(newAccount).Error; err != nil {
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
	id := c.Params("id")

	var account models.Account
	if err := db.Postgres.Where("id = ?", id).First(&account).Error; err != nil {
		return handleDatabaseError(c, err)
	} else {
		return c.JSON(account)
	}
}

func getUserStatuses(c *fiber.Ctx) error {
	id := c.Params("id")

	var statuses []models.Status
	if err := db.Postgres.Where("account_id = ?", id).Find(&statuses).Error; err != nil {
		return handleDatabaseError(c, err)
	} else {
		return c.JSON(statuses)
	}
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
