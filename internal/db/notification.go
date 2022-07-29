package db

import "gorm.io/gorm"

type NotificationType string

const (
	NotificationTypeFollow        NotificationType = "follow"
	NotificationTypeFollowRequest NotificationType = "follow_request"
	NotificationTypeMention       NotificationType = "mention"
	NotificationTypeReblog        NotificationType = "reblog"
	NotificationTypeFavourite     NotificationType = "favourite"
	NotificationTypePoll          NotificationType = "poll"
	NotificationTypeStatus        NotificationType = "status"
)

type Notification struct {
	gorm.Model
	ID               uint64           `json:"id"`
	NotificationType NotificationType `json:"type"`
	AccountID        uint64           `json:"account_id"`
	Account          string           `json:"account"`
}
