package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	ID              int64    `json:"id"`
	SourceAccountID int64    `json:"source_account_id"`
	SourceAccount   *Account `json:"source_account"`
	TargetAccountID int64    `json:"target_account_id"`
	TargetAccount   *Account `json:"target_account"`
}
