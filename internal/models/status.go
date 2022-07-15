package models

import (
	"time"
)

type Visibility string

const (
	VisibilityPublic   Visibility = "public"
	VisibilityUnlisted Visibility = "unlisted"
	VisibilityPrivate  Visibility = "private"
	VisibilityDirect   Visibility = "direct"
)

type Status struct {
	ID                 int        `json:"id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	Content            string     `json:"content"`
	SpoilerText        string     `json:"spoiler_text"`
	Language           string     `json:"language"`
	Text               string     `json:"text"`
	URL                string     `json:"url"`
	InReplyToID        string     `json:"in_reply_to_id"`
	InReplyToAccountID string     `json:"in_reply_to_account_id"`

	Visibility Visibility `json:"visibility"`
	//Application Application
	AccountID int     `json:"-"`
	Account   Account `json:"account"`

	//MediaAttachmentsIDs []uint64      `json:"-"`
	//MediaAttachments    []*Attachment `gorm:"has_many:media_attachments"`
	//Mentions            []Mention    `gorm:"has_many:mentions"`
	//Tags                []Tag        `gorm:"has_many:tags"`
	//Emojis              []Emoji      `gorm:"has_many:emojis"`

	//Reblog *Status
	//Poll   *Poll
	//Card   *Card

	ReblogsCount   int `gorm:"-" json:"reblogs_count"`
	FavoritesCount int `gorm:"-" json:"favourites_count"`
	RepliesCount   int `gorm:"-" json:"replies_count"`

	Sensitive  bool `json:"sensitive"`
	Favourited bool `json:"favourited"`
	Reblogged  bool `json:"reblogged"`
	Muted      bool `json:"muted"`
	Bookmarked bool `json:"bookmarked"`
	Pinned     bool `json:"pinned"`
}
