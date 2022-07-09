package models

type Application struct {
	Name    string      `json:"name"`
	Website interface{} `json:"website"`
}
