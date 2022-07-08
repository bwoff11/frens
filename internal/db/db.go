package db

import (
	"log"

	"github.com/bwoff11/frens/internal/api/v1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbURL := "postgres://test:test@localhost:5432/frens"

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Field{})
	DB.AutoMigrate(&models.Emoji{})
	DB.AutoMigrate(&models.Source{})
}
