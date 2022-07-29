package v1

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreateAccountRequestBody struct {
	Username  string `json:"username" form:"username" validate:"required,min=3,max=32,alphanumunicode"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required,min=8,max=64"`
	Agreement bool   `json:"agreement" form:"agreement" validate:"required"`
	Locale    string `json:"locale" form:"locale" validate:"required"`
	Reason    string `json:"reason" form:"reason" validate:""`
}

func CreateAccount(c *fiber.Ctx) error {
	var reqBody CreateAccountRequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	var newAccount = &db.Account{
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
		LastStatusAt:  nil,
		Bot:           false,
		Suspended:     false,
		MuteExpiresAt: nil,
		Password:      db.Sha256(reqBody.Password),
	}

	if err := db.Postgres.Create(newAccount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create account")
	} else {
		return c.Status(fiber.StatusCreated).SendString("Account created")
	}
}

func VerifyCredentials(c *fiber.Ctx) error {
	accountID, err := utils.GetAccountID(c)
	if err != nil {
		return err
	}
	var Account db.Account
	if err := db.Postgres.Where("id = ?", accountID).First(&Account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get account")
	} else {
		return c.JSON(Account)
	}
}

func UpdateCredentials(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func GetUserFeaturedTags(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

type GetRelationshipsRequest struct {
	IDs []int `json:"ids" form:"ids" validate:"required"`
}

func GetRelationships(c *fiber.Ctx) error {
	sourceID, err := utils.GetAccountID(c)
	if err != nil {
		logrus.Error("Get relationship request failed due to invalid source account ID")
		return err
	}

	idsToCheckStr := c.Query("id[]")
	log.Println(idsToCheckStr)
	if idsToCheckStr == "" {
		logrus.Error("Get relationship request failed due to missing IDs to check")
		return c.Status(fiber.StatusBadRequest).SendString("Missing ids")
	}

	idsToCheckStrArray := strings.Split(idsToCheckStr, ",")
	idsToCheck := make([]int, len(idsToCheckStrArray))
	for i, id := range idsToCheckStrArray {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			logrus.Error("Get relationship request failed due to invalid ID to check")
			return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		}
		idsToCheck[i] = idInt
	}

	var resp []db.Relationship
	if err := db.Postgres.Where("source_account_id = ?", sourceID).Where("target_account_id IN (?)", idsToCheck).Find(&resp).Error; err != nil {
		logrus.Error("Get relationship request failed due to database error")
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get relationships")
	} else {
		return c.JSON(resp)
	}
}

func SearchUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func GetUserInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	var account db.Account
	if err := db.Postgres.First(&account, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get account")
	} else {
		if err := account.CalculateCounts(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to calculate counts")
		} else {
			return c.Status(fiber.StatusOK).JSON(account)
		}
	}
}

func GetUserStatuses(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	var statuses []db.Status
	if err := db.Postgres.Where("account_id = ?", id).Find(&statuses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get statuses")
	} else {
		return c.JSON(statuses)
	}
}

func GetUserFollowers(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	var followerRelationships []db.Relationship
	if err := db.Postgres.Where("target_account_id = ?", id).Find(&followerRelationships).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get followers")
	}

	var followerIDs []int
	for _, relationship := range followerRelationships {
		followerIDs = append(followerIDs, relationship.SourceAccountID)
	}

	var followers []db.Account
	if err := db.Postgres.Where("id IN (?)", followerIDs).Find(&followers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get followers")
	} else {
		return c.JSON(followers)
	}
}

func GetUserFollowing(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	var followingRelationships []db.Relationship
	if err := db.Postgres.Where("source_account_id = ?", id).Find(&followingRelationships).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get following")
	}

	var followingIDs []int
	for _, relationship := range followingRelationships {
		followingIDs = append(followingIDs, relationship.TargetAccountID)
	}

	var following []db.Account
	if err := db.Postgres.Where("id IN (?)", followingIDs).Find(&following).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Get following")
	} else {
		return c.JSON(following)
	}
}

func GetUserLists(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func FollowUser(c *fiber.Ctx) error {
	sourceID, err := utils.GetAccountID(c)
	if err != nil {
		return err
	}

	targetIDStr := c.Params("id")
	if targetIDStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Try to find prexisting relationship
	var relationship db.Relationship
	if err := db.Postgres.Where("source_account_id = ? AND target_account_id = ?", sourceID, targetID).First(&relationship).Error; err != nil {
		// Create new relationship
		relationship = db.Relationship{
			SourceAccountID: *sourceID,
			TargetAccountID: targetID,
			Following:       true,
		}
		if err := db.Postgres.Create(&relationship).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create relationship")
		}
	} else {
		// Check if already following
		if relationship.Following {
			return c.Status(fiber.StatusBadRequest).SendString("Already following")
		}
		// Update existing relationship
		relationship.Following = true
		if err := db.Postgres.Save(&relationship).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update relationship")
		}
	}

	return c.Status(fiber.StatusOK).SendString("Relationship created")
}

func UnfollowUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func BlockUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func UnblockUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func MuteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func UnmuteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func PinUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func UnpinUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func NoteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
