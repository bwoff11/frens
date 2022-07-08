package client

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserFeaturedTags(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFeaturedTags",
	})
}

func Register(c *fiber.Ctx) error {
	return c.Status(200).JSON("desu")
}

func VerifyCredentials(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Verified",
	})
}

func UpdateCredentials(c *fiber.Ctx) error {
	return c.Status(200).JSON("desu")
}

func GetRelationships(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Relationships",
	})
}

func SearchUsers(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "SearchUsers",
	})
}

func GetUserInfo(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserInfo",
	})
}

func GetUserStatuses(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserStatuses",
	})
}

func GetUserFollowers(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFollowers",
	})
}

func GetUserFollowing(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFollowing",
	})
}

func GetUserLists(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserLists",
	})
}

func FollowUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "FollowUser",
	})
}

func UnfollowUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnfollowUser",
	})
}

func BlockUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "BlockUser",
	})
}

func UnblockUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnblockUser",
	})
}

func MuteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "MuteUser",
	})
}

func UnmuteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnmuteUser",
	})
}

func PinUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "PinUser",
	})
}

func UnpinUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnpinUser",
	})
}

func NoteUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "NoteUser",
	})
}

func GetSelfBookmarks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Bookmarks",
	})
}

func GetSelfFavourites(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Favourites",
	})
}

func GetSelfMutes(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Mutes",
	})
}

func GetSelfBlocks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Blocks",
	})
}

func GetSelfEndorsements(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Endorsements",
	})
}

func GetSelfPreferences(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Preferences",
	})
}

func ReportUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Report",
	})
}
