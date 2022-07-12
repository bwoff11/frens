package client

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"gorm.io/gorm"
)

func AddRoutes(app *fiber.App) {
	addOAuthRoutes(app)
	addApplicationRoutes(app)
	addInstanceRoutes(app)

	app.Post("/api/v1/accounts", createAccount) // /api/v1/accounts POST

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	addStatusRoutes(app)
	addAccountRoutes(app)

	arp := app.Group("/admin/reports")                  // Admin - Reports
	arp.Get("/", getReports)                            // /api/v1/admin/reports GET
	arp.Get("/:id", getReport)                          // /api/v1/admin/reports/:id GET
	arp.Post("/:id/assign_to_self", assignReportToSelf) // /api/v1/admin/reports/:id/assign_to_self POST
	arp.Post("/:id/unassign", unassignReport)           // /api/v1/admin/reports/:id/unassign POST
	arp.Post("/:id/resolve", resolveReport)             // /api/v1/admin/reports/:id/resolve POST
	arp.Post("/:id/reopen", reopenReport)               // /api/v1/admin/reports/:id/reopen POST

	v1 := app.Group("/api/v1") // Base Routes

	act := v1.Group("/admin/accounts")           // Admin - Accounts
	act.Get("/", getAccounts)                    // /api/v1/admin/accounts GET
	act.Get("/:id", getAccount)                  // /api/v1/admin/accounts/:id GET
	act.Post("/:id/action", accountAction)       // /api/v1/admin/accounts/:id/action POST
	act.Post("/:id/approve", approveAccount)     // /api/v1/admin/accounts/:id/approve POST
	act.Post("/:id/reject", rejectAccount)       // /api/v1/admin/accounts/:id/reject POST
	act.Post("/:id/enable", enableAccount)       // /api/v1/admin/accounts/:id/enable POST
	act.Post("/:id/unsilence", unsilenceAccount) // /api/v1/admin/accounts/:id/unsilence POST
	act.Post("/:id/unsuspend", unsuspendAccount) // /api/v1/admin/accounts/:id/unsuspend POST

	ann := v1.Group("/announcements")                         // Announcements
	ann.Get("/", getAnnouncements)                            // /api/v1/announcements GET
	ann.Post("/:id/dismiss", dismissAnnouncement)             // /api/v1/announcements/:id/dismiss POST
	ann.Put("/:id/reactions/:name", reactToAnnouncement)      // /api/v1/announcements/:id/reactions/:name PUT
	ann.Delete("/:id/reactions/:name", unreactToAnnouncement) // /api/v1/announcements/:id/reactions/:name DELETE

	bkm := v1.Group("/bookmarks")  // Bookmarks
	bkm.Get("/", GetSelfBookmarks) // /api/v1/bookmarks GET

	blk := v1.Group("/blocks")  // Blocks
	blk.Get("/", GetSelfBlocks) // /api/v1/blocks GET

	cus := v1.Group("/custom_emojis") // Custom Emojis
	cus.Get("/", getCustomEmojis)     // /api/v1/custom_emojis GET

	cvs := v1.Group("/conversations")      // Conversations
	cvs.Get("/", getSelfConversations)     // /api/v1/conversations GET
	cvs.Delete("/:id", deleteConversation) // /api/v1/conversations/:id DELETE
	cvs.Post("/:id", markConversationRead) // /api/v1/conversations/:id POST

	dmb := v1.Group("/domain_blocks") // DomainBlocks
	dmb.Get("/", GetDomainBlocks)     // /api/v1/domain_blocks/ GET
	dmb.Post("/", BlockDomain)        // /api/v1/domain_blocks/ POST
	dmb.Delete("/:id", UnblockDomain) // /api/v1/domain_blocks/:id DELETE

	end := v1.Group("/endorsements")  // Endorsements
	end.Get("/", GetSelfEndorsements) // /api/v1/endorsements GET

	fvt := v1.Group("/favourites")  // Favourites
	fvt.Get("/", GetSelfFavourites) // /api/v1/favourites GET

	ftt := v1.Group("/featured_tags")                  // FeaturedTags
	ftt.Get("/", getFeaturedTags)                      // /api/v1/featured_tags/ GET
	ftt.Post("/", createFeaturedTag)                   // /api/v1/featured_tags/ POST
	ftt.Delete("/:id", deleteFeaturedTag)              // /api/v1/featured_tags/:id DELETE
	ftt.Get("/suggestions", getFeaturedTagSuggestions) // /api/v1/featured_tags/suggestions GET

	flt := v1.Group("/filters")      // Filters
	flt.Get("/", getFilters)         // /api/v1/filters/ GET
	flt.Get("/:id", getFilter)       // /api/v1/filters/:id GET
	flt.Post("/", createFilter)      // /api/v1/filters/ POST
	flt.Put("/:id", updateFilter)    // /api/v1/filters/:id PUT
	flt.Delete("/:id", deleteFilter) // /api/v1/filters/:id DELETE

	flr := v1.Group("/follow_requests")             // api/v1/follow_requests GET
	flr.Get("/", getFollowRequests)                 // api/v1/follow_requests GET
	flr.Post("/:id/authorize", acceptFollowRequest) // api/v1/follow_requests/:id/authorize POST
	flr.Post("/:id/reject", rejectFollowRequest)    // api/v1/follow_requests/:id/reject POST

	lst := v1.Group("/lists")                          // Lists
	lst.Get("/", getLists)                             // /api/v1/lists GET
	lst.Get("/:id", getList)                           // /api/v1/lists/:id GET
	lst.Post("/", createList)                          // /api/v1/lists POST
	lst.Put("/:id", updateList)                        // /api/v1/lists/:id PUT
	lst.Delete("/:id", deleteList)                     // /api/v1/lists/:id DELETE
	lst.Get("/:id/accounts", getListAccounts)          // /api/v1/lists/:id/accounts GET
	lst.Post("/:id/accounts", addAccountToList)        // /api/v1/lists/:id/accounts POST
	lst.Delete("/:id/accounts", removeAccountFromList) // /api/v1/lists/:id/accounts DELETE

	mda := v1.Group("/media")    // Media
	mda.Post("/", attachMedia)   // /api/v1/media POST
	mda.Get("/:id", getMedia)    // /api/v1/media GET
	mda.Put("/:id", updateMedia) // /api/v1/media/:id PUT

	mte := v1.Group("/mutes")  // Mutes
	mte.Get("/", GetSelfMutes) // /api/v1/mutes GET

	mrk := v1.Group("/markers")   // Markers
	mrk.Get("/", getSelfMarker)   // /api/v1/markers GET
	mrk.Post("/", saveSelfMarker) // /api/v1/markers POST

	ntf := v1.Group("/notifications")                 // Notifications
	ntf.Get("/", getSelfNotifications)                // /api/v1/notifications GET
	ntf.Get("/:id", getSelfNotification)              // /api/v1/notifications GET
	ntf.Post("/clear", clearSelfNotifications)        // /api/v1/notifications/clear POST
	ntf.Post("/:id/dismiss", dismissSelfNotification) // /api/v1/notifications/:id/dismiss POST

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
}

