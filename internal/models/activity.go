package models

type Activity struct {
	Week          int `json:"week"`
	Statuses      int `json:"statuses"`
	Logins        int `json:"logins"`
	Registrations int `json:"registrations"`
}
