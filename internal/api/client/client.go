package client

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return c.Status(200).JSON(req)
}

func VerifyCredentials(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Verified",
	})
}

func UpdateCredentials(c *fiber.Ctx) error {
	var req UpdateCredentialsRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return c.Status(200).JSON(req)
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

func GetUserFeaturedTags(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UserFeaturedTags",
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
