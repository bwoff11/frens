package v1

import (
	v2 "github.com/bwoff11/frens/internal/api/v2"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AddRoutes(app *fiber.App) {
	addPublicRoutes(app)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	addAuthenticatedRoutes(app)
	addAdminRoutes(app)
}

func addPublicRoutes(app *fiber.App) {
	app.Post("/api/v1/accounts", CreateAccount)            // /api/v1/accounts POST
	app.Get("/api/v1/timelines/public", GetPublicTimeline) // /api/v1/timelines/public GET
	app.Get("/api/v2/search/", v2.Search)                  // /api/v2/search GET

	aps := app.Group("/api/v1/apps")                     // Apps
	aps.Post("/", createApp)                             // /api/v1/apps/ POST
	aps.Get("/verify_credentials", verifyAppCredentials) // /api/v1/apps/verify_credentials GET

	ins := app.Group("/api/v1/instance")      // Instances
	ins.Get("/", getInstance)                 // /api/v1/instances GET
	ins.Get("/peers", getPeers)               // /api/v1/instances/peers GET
	ins.Get("/activity", getInstanceActivity) // /api/v1/instances/activity GET

	oat := app.Group("/oauth")                // OAuth
	oat.Get("/authorize", GetLoginPage)       // /oauth/authorize GET
	oat.Post("/authorize", Login)             // /oauth/authorize POST
	oat.Get("/login_style.css", LoginStyle)   // /oauth/login_style.css GET
	oat.Get("/signup.html", SignupPage)       // /oauth/signup.html GET
	oat.Get("/signup_style.css", SignupStyle) // /oauth/signup_style.css GET
	oat.Post("/token", GetToken)              // /oauth/token POST
	oat.Post("/revoke", RevokeToken)          // /oauth/revoke POST
}

func addAuthenticatedRoutes(app *fiber.App) {
	app.Get("/api/v1/blocks/", GetSelfBlocks)               // /api/v1/blocks GET
	app.Get("/api/v1/bookmarks/", GetSelfBookmarks)         // /api/v1/bookmarks GET
	app.Get("/api/v1/mutes", GetSelfMutes)                  // /api/v1/mutes GET
	app.Get("/api/v1/suggestions", GetSuggestions)          // /api/v1/suggestions/ GET
	app.Delete("/api/v1/suggestions/:id", DeleteSuggestion) // /api/v1/suggestions/:id DELETE
	app.Get("/api/v1/preferences", GetSelfPreferences)      // /api/v1/preferences GET
	app.Get("/api/v1/custom_emojis", getCustomEmojis)       // /api/v1/custom_emojis GET
	app.Get("/api/v1/reports", ReportUser)                  // /api/v1/reports GET
	app.Get("/api/v1/polls/:id", viewPoll)                  // /api/v1/polls GET
	app.Post("/api/v1/polls/", voteOnPoll)                  // /api/v1/polls POST

	ntf := app.Group("/api/v1/notifications")         // Notifications
	ntf.Get("/", getSelfNotifications)                // /api/v1/notifications GET
	ntf.Post("/clear", clearSelfNotifications)        // /api/v1/notifications/clear POST
	ntf.Get("/:id", getSelfNotification)              // /api/v1/notifications GET
	ntf.Post("/:id/dismiss", dismissSelfNotification) // /api/v1/notifications/:id/dismiss POST

	acc := app.Group("/api/v1/accounts")                // Accounts
	acc.Get("/verify_credentials", VerifyCredentials)   // /api/v1/accounts/verify_credentials GET
	acc.Patch("/update_credentials", UpdateCredentials) // /api/v1/accounts/update_credentials PATCH
	acc.Get("/:id/statuses", GetUserStatuses)           // /api/v1/accounts/:id/statuses GET
	acc.Get("/:id/followers", GetUserFollowers)         // /api/v1/accounts/:id/followers GET
	acc.Get("/:id/following", GetUserFollowing)         // /api/v1/accounts/:id/following GET
	acc.Get("/:id/featured_tags", GetUserFeaturedTags)  // /api/v1/accounts/:id/featured_tags GET
	acc.Get("/:id/lists", GetUserLists)                 // /api/v1/accounts/:id/lists GET
	// Todo: The following isnt in the right place for some reason
	//acc.Get("/:id/identity_proofs", GetUserIdentityProofs) // /api/v1/accounts/:id/identity_proofs GET
	acc.Post("/:id/follow", FollowUser)         // /api/v1/accounts/:id/follow POST
	acc.Post("/:id/unfollow", UnfollowUser)     // /api/v1/accounts/:id/unfollow POST
	acc.Post("/:id/block", BlockUser)           // /api/v1/accounts/:id/block POST
	acc.Post("/:id/unblock", UnblockUser)       // /api/v1/accounts/:id/unblock POST
	acc.Post("/:id/mute", MuteUser)             // /api/v1/accounts/:id/mute POST
	acc.Post("/:id/unmute", UnmuteUser)         // /api/v1/accounts/:id/unmute POST
	acc.Post("/:id/pin", PinUser)               // /api/v1/accounts/:id/pin POST
	acc.Post("/:id/unpin", UnpinUser)           // /api/v1/accounts/:id/unpin POST
	acc.Post("/:id/note", NoteUser)             // /api/v1/accounts/:id/note POST
	acc.Get("/relationships", GetRelationships) // /api/v1/accounts/relationships GET
	//acc.Get("/search", RearchUsers)             // /api/v1/accounts/search GET
	acc.Get("/:id", GetUserInfo) // /api/v1/accounts/:id GET

	sts := app.Group("/api/v1/statuses")            // Statuses
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

	tml := app.Group("/api/v1/timelines")        // Timelines
	tml.Get("/tag/:hashtag", GetHashtagTimeline) // /api/v1/timelines/tag/:hashtag GET
	tml.Get("/home", GetHomeTimeline)            // /api/v1/timelines/home GET
	tml.Get("/list/:ID", GetListTimeline)        // /api/v1/timelines/list/:ID GET

	act := app.Group("/api/v1/admin/accounts")   // Admin - Accounts
	act.Get("/", GetAccounts)                    // /api/v1/admin/accounts GET
	act.Get("/:id", GetAccount)                  // /api/v1/admin/accounts/:id GET
	act.Post("/:id/action", AccountAction)       // /api/v1/admin/accounts/:id/action POST
	act.Post("/:id/approve", ApproveAccount)     // /api/v1/admin/accounts/:id/approve POST
	act.Post("/:id/reject", RejectAccount)       // /api/v1/admin/accounts/:id/reject POST
	act.Post("/:id/enable", EnableAccount)       // /api/v1/admin/accounts/:id/enable POST
	act.Post("/:id/unsilence", UnsilenceAccount) // /api/v1/admin/accounts/:id/unsilence POST
	act.Post("/:id/unsuspend", UnsuspendAccount) // /api/v1/admin/accounts/:id/unsuspend POST

	ann := app.Group("/api/v1/announcements")                 // Announcements
	ann.Get("/", GetAnnouncements)                            // /api/v1/announcements GET
	ann.Post("/:id/dismiss", DismissAnnouncement)             // /api/v1/announcements/:id/dismiss POST
	ann.Put("/:id/reactions/:name", ReactToAnnouncement)      // /api/v1/announcements/:id/reactions/:name PUT
	ann.Delete("/:id/reactions/:name", UnreactToAnnouncement) // /api/v1/announcements/:id/reactions/:name DELETE

	cvs := app.Group("/api/v1/conversations") // Conversations
	cvs.Get("/", getSelfConversations)        // /api/v1/conversations GET
	cvs.Delete("/:id", deleteConversation)    // /api/v1/conversations/:id DELETE
	cvs.Post("/:id", markConversationRead)    // /api/v1/conversations/:id POST

	dmb := app.Group("/api/v1/domain_blocks") // DomainBlocks
	dmb.Get("/", GetDomainBlocks)             // /api/v1/domain_blocks/ GET
	dmb.Post("/", BlockDomain)                // /api/v1/domain_blocks/ POST
	dmb.Delete("/:id", UnblockDomain)         // /api/v1/domain_blocks/:id DELETE

	end := app.Group("/api/v1/endorsements") // Endorsements
	end.Get("/", GetSelfEndorsements)        // /api/v1/endorsements GET

	fvt := app.Group("/api/v1/favourites") // Favourites
	fvt.Get("/", GetSelfFavourites)        // /api/v1/favourites GET

	ftt := app.Group("/api/v1/featured_tags")          // FeaturedTags
	ftt.Get("/", getFeaturedTags)                      // /api/v1/featured_tags/ GET
	ftt.Post("/", createFeaturedTag)                   // /api/v1/featured_tags/ POST
	ftt.Delete("/:id", deleteFeaturedTag)              // /api/v1/featured_tags/:id DELETE
	ftt.Get("/suggestions", getFeaturedTagSuggestions) // /api/v1/featured_tags/suggestions GET

	flt := app.Group("/api/v1/filters") // Filters
	flt.Get("/", getFilters)            // /api/v1/filters/ GET
	flt.Get("/:id", getFilter)          // /api/v1/filters/:id GET
	flt.Post("/", createFilter)         // /api/v1/filters/ POST
	flt.Put("/:id", updateFilter)       // /api/v1/filters/:id PUT
	flt.Delete("/:id", deleteFilter)    // /api/v1/filters/:id DELETE

	flr := app.Group("/api/v1/follow_requests")     // api/v1/follow_requests GET
	flr.Get("/", getFollowRequests)                 // api/v1/follow_requests GET
	flr.Post("/:id/authorize", acceptFollowRequest) // api/v1/follow_requests/:id/authorize POST
	flr.Post("/:id/reject", rejectFollowRequest)    // api/v1/follow_requests/:id/reject POST

	lst := app.Group("/api/v1/lists")                  // Lists
	lst.Get("/", getLists)                             // /api/v1/lists GET
	lst.Get("/:id", getList)                           // /api/v1/lists/:id GET
	lst.Post("/", createList)                          // /api/v1/lists POST
	lst.Put("/:id", updateList)                        // /api/v1/lists/:id PUT
	lst.Delete("/:id", deleteList)                     // /api/v1/lists/:id DELETE
	lst.Get("/:id/accounts", getListAccounts)          // /api/v1/lists/:id/accounts GET
	lst.Post("/:id/accounts", addAccountToList)        // /api/v1/lists/:id/accounts POST
	lst.Delete("/:id/accounts", removeAccountFromList) // /api/v1/lists/:id/accounts DELETE

	mda := app.Group("/api/v1/media") // Media
	mda.Post("/", attachMedia)        // /api/v1/media POST
	mda.Get("/:id", getMedia)         // /api/v1/media GET
	mda.Put("/:id", updateMedia)      // /api/v1/media/:id PUT

	mrk := app.Group("/api/v1/markers") // Markers
	mrk.Get("/", getSelfMarker)         // /api/v1/markers GET
	mrk.Post("/", saveSelfMarker)       // /api/v1/markers POST

	sch := app.Group("/api/v1/scheduled_statuses") // ScheduledStatuses
	sch.Get("/", getScheduledStatuses)             // /api/v1/scheduled_statuses GET
	sch.Get("/:id", getScheduledStatus)            // /api/v1/scheduled_statuses/:id GET
	sch.Put("/:id", scheduleStatus)                // /api/v1/scheduled_statuses/:id PUT
	sch.Delete("/:id", deleteScheduledStatus)      // /api/v1/scheduled_statuses/:id DELETE
}

func addAdminRoutes(app *fiber.App) {
	arp := app.Group("/admin/reports")                  // Admin - Reports
	arp.Get("/", GetReports)                            // /api/v1/admin/reports GET
	arp.Get("/:id", GetReport)                          // /api/v1/admin/reports/:id GET
	arp.Post("/:id/assign_to_self", AssignReportToSelf) // /api/v1/admin/reports/:id/assign_to_self POST
	arp.Post("/:id/unassign", UnassignReport)           // /api/v1/admin/reports/:id/unassign POST
	arp.Post("/:id/resolve", ResolveReport)             // /api/v1/admin/reports/:id/resolve POST
	arp.Post("/:id/reopen", ReopenReport)               // /api/v1/admin/reports/:id/reopen POST
}
