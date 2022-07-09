package models

type Emoji struct {
	Shortcode       string `gorm:"column:shortcode;primary_key"`
	URL             string `gorm:"column:url"`
	StaticURL       string `gorm:"column:static_url"`
	VisibleInPicker bool   `gorm:"column:visible_in_picker"`
	Category        string `gorm:"column:category"`
}