// https://github.com/go-gorm/gorm/blob/master/errors.go
func handleDatabaseError(c *fiber.Ctx, err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "record not found",
		})
	case gorm.ErrInvalidTransaction:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid transaction",
		})
	case gorm.ErrNotImplemented:
		return c.Status(http.StatusNotImplemented).JSON(fiber.Map{
			"error": "not implemented",
		})
	case gorm.ErrMissingWhereClause:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "missing where clause",
		})
	case gorm.ErrUnsupportedRelation:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "unsupported relation",
		})
	case gorm.ErrPrimaryKeyRequired:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "primary key required",
		})
	case gorm.ErrModelValueRequired:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "model value required",
		})
	case gorm.ErrInvalidData:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
		})
	case gorm.ErrUnsupportedDriver:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "unsupported driver",
		})
	case gorm.ErrRegistered:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "registered",
		})
	case gorm.ErrInvalidField:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid field",
		})
	case gorm.ErrEmptySlice:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "empty slice",
		})
	case gorm.ErrDryRunModeUnsupported:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "dry run mode unsupported",
		})
	case gorm.ErrInvalidDB:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "invalid db",
		})
	case gorm.ErrInvalidValue:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid value",
		})
	case gorm.ErrInvalidValueOfLength:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid value of length",
		})
	case gorm.ErrPreloadNotAllowed:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "preload not allowed",
		})
	default:
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
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
