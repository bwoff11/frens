package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	// The account idheader.
	ID uint64 `gorm:"primary_key" json:"id,omitempty"`
	// The username of the account, not including domain.
	Username string `gorm:";unique" json:"username,omitempty"`
	// The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	Acct string `gorm:";unique" json:"acct,omitempty"`
	// The location of the user's profile page.
	URL string `gorm:"" json:"url,omitempty"`
	// The profile's display name.
	DisplayName string `gorm:"" json:"display_name,omitempty"`
	// The profile's bio / description.
	Note string `gorm:"" json:"note,omitempty"`
	// An image icon that is shown next to statuses and in the profile.
	Avatar string `gorm:"" json:"avatar,omitempty"`
	// A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	AvatarStatic string `gorm:"" json:"avatar_static,omitempty"`
	// An image banner that is shown above the profile and in profile cards.
	Header string `gorm:"" json:"header,omitempty"`
	// A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	HeaderStatic string `gorm:"" json:"header_static,omitempty"`
	// Whether the account manually approves follow requests.
	Locked bool `gorm:"" json:"locked,omitempty"`
	// Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	//Emojis         []Emoji   `gorm:"" json:"emojis,omitempty"`
	// Whether the account has opted into discovery features such as the profile directory.
	Discoverable bool `gorm:"" json:"discoverable,omitempty"`
	// When the account was created.
	CreatedAt time.Time `gorm:"" json:"created_at,omitempty"`
	// When the most recent status was posted.
	LastStatusAt time.Time `gorm:"" json:"last_status_at,omitempty"`
	// How many statuses are attached to this account.
	StatusCount int `gorm:"" json:"status_count,omitempty"`
	// The reported followers of this profile.
	FollowersCount int `gorm:"" json:"followers_count,omitempty"`
	// The reported follows of this profile.
	FollowingCount int `gorm:"" json:"following_count,omitempty"`
	// Indicates that the profile is currently inactive and that its user has moved to a new account.
	//Moved          Account   `gorm:"" json:"moved,omitempty"`
	// Additional metadata attached to a profile as name-value pairs.
	//Fields        []Field   `gorm:"" json:"fields,omitempty"`
	// A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Bot bool `gorm:"" json:"bot,omitempty"`
	// An extra entity to be used with API methods to verify credentials and update credentials.
	//Source        Source    `gorm:"" json:"source,omitempty"`
	// An extra entity returned when an account is suspended.
	Suspended bool `gorm:"" json:"suspended,omitempty"`
	// When a timed mute will expire, if applicable.
	MuteExpiresAt time.Time `gorm:"" json:"mute_expires_at,omitempty"`

	// The following are objects not part of the pubic API spec but do neet to be stored.
	// I should figure out how to separate these out into their own structs later.
	Password string `gorm:"" json:"password,omitempty"` // The account's password.
}
