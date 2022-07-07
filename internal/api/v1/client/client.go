package client

import "github.com/gofiber/fiber/v2"

func AddRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Get("/bookmarks", getSelfBookmarks)
	v1.Get("/favourites", getSelfFavourites)
	v1.Get("/mutes", getSelfMutes)
	v1.Get("/blocks", getSelfBlocks)
	v1.Get("/endorsements", getSelfEndorsements)
	v1.Get("/preferences", getSelfPreferences)
	v1.Post("/reports", reportUser)

	sug := v1.Group("/suggestions")
	sug.Get("/", getSuggestions)
	sug.Delete("/:id", deleteSuggestion)

	ftt := v1.Group("/featured_tags")
	ftt.Get("/", getFeaturedTags)
	ftt.Post("/", createFeaturedTag)
	ftt.Delete("/:id", deleteFeaturedTag)
	ftt.Get("/suggestions", getFeaturedTagSuggestions)

	dmb := v1.Group("/domain_blocks")
	dmb.Get("/", getDomainBlocks)
	dmb.Post("/", blockDomain)
	dmb.Delete("/:id", unblockDomain)

	flt := v1.Group("/filters")
	flt.Get("/", getFilters)
	flt.Get("/:id", getFilter)
	flt.Post("/", createFilter)
	flt.Put("/:id", updateFilter)
	flt.Delete("/:id", deleteFilter)

	flr := v1.Group("/follow_requests")
	flr.Get("/", getFollowRequests)
	flr.Post("/:id/authorize", acceptFollowRequest)
	flr.Post("/:id/reject", rejectFollowRequest)

	acc := v1.Group("account")
	acc.Post("/", register)
	acc.Get("/verify_credentials", verifyCredentials)
	acc.Patch("/update_credentials", updateCredentials)
	acc.Get("/relationships", getRelationships)
	acc.Get("/search", searchUsers)
	acc.Get("/:id", getUserInfo)
	acc.Get("/:id/statuses", getUserStatuses)
	acc.Get("/:id/followers", getUserFollowers)
	acc.Get("/:id/following", getUserFollowing)
	acc.Get("/:id/featured_tags", getUserFeaturedTags)
	acc.Get("/:id/lists", getUserLists)
	acc.Post("/:id/follow", followUser)
	acc.Post("/:id/unfollow", unfollowUser)
	acc.Post("/:id/block", blockUser)
	acc.Post("/:id/unblock", unblockUser)
	acc.Post("/:id/mute", muteUser)
	acc.Post("/:id/unmute", unmuteUser)
	acc.Post("/:id/pin", pinUser)
	acc.Post("/:id/unpin", unpinUser)
	acc.Post("/:id/note", noteUser)
}

func getSelfBookmarks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Bookmarks",
	})
}

func getSelfFavourites(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Favourites",
	})
}

func getSelfMutes(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Mutes",
	})
}

func getSelfBlocks(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Blocks",
	})
}

func getSelfEndorsements(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Endorsements",
	})
}

func getSelfPreferences(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Preferences",
	})
}

func reportUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Report",
	})
}
