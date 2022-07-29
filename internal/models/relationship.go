package models

type Relationship struct {
	ID              int `json:"-"`
	SourceAccountID int `json:"-"`
	TargetAccountID int `json:"id"`

	Following  bool
	FollowedBy bool `gorm:"-"`
	Requested  bool `gorm:"-"` // Do you have a pending follow request for this user?

	Blocking  bool
	BlockedBy bool `gorm:"-"`

	Muting              bool
	MutingNotifications bool `gorm:"-"` // Are you muting notifications from this user?

	ShowingReblogs bool // Are you receiving this user's boosts in your home timeline?
	Notifying      bool // Have you enabled notifications for this user?
	DomainBlocking bool // Have you blocked this user's domain?
}
