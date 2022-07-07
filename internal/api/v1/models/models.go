package models

type Emoji struct {
}

type Account struct {
	ID           uint64  `gorm:"primary_key" json:"id,omitempty"`
	Username     string  `gorm:"not null;unique" json:"username,omitempty"`
	Acct         string  `gorm:"not null;unique" json:"acct,omitempty"`
	URL          string  `gorm:"not null" json:"url,omitempty"`
	DisplayName  string  `gorm:"not null" json:"display_name,omitempty"`
	Note         string  `gorm:"not null" json:"note,omitempty"`
	Avatar       string  `gorm:"not null" json:"avatar,omitempty"`
	AvatarStatic string  `gorm:"not null" json:"avatar_static,omitempty"`
	Header       string  `gorm:"not null" json:"header,omitempty"`
	HeaderStatic string  `gorm:"not null" json:"header_static,omitempty"`
	Locked       bool    `gorm:"not null" json:"locked,omitempty"`
	Emojis       []Emoji `gorm:"many2many:emojis" json:"emojis,omitempty"`
	Discoverable bool    `gorm:"not null" json:"discoverable,omitempty"`
}
