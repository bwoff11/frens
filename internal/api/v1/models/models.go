package models

type Emoji struct {
}

type Account struct {
	ID           uint64  `gorm:"primary_key"`
	Username     string  `gorm:"not null;unique"`
	Acct         string  `gorm:"not null;unique"`
	URL          string  `gorm:"not null"`
	DisplayName  string  `gorm:"not null"`
	Note         string  `gorm:"not null"`
	Avatar       string  `gorm:"not null"`
	AvatarStatic string  `gorm:"not null"`
	Header       string  `gorm:"not null"`
	HeaderStatic string  `gorm:"not null"`
	Locked       bool    `gorm:"not null"`
	Emojis       []Emoji `gorm:"many2many:emojis"`
	Discoverable bool    `gorm:"not null"`
}
