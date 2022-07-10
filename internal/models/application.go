package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Name         string `json:"name"`
	Website      string `json:"website"`
	RedirectURIs string `json:"redirect_uris"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scopes       string `json:"scopes"`
}

type OAuthApplication struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Domain string `json:"domain"`
	UserID string `json:"user_id"`
}

type ApplicationResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Website      string `json:"website"`
	RedirectURIs string `json:"redirect_uris"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
