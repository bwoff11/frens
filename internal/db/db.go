package db

import (
	"log"

	"github.com/bwoff11/frens/internal/models"
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
	DB.AutoMigrate(&models.Status{})
	//DB.AutoMigrate(&models.Field{})
	//DB.AutoMigrate(&models.Activity{})
	//DB.AutoMigrate(&models.Emoji{})
	//DB.AutoMigrate(&models.Source{})
	DB.AutoMigrate(&models.Hashtag{})
	DB.AutoMigrate(&models.Application{})
}

func AddAccount(account *models.Account) {
	DB.Create(account)
}

func DeleteAccount(account *models.Account) {
	DB.Delete(account)
}

func GetAccountFollowingIDs(id *int64) []int64 {
	var following []int64
	DB.Model(&models.Account{}).Where("id = ?", id).Pluck("following", &following)
	return following
}
