package client

import "github.com/gofiber/fiber/v2"

func register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	return c.Status(200).JSON(req)
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
