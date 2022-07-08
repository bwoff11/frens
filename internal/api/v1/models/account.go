package models

import "time"

type Account struct {
	// The account idheader.
	ID uint64 `gorm:"primary_key" json:"id,omitempty"`
	// The username of the account, not including domain.
	Username string `gorm:"not null;unique" json:"username,omitempty"`
	// The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	Acct string `gorm:"not null;unique" json:"acct,omitempty"`
	// The location of the user's profile page.
	URL string `gorm:"not null" json:"url,omitempty"`
	// The profile's display name.
	DisplayName string `gorm:"not null" json:"display_name,omitempty"`
	// The profile's bio / description.
	Note string `gorm:"not null" json:"note,omitempty"`
	// An image icon that is shown next to statuses and in the profile.
	Avatar string `gorm:"not null" json:"avatar,omitempty"`
	// A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	AvatarStatic string `gorm:"not null" json:"avatar_static,omitempty"`
	// An image banner that is shown above the profile and in profile cards.
	Header string `gorm:"not null" json:"header,omitempty"`
	// A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	HeaderStatic string `gorm:"not null" json:"header_static,omitempty"`
	// Whether the account manually approves follow requests.
	Locked bool `gorm:"not null" json:"locked,omitempty"`
	// Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	Emojis []Emoji `gorm:"many2many:emojis" json:"emojis,omitempty"`
	// Whether the account has opted into discovery features such as the profile directory.
	Discoverable bool `gorm:"not null" json:"discoverable,omitempty"`
	// When the account was created.
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	// When the most recent status was posted.
	LastStatusAt time.Time `gorm:"not null" json:"last_status_at,omitempty"`
	// How many statuses are attached to this account.
	StatusCount int `gorm:"not null" json:"status_count,omitempty"`
	// The reported followers of this profile.
	FollowersCount int `gorm:"not null" json:"followers_count,omitempty"`
	// The reported follows of this profile.
	FollowingCount int `gorm:"not null" json:"following_count,omitempty"`
	// Indicates that the profile is currently inactive and that its user has moved to a new account.
	// Moved Account `gorm:"not null" json:"moved,omitempty"`
	// Additional metadata attached to a profile as name-value pairs.
	Fields []Field `gorm:"many2many:fields" json:"fields,omitempty"`
	// A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Bot bool `gorm:"not null" json:"bot,omitempty"`
	// An extra entity to be used with API methods to verify credentials and update credentials.
	Source Source `gorm:"not null" json:"source,omitempty"`
	// An extra entity returned when an account is suspended.
	//Suspended Account `gorm:"not null" json:"suspended,omitempty"`
	// When a timed mute will expire, if applicable.
	MuteExpiresAt time.Time `gorm:"not null" json:"mute_expires_at,omitempty"`
}
