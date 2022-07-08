package client

import (
	"github.com/bwoff11/frens/internal/api/v1/client/accounts"
	"github.com/bwoff11/frens/internal/api/v1/client/admin"
	"github.com/bwoff11/frens/internal/api/v1/client/apps"
	"github.com/bwoff11/frens/internal/api/v1/client/instance"
	"github.com/bwoff11/frens/internal/api/v1/client/notifications"
	"github.com/bwoff11/frens/internal/api/v1/client/oembed"
	"github.com/bwoff11/frens/internal/api/v1/client/proofs"
	"github.com/bwoff11/frens/internal/api/v1/client/search"
	"github.com/bwoff11/frens/internal/api/v1/client/statuses"
	"github.com/bwoff11/frens/internal/api/v1/client/timelines"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")                            // Base Routes
	v1.Get("/blocks", accounts.GetSelfBlocks)             // /api/v1/blocks GET
	v1.Get("/bookmarks", accounts.GetSelfBookmarks)       // /api/v1/bookmarks GET
	v1.Get("/endorsements", accounts.GetSelfEndorsements) // /api/v1/endorsements GET
	v1.Get("/favourites", accounts.GetSelfFavourites)     // /api/v1/favourites GET
	v1.Get("/mutes", accounts.GetSelfMutes)               // /api/v1/mutes GET
	v1.Get("/preferences", accounts.GetSelfPreferences)   // /api/v1/preferences GET
	v1.Post("/reports", accounts.ReportUser)              // /api/v1/reports POST

	ftt := v1.Group("/featured_tags")                           // FeaturedTags
	ftt.Get("/", accounts.GetFeaturedTags)                      // /api/v1/featured_tags/ GET
	ftt.Post("/", accounts.CreateFeaturedTag)                   // /api/v1/featured_tags/ POST
	ftt.Delete("/:id", accounts.DeleteFeaturedTag)              // /api/v1/featured_tags/:id DELETE
	ftt.Get("/suggestions", accounts.GetFeaturedTagSuggestions) // /api/v1/featured_tags/suggestions GET

	dmb := app.Group("/api/v1/domain_blocks")  // DomainBlocks
	dmb.Get("/", accounts.GetDomainBlocks)     // /api/v1/domain_blocks/ GET
	dmb.Post("/", accounts.BlockDomain)        // /api/v1/domain_blocks/ POST
	dmb.Delete("/:id", accounts.UnblockDomain) // /api/v1/domain_blocks/:id DELETE

	flt := app.Group("/api/v1/filters")       // Filters
	flt.Get("/", accounts.GetFilters)         // /api/v1/filters/ GET
	flt.Get("/:id", accounts.GetFilter)       // /api/v1/filters/:id GET
	flt.Post("/", accounts.CreateFilter)      // /api/v1/filters/ POST
	flt.Put("/:id", accounts.UpdateFilter)    // /api/v1/filters/:id PUT
	flt.Delete("/:id", accounts.DeleteFilter) // /api/v1/filters/:id DELETE

	flr := v1.Group("/follow_requests")                      // api/v1/follow_requests GET
	flr.Get("/", accounts.GetFollowRequests)                 // api/v1/follow_requests GET
	flr.Post("/:id/authorize", accounts.AcceptFollowRequest) // api/v1/follow_requests/:id/authorize POST
	flr.Post("/:id/reject", accounts.RejectFollowRequest)    // api/v1/follow_requests/:id/reject POST

	acc := v1.Group("/accounts")                                 // Accounts
	acc.Post("/", accounts.Register)                             // /api/v1/accounts POST
	acc.Get("/verify_credentials", accounts.VerifyCredentials)   // /api/v1/accounts/verify_credentials GET
	acc.Patch("/update_credentials", accounts.UpdateCredentials) // /api/v1/accounts/update_credentials PATCH
	acc.Get("/relationships", accounts.GetRelationships)         // /api/v1/accounts/relationships GET
	acc.Get("/search", accounts.SearchUsers)                     // /api/v1/accounts/search GET
	acc.Get("/:id", accounts.GetUserInfo)                        // /api/v1/accounts/:id GET
	acc.Get("/:id/statuses", accounts.GetUserStatuses)           // /api/v1/accounts/:id/statuses GET
	acc.Get("/:id/followers", accounts.GetUserFollowers)         // /api/v1/accounts/:id/followers GET
	acc.Get("/:id/following", accounts.GetUserFollowing)         // /api/v1/accounts/:id/following GET
	acc.Get("/:id/featured_tags", accounts.GetUserFeaturedTags)  // /api/v1/accounts/:id/featured_tags GET
	acc.Get("/:id/lists", accounts.GetUserLists)                 // /api/v1/accounts/:id/lists GET
	acc.Post("/:id/follow", accounts.FollowUser)                 // /api/v1/accounts/:id/follow POST
	acc.Post("/:id/unfollow", accounts.UnfollowUser)             // /api/v1/accounts/:id/unfollow POST
	acc.Post("/:id/block", accounts.BlockUser)                   // /api/v1/accounts/:id/block POST
	acc.Post("/:id/unblock", accounts.UnblockUser)               // /api/v1/accounts/:id/unblock POST
	acc.Post("/:id/mute", accounts.MuteUser)                     // /api/v1/accounts/:id/mute POST
	acc.Post("/:id/unmute", accounts.UnmuteUser)                 // /api/v1/accounts/:id/unmute POST
	acc.Post("/:id/pin", accounts.PinUser)                       // /api/v1/accounts/:id/pin POST
	acc.Post("/:id/unpin", accounts.UnpinUser)                   // /api/v1/accounts/:id/unpin POST
	acc.Post("/:id/note", accounts.NoteUser)                     // /api/v1/accounts/:id/note POST

	sug := app.Group("/api/v1/suggestions")       // Suggestions
	sug.Get("/", accounts.GetSuggestions)         // /api/v1/suggestions/ GET
	sug.Delete("/:id", accounts.DeleteSuggestion) // /api/v1/suggestions/:id DELETE

	act := app.Group("/api/v1/admin/accounts")                  // Admin - Accounts
	act.Get("/accounts", admin.GetAccounts)                     // /api/v1/admin/accounts GET
	act.Get("/accounts/:id", admin.GetAccount)                  // /api/v1/admin/accounts/:id GET
	act.Post("/accounts/:id/action", admin.AccountAction)       // /api/v1/admin/accounts/:id/action POST
	act.Post("/accounts/:id/approve", admin.ApproveAccount)     // /api/v1/admin/accounts/:id/approve POST
	act.Post("/accounts/:id/reject", admin.RejectAccount)       // /api/v1/admin/accounts/:id/reject POST
	act.Post("/accounts/:id/enable", admin.EnableAccount)       // /api/v1/admin/accounts/:id/enable POST
	act.Post("/accounts/:id/unsilence", admin.UnsilenceAccount) // /api/v1/admin/accounts/:id/unsilence POST
	act.Post("/accounts/:id/unsuspend", admin.UnsuspendAccount) // /api/v1/admin/accounts/:id/unsuspend POST

	rpt := app.Group("/api/v1/reports")                       // Admin - Reports
	rpt.Get("/", admin.GetReports)                            // /api/v1/admin/reports GET
	rpt.Get("/:id", admin.GetReport)                          // /api/v1/admin/reports/:id GET
	rpt.Post("/:id/assign_to_self", admin.AssignReportToSelf) // /api/v1/admin/reports/:id/assign_to_self POST
	rpt.Post("/:id/unassign", admin.UnassignReport)           // /api/v1/admin/reports/:id/unassign POST
	rpt.Post("/:id/resolve", admin.ResolveReport)             // /api/v1/admin/reports/:id/resolve POST
	rpt.Post("/:id/reopen", admin.ReopenReport)               // /api/v1/admin/reports/:id/reopen POST

	apps.AddRoutes(app)
	instance.AddRoutes(app)
	notifications.AddRoutes(app)
	oembed.AddRoutes(app)
	proofs.AddRoutes(app)
	search.AddRoutes(app)
	statuses.AddRoutes(app)
	timelines.AddRoutes(app)
}
