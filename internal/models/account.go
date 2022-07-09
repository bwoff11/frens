package models

import "time"

type Account struct {
	// The account idheader.
	ID           uint64 `gorm:"primary_key" json:"id,omitempty"`   // The account idheader.
	Username     string `gorm:";unique" json:"username,omitempty"` // The username of the account, not including domain.
	Acct         string `gorm:";unique" json:"acct,omitempty"`     // The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	URL          string `gorm:"" json:"url,omitempty"`             // The location of the user's profile page.
	DisplayName  string `gorm:"" json:"display_name,omitempty"`    // The profile's display name.
	Note         string `gorm:"" json:"note,omitempty"`            // The profile's bio / description.
	Avatar       string `gorm:"" json:"avatar,omitempty"`          // An image icon that is shown next to statuses and in the profile.
	AvatarStatic string `gorm:"" json:"avatar_static,omitempty"`   // A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	Header       string `gorm:"" json:"header,omitempty"`          // An image banner that is shown above the profile and in profile cards.
	HeaderStatic string `gorm:"" json:"header_static,omitempty"`   // A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	Locked       bool   `gorm:"" json:"locked,omitempty"`          // Whether the account manually approves follow requests.
	//Emojis         []Emoji   `gorm:"" json:"emojis,omitempty"`          // Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	Discoverable   bool      `gorm:"" json:"discoverable,omitempty"`    // Whether the account has opted into discovery features such as the profile directory.
	CreatedAt      time.Time `gorm:"" json:"created_at,omitempty"`      // When the account was created.
	LastStatusAt   time.Time `gorm:"" json:"last_status_at,omitempty"`  // When the most recent status was posted.
	StatusCount    int       `gorm:"" json:"status_count,omitempty"`    // How many statuses are attached to this account.
	FollowersCount int       `gorm:"" json:"followers_count,omitempty"` // The reported followers of this profile.
	FollowingCount int       `gorm:"" json:"following_count,omitempty"` // The reported follows of this profile.
	//Moved          Account   `gorm:"" json:"moved,omitempty"`           // Indicates that the profile is currently inactive and that its user has moved to a new account.
	//Fields        []Field   `gorm:"" json:"fields,omitempty"`                  // Additional metadata attached to a profile as name-value pairs.
	Bot bool `gorm:"" json:"bot,omitempty"` // A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	//Source        Source    `gorm:"" json:"source,omitempty"`          // An extra entity to be used with API methods to verify credentials and update credentials.
	Suspended     bool      `gorm:"" json:"suspended,omitempty"`       // An extra entity returned when an account is suspended.
	MuteExpiresAt time.Time `gorm:"" json:"mute_expires_at,omitempty"` // When a timed mute will expire, if applicable.

	// The following are objects not part of the pubic API spec but do neet to be stored.
	// I should figure out how to separate these out into their own structs later.
	Password  string `gorm:"" json:"password,omitempty"`             // The account's password.
	Following []int  `gorm:"type:bigint" json:"following,omitempty"` // The accounts that this account follows.
}
