package models

import (
	"time"
)

type Account struct {
	ID           int        `gorm:"primary_key" json:"id"`   // The account idheader.
	CreatedAt    time.Time  `gorm:"" json:"created_at"`      // When the account was created.
	UpdatedAt    *time.Time `gorm:"" json:"updated_at"`      // When the account was last updated.
	DeletedAt    *time.Time `gorm:"" json:"deleted_at"`      // When the account was deleted.
	Username     string     `gorm:";unique" json:"username"` // The username of the account, not including domain.
	Acct         string     `gorm:";unique" json:"acct"`     // The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	URL          string     `gorm:"" json:"url"`             // The location of the user's profile page.
	DisplayName  string     `gorm:"" json:"display_name"`    // The profile's display name.
	Note         string     `gorm:"" json:"note"`            // The profile's bio / description.
	Avatar       string     `gorm:"" json:"avatar"`          // An image icon that is shown next to statuses and in the profile.
	AvatarStatic string     `gorm:"" json:"avatar_static"`   // A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	Header       string     `gorm:"" json:"header"`          // An image banner that is shown above the profile and in profile cards.
	HeaderStatic string     `gorm:"" json:"header_static"`   // A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	Locked       bool       `gorm:"" json:"locked"`          // Whether the account manually approves follow requests.
	Discoverable bool       `gorm:"" json:"discoverable"`    // Whether the account has opted into discovery features such as the profile directory.
	Bot          bool       `gorm:"" json:"bot"`             // A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Suspended    bool       `gorm:"" json:"suspended"`       // An extra entity returned when an account is suspended.

	LastStatusAt  *time.Time `gorm:"" json:"last_status_at"`  // When the most recent status was posted.
	MuteExpiresAt *time.Time `gorm:"" json:"mute_expires_at"` // When a timed mute will expire, if applicable.

	//StatusCount   int        `gorm:"-" json:"status_count"`   // How many statuses are attached to this account.
	//FollowersCount int       `gorm:"-" json:"followers_count"` // The reported followers of this profile.
	//FollowingCount int       `gorm:"-" json:"following_count"` // The reported follows of this profile.

	//Moved          Account   `gorm:"" json:"moved"`           // Indicates that the profile is currently inactive and that its user has moved to a new account.
	//SourceID uint64 `gorm:"" json:"id"`     // The ID of the account that the profile was moved to.
	//Source   Source `gorm:"" json:"source"` // An extra entity to be used with API methods to verify credentials and update credentials.

	//EmojiIDs []uint64 `gorm:"" json:"emoji_ids"` // The IDs of the custom emoji that are attached to this account.
	//Emojis   []Emoji  `gorm:"" json:"emojis"`    // Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	//Fields []Field `gorm:"" json:"fields"` // Additional metadata attached to a profile as name-value pairs.

	Password string `gorm:"" json:"password"` // The account's password.
}
