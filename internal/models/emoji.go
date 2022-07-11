package models

type Emoji struct {
	Shortcode       string
	URL             string
	StaticURL       string
	VisibleInPicker bool
	Category        string
}
