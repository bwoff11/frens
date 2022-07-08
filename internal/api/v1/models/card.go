package models

type Card struct {
	URL          string      `json:"url"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Type         string      `json:"type"`
	AuthorName   string      `json:"author_name"`
	AuthorURL    string      `json:"author_url"`
	ProviderName string      `json:"provider_name"`
	ProviderURL  string      `json:"provider_url"`
	HTML         string      `json:"html"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	Image        interface{} `json:"image"`
	EmbedURL     string      `json:"embed_url"`
}
