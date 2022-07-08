package models

import "time"

type Status struct {
	ID                 string        `json:"id"`
	CreatedAt          time.Time     `json:"created_at"`
	InReplyToID        interface{}   `json:"in_reply_to_id"`
	InReplyToAccountID interface{}   `json:"in_reply_to_account_id"`
	Sensitive          bool          `json:"sensitive"`
	SpoilerText        string        `json:"spoiler_text"`
	Visibility         string        `json:"visibility"`
	Language           string        `json:"language"`
	URI                string        `json:"uri"`
	URL                string        `json:"url"`
	RepliesCount       int           `json:"replies_count"`
	ReblogsCount       int           `json:"reblogs_count"`
	FavouritesCount    int           `json:"favourites_count"`
	Favourited         bool          `json:"favourited"`
	Reblogged          bool          `json:"reblogged"`
	Muted              bool          `json:"muted"`
	Bookmarked         bool          `json:"bookmarked"`
	Content            string        `json:"content"`
	Reblog             interface{}   `json:"reblog"`
	Application        Application   `json:"application"`
	Account            Account       `json:"account"`
	MediaAttachments   []interface{} `json:"media_attachments"`
	Mentions           []interface{} `json:"mentions"`
	Tags               []interface{} `json:"tags"`
	Emojis             []interface{} `json:"emojis"`
	Card               Card          `json:"card"`
	Poll               interface{}   `json:"poll"`
}
