package client

import (
	"github.com/bwoff11/frens/internal/api/v1/models"
	"github.com/bwoff11/frens/internal/db"
	"github.com/gofiber/fiber/v2"
)

func register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	var account models.Account
	account.Username = req.Username

	// Insert new account into database
	if err := db.DB.Create(&account).Error; err != nil {
		return err
	}

	return c.Status(200).JSON(account)
}

func verifyCredentials(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Verified",
	})
}

func updateCredentials(c *fiber.Ctx) error {
	var req UpdateCredentialsRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return c.Status(200).JSON(req)
}

func getRelationships(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Relationships",
	})
}

func searchUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "SearchUsers",
	})
}

func getUserInfo(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserInfo",
	})
}

func getUserStatuses(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserStatuses",
	})
}

func getUserFollowers(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFollowers",
	})
}

func getUserFollowing(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFollowing",
	})
}

func getUserFeaturedTags(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFeaturedTags",
	})
}

func getUserLists(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserLists",
	})
}

func followUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FollowUser",
	})
}

func unfollowUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnfollowUser",
	})
}

func blockUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "BlockUser",
	})
}

func unblockUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnblockUser",
	})
}

func muteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "MuteUser",
	})
}

func unmuteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnmuteUser",
	})
}

func pinUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "PinUser",
	})
}

func unpinUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnpinUser",
	})
}

func noteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "NoteUser",
	})
}

func getDomainBlocks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DomainBlocks",
	})
}

func blockDomain(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "BlockDomain",
	})
}

func unblockDomain(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnblockDomain",
	})
}

func getFeaturedTags(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FeaturedTags",
	})
}

func createFeaturedTag(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFeaturedTag",
	})
}

func deleteFeaturedTag(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFeaturedTag",
	})
}

func getFeaturedTagSuggestions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FeaturedTagSuggestions",
	})
}

func getFilters(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Filters",
	})
}

func getFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Filter",
	})
}

func createFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "CreateFilter",
	})
}

func updateFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UpdateFilter",
	})
}

func deleteFilter(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteFilter",
	})
}

func getFollowRequests(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FollowRequests",
	})
}

func acceptFollowRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AcceptFollowRequest",
	})
}

func rejectFollowRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "RejectFollowRequest",
	})
}

func getSuggestions(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Suggestions",
	})
}

func deleteSuggestion(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "DeleteSuggestion",
	})
}
