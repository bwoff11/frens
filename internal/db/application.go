package db

import "time"

type Application struct {
	ID           int        `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Name         string     `json:"name"`
	Website      string     `json:"website"`
	RedirectURIs string     `json:"redirect_uris"`
	ClientID     string     `json:"client_id"`
	ClientSecret string     `json:"client_secret"`
	Scopes       string     `json:"scopes"`
}
