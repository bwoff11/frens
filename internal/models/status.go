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

	Visibility    Visibility   `json:"visibility"`
	ApplicationID int          `json:"-"`
	Application   *Application `json:"application"`
	AccountID     int          `json:"-"`
	Account       *Account     `json:"account"`
	//MediaAttachmentsIDs []uint64      `json:"-"`
	//MediaAttachments    []*Attachment `json:"media_attachments"`
	//MendionsIDs []int     `json:"-"`
	//Mentions    []Mention `json:"mentions"`
	//Tags                []Tag        `gorm:"has_many:tags"`
	//Emojis              []Emoji      `gorm:"has_many:emojis"`

	//ReblogID int     `json:"-"`
	//Reblog   *Status `json:"reblog"`
	//PollID   int     `json:"-"`
	//Poll     *Poll   `json:"poll"`
	//CardID   int     `json:"-"`
	//Card     *Card   `json:"card"`

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
