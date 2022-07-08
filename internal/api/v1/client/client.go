package client

import (
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")                   // Base Routes
	v1.Get("/blocks", GetSelfBlocks)             // /api/v1/blocks GET
	v1.Get("/bookmarks", GetSelfBookmarks)       // /api/v1/bookmarks GET
	v1.Get("/endorsements", GetSelfEndorsements) // /api/v1/endorsements GET
	v1.Get("/favourites", GetSelfFavourites)     // /api/v1/favourites GET
	v1.Get("/mutes", GetSelfMutes)               // /api/v1/mutes GET
	v1.Get("/preferences", GetSelfPreferences)   // /api/v1/preferences GET
	v1.Post("/reports", ReportUser)              // /api/v1/reports POST

	act := v1.Group("/admin/accounts")           // Admin - Accounts
	act.Get("/", GetAccounts)                    // /api/v1/admin/accounts GET
	act.Get("/:id", GetAccount)                  // /api/v1/admin/accounts/:id GET
	act.Post("/:id/action", AccountAction)       // /api/v1/admin/accounts/:id/action POST
	act.Post("/:id/approve", ApproveAccount)     // /api/v1/admin/accounts/:id/approve POST
	act.Post("/:id/reject", RejectAccount)       // /api/v1/admin/accounts/:id/reject POST
	act.Post("/:id/enable", EnableAccount)       // /api/v1/admin/accounts/:id/enable POST
	act.Post("/:id/unsilence", UnsilenceAccount) // /api/v1/admin/accounts/:id/unsilence POST
	act.Post("/:id/unsuspend", UnsuspendAccount) // /api/v1/admin/accounts/:id/unsuspend POST

	ath := app.Group("/auth/sign_in")
	ath.Static("/", "./public/auth")

	ftt := v1.Group("/featured_tags")                  // FeaturedTags
	ftt.Get("/", GetFeaturedTags)                      // /api/v1/featured_tags/ GET
	ftt.Post("/", CreateFeaturedTag)                   // /api/v1/featured_tags/ POST
	ftt.Delete("/:id", DeleteFeaturedTag)              // /api/v1/featured_tags/:id DELETE
	ftt.Get("/suggestions", GetFeaturedTagSuggestions) // /api/v1/featured_tags/suggestions GET

	dmb := v1.Group("/domain_blocks") // DomainBlocks
	dmb.Get("/", GetDomainBlocks)     // /api/v1/domain_blocks/ GET
	dmb.Post("/", BlockDomain)        // /api/v1/domain_blocks/ POST
	dmb.Delete("/:id", UnblockDomain) // /api/v1/domain_blocks/:id DELETE

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

	acc := v1.Group("/accounts")                        // Accounts
	acc.Post("/", Register)                             // /api/v1/accounts POST
	acc.Get("/verify_credentials", VerifyCredentials)   // /api/v1/accounts/verify_credentials GET
	acc.Patch("/update_credentials", UpdateCredentials) // /api/v1/accounts/update_credentials PATCH
	acc.Get("/relationships", GetRelationships)         // /api/v1/accounts/relationships GET
	acc.Get("/search", SearchUsers)                     // /api/v1/accounts/search GET
	acc.Get("/:id", GetUserInfo)                        // /api/v1/accounts/:id GET
	acc.Get("/:id/statuses", GetUserStatuses)           // /api/v1/accounts/:id/statuses GET
	acc.Get("/:id/followers", GetUserFollowers)         // /api/v1/accounts/:id/followers GET
	acc.Get("/:id/following", GetUserFollowing)         // /api/v1/accounts/:id/following GET
	acc.Get("/:id/featured_tags", GetUserFeaturedTags)  // /api/v1/accounts/:id/featured_tags GET
	acc.Get("/:id/lists", GetUserLists)                 // /api/v1/accounts/:id/lists GET
	acc.Post("/:id/follow", FollowUser)                 // /api/v1/accounts/:id/follow POST
	acc.Post("/:id/unfollow", UnfollowUser)             // /api/v1/accounts/:id/unfollow POST
	acc.Post("/:id/block", BlockUser)                   // /api/v1/accounts/:id/block POST
	acc.Post("/:id/unblock", UnblockUser)               // /api/v1/accounts/:id/unblock POST
	acc.Post("/:id/mute", MuteUser)                     // /api/v1/accounts/:id/mute POST
	acc.Post("/:id/unmute", UnmuteUser)                 // /api/v1/accounts/:id/unmute POST
	acc.Post("/:id/pin", PinUser)                       // /api/v1/accounts/:id/pin POST
	acc.Post("/:id/unpin", UnpinUser)                   // /api/v1/accounts/:id/unpin POST
	acc.Post("/:id/note", NoteUser)                     // /api/v1/accounts/:id/note POST

	sug := v1.Group("/suggestions")      // Suggestions
	sug.Get("/", GetSuggestions)         // /api/v1/suggestions/ GET
	sug.Delete("/:id", DeleteSuggestion) // /api/v1/suggestions/:id DELETE

	rpt := app.Group("/admin/reports")                  // Admin - Reports
	rpt.Get("/", GetReports)                            // /api/v1/admin/reports GET
	rpt.Get("/:id", GetReport)                          // /api/v1/admin/reports/:id GET
	rpt.Post("/:id/assign_to_self", AssignReportToSelf) // /api/v1/admin/reports/:id/assign_to_self POST
	rpt.Post("/:id/unassign", UnassignReport)           // /api/v1/admin/reports/:id/unassign POST
	rpt.Post("/:id/resolve", ResolveReport)             // /api/v1/admin/reports/:id/resolve POST
	rpt.Post("/:id/reopen", ReopenReport)               // /api/v1/admin/reports/:id/reopen POST

	aps := v1.Group("/apps")                             // Apps
	aps.Options("/", getAppOptions)                      // /api/v1/apps/ OPTIONS
	aps.Post("/", createApp)                             // /api/v1/apps/ POST
	aps.Get("/verify_credentials", verifyAppCredentials) // /api/v1/apps/verify_credentials GET

	oat := app.Group("/oauth")           // OAuth
	oat.Get("/authorize", authorizeUser) // /oauth/authorize GET
	oat.Post("/token", getToken)         // /oauth/token POST
	oat.Post("/revoke", revokeToken)     // /oauth/revoke POST

	ins := v1.Group("/instance")              // Instances
	ins.Get("/", getInstance)                 // /api/v1/instances GET
	ins.Get("/peers", getPeers)               // /api/v1/instances/peers GET
	ins.Get("/activity", getInstanceActivity) // /api/v1/instances/activity GET

	tml := v1.Group("/timelines")                // Timelines
	tml.Get("/public", getPublicTimeline)        // /api/v1/timelines/public GET
	tml.Get("/tag/:hashtag", getHashtagTimeline) // /api/v1/timelines/tag/:hashtag GET
	tml.Get("/home", getHomeTimeline)            // /api/v1/timelines/home GET
	tml.Get("/list/:ID", getListTimeline)        // /api/v1/timelines/list/:ID GET
}
