package models

import (
	"time"

	"gorm.io/gorm"
)

type Visibility string

const (
	VisibilityPublic   Visibility = "public"
	VisibilityUnlisted Visibility = "unlisted"
	VisibilityPrivate  Visibility = "private"
	VisibilityDirect   Visibility = "direct"
)

type Status struct {
	gorm.Model
	ID                 uint64    `gorm:"primary_key"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	DeletedAt          time.Time `json:"deleted_at"`
	Content            string    `json:"content"`
	SpoilerText        string    `json:"spoiler_text"`
	Language           string    `json:"language"`
	Text               string    `json:"text"`
	URL                string    `json:"url"`
	InReplyToID        string    `json:"in_reply_to_id"`
	InReplyToAccountID string    `json:"in_reply_to_account_id"`

	//Visibility         Visibility  `json:"visibility"`
	//Application        Application `gorm:"foreignkey:ApplicationID" json:"application"`
	AccountID uint64   `json:"-"`
	Account   *Account `gorm:"foreignkey:AccountID" json:"account"`

	//MediaAttachments []Attachment `json:"media_attachments"`
	//Mentions []Mention `json:"mentions"`
	//Tags     []Tag     `json:"tags"`
	//Emojis   []Emoji   `json:"emojis"`

	//Reblog             *Status `json:"reblog"`
	//Poll               *Poll   `json:"poll"`
	//Card               *Card   `json:"card"`

	//ReblogsCount   int `json:"reblogs_count"`
	//FavoritesCount int `json:"favourites_count"`
	//RepliesCount   int `json:"replies_count"`

	Sensitive  bool `json:"sensitive"`
	Favourited bool `json:"favourited"`
	Reblogged  bool `json:"reblogged"`
	Muted      bool `json:"muted"`
	Bookmarked bool `json:"bookmarked"`
	Pinned     bool `json:"pinned"`
}
