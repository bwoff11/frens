package main

import (
	"github.com/bwoff11/frens/internal/api/client"
	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	v1 := api.Group("/v1")

	acc := v1.Group("/accounts")
	acc.Post("/", client.Register)
	acc.Get("/verify_credentials", client.VerifyCredentials)
	acc.Patch("/update_credentials", client.UpdateCredentials)
	acc.Get("/relationships", client.GetRelationships)
	acc.Get("/search", client.SearchUsers)
	acc.Get("/:userID", client.GetUserInfo)
	acc.Get("/:userID/statuses", client.GetUserStatuses)
	acc.Get("/:userID/followers", client.GetUserFollowers)
	acc.Get("/:userID/following", client.GetUserFollowing)
	acc.Get("/:userID/featured_tags", client.GetUserFeaturedTags)
	acc.Get("/:userID/lists", client.GetUserLists)
	acc.Post("/:userID/follow", client.FollowUser)
	acc.Post("/:userID/unfollow", client.UnfollowUser)
	acc.Post("/:userID/block", client.BlockUser)
	acc.Post("/:userID/unblock", client.UnblockUser)
	acc.Post("/:userID/mute", client.MuteUser)
	acc.Post("/:userID/unmute", client.UnmuteUser)
	acc.Post("/:userID/pin", client.PinUser)
	acc.Post("/:userID/unpin", client.UnpinUser)
	acc.Post("/:userID/note", client.NoteUser)

	api.Listen(":8080")
}
