package db

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/bwoff11/frens/internal/models"
	badger "github.com/dgraph-io/badger/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB
var Badger *badger.DB

func Connect() {
	connectToPostgresql()
	connectToBadger()
}

func connectToPostgresql() {
	dbURL := "postgres://test:test@localhost:5432/frens"

	var err error
	Postgres, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Postgres.Set("gorm:auto_preload", true)

	Postgres.AutoMigrate(&models.Account{})
	Postgres.AutoMigrate(&models.Application{})
	Postgres.AutoMigrate(&models.Attachment{})
	Postgres.AutoMigrate(&models.Mention{})
	Postgres.AutoMigrate(&models.Tag{})
	Postgres.AutoMigrate(&models.Status{})
	Postgres.AutoMigrate(&models.Field{})
	Postgres.AutoMigrate(&models.Activity{})
	Postgres.AutoMigrate(&models.Emoji{})
	Postgres.AutoMigrate(&models.Source{})
	Postgres.AutoMigrate(&models.Hashtag{})
	Postgres.AutoMigrate(&models.OAuthApplication{})

	Postgres.Preload("Account").Find(&models.Status{})
}

func connectToBadger() {
	var err error
	Badger, err = badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		log.Fatal(err)
	}
}

func Sha256(password string) string {
	data := []byte(password)
	encrypted := fmt.Sprintf("%x", sha256.Sum256(data))
	return encrypted
}

func AddAccount(account *models.Account) {
	Postgres.Create(account)
}

func DeleteAccount(account *models.Account) {
	Postgres.Delete(account)
}

func GetAccountFollowingIDs(id *int64) []int64 {
	var following []int64
	Postgres.Model(&models.Account{}).Where("id = ?", id).Pluck("following", &following)
	return following
}
