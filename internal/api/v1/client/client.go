package client

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	ath := app.Group("/auth")
	ath.Static("/sign_in", "./public/login.html")
	ath.Static("/login_style", "./public/login_styles.css")

	oat := app.Group("/oauth")           // OAuth
	oat.Get("/authorize", authorizeUser) // /oauth/authorize GET
	oat.Post("/token", getToken)         // /oauth/token POST
	oat.Post("/revoke", revokeToken)     // /oauth/revoke POST

	arp := app.Group("/admin/reports")                  // Admin - Reports
	arp.Get("/", GetReports)                            // /api/v1/admin/reports GET
	arp.Get("/:id", GetReport)                          // /api/v1/admin/reports/:id GET
	arp.Post("/:id/assign_to_self", AssignReportToSelf) // /api/v1/admin/reports/:id/assign_to_self POST
	arp.Post("/:id/unassign", UnassignReport)           // /api/v1/admin/reports/:id/unassign POST
	arp.Post("/:id/resolve", ResolveReport)             // /api/v1/admin/reports/:id/resolve POST
	arp.Post("/:id/reopen", ReopenReport)               // /api/v1/admin/reports/:id/reopen POST

	v1 := app.Group("/api/v1") // Base Routes

	acc := v1.Group("/accounts")                           // Accounts
	acc.Post("/", Register)                                // /api/v1/accounts POST
	acc.Get("/verify_credentials", VerifyCredentials)      // /api/v1/accounts/verify_credentials GET
	acc.Patch("/update_credentials", UpdateCredentials)    // /api/v1/accounts/update_credentials PATCH
	acc.Get("/:id", GetUserInfo)                           // /api/v1/accounts/:id GET
	acc.Get("/:id/statuses", GetUserStatuses)              // /api/v1/accounts/:id/statuses GET
	acc.Get("/:id/followers", GetUserFollowers)            // /api/v1/accounts/:id/followers GET
	acc.Get("/:id/following", GetUserFollowing)            // /api/v1/accounts/:id/following GET
	acc.Get("/:id/featured_tags", GetUserFeaturedTags)     // /api/v1/accounts/:id/featured_tags GET
	acc.Get("/:id/lists", GetUserLists)                    // /api/v1/accounts/:id/lists GET
	acc.Get("/:id/identity_proofs", getUserIdentityProofs) // /api/v1/accounts/:id/identity_proofs GET
	acc.Post("/:id/follow", FollowUser)                    // /api/v1/accounts/:id/follow POST
	acc.Post("/:id/unfollow", UnfollowUser)                // /api/v1/accounts/:id/unfollow POST
	acc.Post("/:id/block", BlockUser)                      // /api/v1/accounts/:id/block POST
	acc.Post("/:id/unblock", UnblockUser)                  // /api/v1/accounts/:id/unblock POST
	acc.Post("/:id/mute", MuteUser)                        // /api/v1/accounts/:id/mute POST
	acc.Post("/:id/unmute", UnmuteUser)                    // /api/v1/accounts/:id/unmute POST
	acc.Post("/:id/pin", PinUser)                          // /api/v1/accounts/:id/pin POST
	acc.Post("/:id/unpin", UnpinUser)                      // /api/v1/accounts/:id/unpin POST
	acc.Post("/:id/note", NoteUser)                        // /api/v1/accounts/:id/note POST
	acc.Get("/relationships", GetRelationships)            // /api/v1/accounts/relationships GET
	acc.Get("/search", SearchUsers)                        // /api/v1/accounts/search GET

	act := v1.Group("/admin/accounts")           // Admin - Accounts
	act.Get("/", getAccounts)                    // /api/v1/admin/accounts GET
	act.Get("/:id", getAccount)                  // /api/v1/admin/accounts/:id GET
	act.Post("/:id/action", accountAction)       // /api/v1/admin/accounts/:id/action POST
	act.Post("/:id/approve", approveAccount)     // /api/v1/admin/accounts/:id/approve POST
	act.Post("/:id/reject", rejectAccount)       // /api/v1/admin/accounts/:id/reject POST
	act.Post("/:id/enable", enableAccount)       // /api/v1/admin/accounts/:id/enable POST
	act.Post("/:id/unsilence", unsilenceAccount) // /api/v1/admin/accounts/:id/unsilence POST
	act.Post("/:id/unsuspend", unsuspendAccount) // /api/v1/admin/accounts/:id/unsuspend POST

	aps := v1.Group("/apps")                             // Apps
	aps.Post("/", createApp)                             // /api/v1/apps/ POST
	aps.Get("/verify_credentials", verifyAppCredentials) // /api/v1/apps/verify_credentials GET

	bkm := v1.Group("/bookmarks")  // Bookmarks
	bkm.Get("/", GetSelfBookmarks) // /api/v1/bookmarks GET

	blk := v1.Group("/blocks")  // Blocks
	blk.Get("/", GetSelfBlocks) // /api/v1/blocks GET

	dmb := v1.Group("/domain_blocks") // DomainBlocks
	dmb.Get("/", GetDomainBlocks)     // /api/v1/domain_blocks/ GET
	dmb.Post("/", BlockDomain)        // /api/v1/domain_blocks/ POST
	dmb.Delete("/:id", UnblockDomain) // /api/v1/domain_blocks/:id DELETE

	end := v1.Group("/endorsements")  // Endorsements
	end.Get("/", GetSelfEndorsements) // /api/v1/endorsements GET

	fvt := v1.Group("/favourites")  // Favourites
	fvt.Get("/", GetSelfFavourites) // /api/v1/favourites GET

	ftt := v1.Group("/featured_tags")                  // FeaturedTags
	ftt.Get("/", GetFeaturedTags)                      // /api/v1/featured_tags/ GET
	ftt.Post("/", CreateFeaturedTag)                   // /api/v1/featured_tags/ POST
	ftt.Delete("/:id", DeleteFeaturedTag)              // /api/v1/featured_tags/:id DELETE
	ftt.Get("/suggestions", GetFeaturedTagSuggestions) // /api/v1/featured_tags/suggestions GET

	flt := v1.Group("/filters")      // Filters
	flt.Get("/", GetFilters)         // /api/v1/filters/ GET
	flt.Get("/:id", GetFilter)       // /api/v1/filters/:id GET
	flt.Post("/", CreateFilter)      // /api/v1/filters/ POST
	flt.Put("/:id", UpdateFilter)    // /api/v1/filters/:id PUT
	flt.Delete("/:id", DeleteFilter) // /api/v1/filters/:id DELETE

	flr := v1.Group("/follow_requests")             // api/v1/follow_requests GET
	flr.Get("/", GetFollowRequests)                 // api/v1/follow_requests GET
	flr.Post("/:id/authorize", AcceptFollowRequest) // api/v1/follow_requests/:id/authorize POST
	flr.Post("/:id/reject", RejectFollowRequest)    // api/v1/follow_requests/:id/reject POST

	mda := v1.Group("/media")    // Media
	mda.Post("/", attachMedia)   // /api/v1/media POST
	mda.Get("/:id", getMedia)    // /api/v1/media GET
	mda.Put("/:id", updateMedia) // /api/v1/media/:id PUT

	mte := v1.Group("/mutes")  // Mutes
	mte.Get("/", GetSelfMutes) // /api/v1/mutes GET

	pol := v1.Group("/polls") // Polls
	pol.Get("/:id", viewPoll) // /api/v1/polls GET
	pol.Post("/", voteOnPoll) // /api/v1/polls POST

	prf := v1.Group("/preferences")  // Preferences
	prf.Get("/", GetSelfPreferences) // /api/v1/preferences GET

	rpt := v1.Group("/reports") // Reports
	rpt.Get("/", ReportUser)    // /api/v1/reports GET

	sug := v1.Group("/suggestions")      // Suggestions
	sug.Get("/", GetSuggestions)         // /api/v1/suggestions/ GET
	sug.Delete("/:id", DeleteSuggestion) // /api/v1/suggestions/:id DELETE

	ins := v1.Group("/instance")              // Instances
	ins.Get("/", getInstance)                 // /api/v1/instances GET
	ins.Get("/peers", getPeers)               // /api/v1/instances/peers GET
	ins.Get("/activity", getInstanceActivity) // /api/v1/instances/activity GET

	tml := v1.Group("/timelines")                // Timelines
	tml.Get("/public", getPublicTimeline)        // /api/v1/timelines/public GET
	tml.Get("/tag/:hashtag", getHashtagTimeline) // /api/v1/timelines/tag/:hashtag GET
	tml.Get("/home", getHomeTimeline)            // /api/v1/timelines/home GET
	tml.Get("/list/:ID", getListTimeline)        // /api/v1/timelines/list/:ID GET

	sch := v1.Group("/scheduled_statuses")    // ScheduledStatuses
	sch.Get("/", getScheduledStatuses)        // /api/v1/scheduled_statuses GET
	sch.Get("/:id", getScheduledStatus)       // /api/v1/scheduled_statuses/:id GET
	sch.Put("/:id", scheduleStatus)           // /api/v1/scheduled_statuses/:id PUT
	sch.Delete("/:id", deleteScheduledStatus) // /api/v1/scheduled_statuses/:id DELETE

	sts := v1.Group("/statuses")                    // Statuses
	sts.Post("/", createStatus)                     // /api/v1/statuses POST
	sts.Get("/:id", getStatus)                      // /api/v1/statuses/:id GET
	sts.Delete("/:id", deleteStatus)                // /api/v1/statuses/:id DELETE
	sts.Get("/:id/context", getStatusContext)       // /api/v1/statuses/:id/context GET
	sts.Get("/:id/favorited_by", getFavoritedBy)    // /api/v1/statuses/:id/favorited_by GET
	sts.Post("/:id/favourite", favouriteStatus)     // /api/v1/statuses/:id/favourite POST
	sts.Post("/:id/unfavourite", unfavouriteStatus) // /api/v1/statuses/:id/unfavourite POST
	sts.Post("/:id/reblog", reblogStatus)           // /api/v1/statuses/:id/reblog POST
	sts.Post("/:id/unreblog", unreblogStatus)       // /api/v1/statuses/:id/unreblog POST
	sts.Post("/:id/bookmark", bookmarkStatus)       // /api/v1/statuses/:id/bookmark POST
	sts.Post("/:id/unbookmark", unbookmarkStatus)   // /api/v1/statuses/:id/unbookmark POST
	sts.Post("/:id/mute", muteStatus)               // /api/v1/statuses/:id/mute POST
	sts.Post("/:id/unmute", unmuteStatus)           // /api/v1/statuses/:id/unmute POST
	sts.Post("/:id/pin", pinStatus)                 // /api/v1/statuses/:id/pin POST
	sts.Post("/:id/unpin", unpinStatus)             // /api/v1/statuses/:id/unpin POST
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
