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
	dbURL := "postgres://test:test@localhost:5432/postgres"

	var err error
	Postgres, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Postgres.Set("gorm:auto_preload", true)

	Postgres.AutoMigrate(
		&models.Account{},
		&models.Application{},
		&models.Mention{},
		&models.Activity{},
		&models.Status{},
		&models.Field{},
		&models.Attachment{},
		&models.Emoji{},
		&models.Tag{},
		&models.Source{},
		&models.Hashtag{},
	)

	Postgres.Preload("Account").Find(&models.Status{})
	//Postgres.Preload("Status").Find(&models.Account{})
}

func connectToBadger() {
	ops := badger.DefaultOptions("")
	ops.Logger = nil
	var err error
	Badger, err = badger.Open(ops.WithInMemory(true))
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
