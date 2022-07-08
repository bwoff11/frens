package models

import "time"

type Account struct {
	// The account idheader.
	ID             uint64    `gorm:"primary_key" json:"id,omitempty"`
	Username       string    `gorm:"not null;unique" json:"username,omitempty"` // The username of the account, not including domain.
	Acct           string    `gorm:"not null;unique" json:"acct,omitempty"`     // The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	URL            string    `gorm:"not null" json:"url,omitempty"`             // The location of the user's profile page.
	DisplayName    string    `gorm:"not null" json:"display_name,omitempty"`    // The profile's display name.
	Note           string    `gorm:"not null" json:"note,omitempty"`            // The profile's bio / description.
	Avatar         string    `gorm:"not null" json:"avatar,omitempty"`          // An image icon that is shown next to statuses and in the profile.
	AvatarStatic   string    `gorm:"not null" json:"avatar_static,omitempty"`   // A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	Header         string    `gorm:"not null" json:"header,omitempty"`          // An image banner that is shown above the profile and in profile cards.
	HeaderStatic   string    `gorm:"not null" json:"header_static,omitempty"`   // A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	Locked         bool      `gorm:"not null" json:"locked,omitempty"`          // Whether the account manually approves follow requests.
	Emojis         []Emoji   `gorm:"many2many:emojis" json:"emojis,omitempty"`  // Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	Discoverable   bool      `gorm:"not null" json:"discoverable,omitempty"`    // Whether the account has opted into discovery features such as the profile directory.
	CreatedAt      time.Time `gorm:"not null" json:"created_at,omitempty"`      // When the account was created.
	LastStatusAt   time.Time `gorm:"not null" json:"last_status_at,omitempty"`  // When the most recent status was posted.
	StatusCount    int       `gorm:"not null" json:"status_count,omitempty"`    // How many statuses are attached to this account.
	FollowersCount int       `gorm:"not null" json:"followers_count,omitempty"` // The reported followers of this profile.
	FollowingCount int       `gorm:"not null" json:"following_count,omitempty"` // The reported follows of this profile.
	//Moved          Account   `gorm:"not null" json:"moved,omitempty"`           // Indicates that the profile is currently inactive and that its user has moved to a new account.
	Fields []Field `gorm:"many2many:fields" json:"fields,omitempty"` // Additional metadata attached to a profile as name-value pairs.
	Bot    bool    `gorm:"not null" json:"bot,omitempty"`            // A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Source Source  `gorm:"not null" json:"source,omitempty"`         // An extra entity to be used with API methods to verify credentials and update credentials.
	//Suspended     Account   `gorm:"not null" json:"suspended,omitempty"`       // An extra entity returned when an account is suspended.
	MuteExpiresAt time.Time `gorm:"not null" json:"mute_expires_at,omitempty"` // When a timed mute will expire, if applicable.
}
