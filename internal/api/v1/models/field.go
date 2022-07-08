package models

import "time"

type Field struct {
	// The key of a given field's key-value pair.
	Name string `gorm:"column:name" json:"name"`
	// The value associated with the name key.
	Value string `gorm:"column:value" json:"value"`
	// Timestamp of when the server verified a URL value for a rel="me” link. Type: String (ISO 8601 Datetime) if value is a verified URL. Otherwise, null.
	VerifiedAt time.Time `gorm:"column:verified_at" json:"verified_at"`
}
