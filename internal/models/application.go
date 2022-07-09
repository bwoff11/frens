package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Name    string `json:"name"`
	Website string `json:"website"`
}
