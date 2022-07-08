package models

import "time"

type Filter struct {
	ID           string    `json:"id"`
	Phrase       string    `json:"phrase"`
	Context      []string  `json:"context"`
	WholeWord    bool      `json:"whole_word"`
	ExpiresAt    time.Time `json:"expires_at"`
	Irreversible bool      `json:"irreversible"`
}
