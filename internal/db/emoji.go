package db

type Emoji struct {
	Shortcode       string
	URL             string
	StaticURL       string
	VisibleInPicker bool
	Category        string
}
