package db

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/bwoff11/frens/internal/config"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB
var Badger *badger.DB

func Connect() {
	logrus.Info("Initation database connections")
	connectToPostgresql()
	connectToBadger()
}

func connectToPostgresql() {
	logrus.Info("Attempting to connect to Postgresql")

	dbURL := "postgres://" +
		config.C.Database.User +
		":" + config.C.Database.Password +
		"@" + config.C.Database.Host +
		":" + config.C.Database.Port +
		"/" + config.C.Database.Database

	var err error
	Postgres, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Postgres.Set("gorm:auto_preload", true)

	Postgres.AutoMigrate(
		&Account{},
		&Application{},
		&Mention{},
		&Activity{},
		&Status{},
		&Field{},
		&Attachment{},
		&Emoji{},
		&Reblog{},
		&Poll{},
		&Card{},
		&Tag{},
		&Source{},
		&Hashtag{},
		&Relationship{},
	)

	logrus.Info("Finished migrating models to Postgresql")

	Postgres.Preload("Account").Find(&Status{})
	//Postgres.Preload("Status").Find(&Account{})

	logrus.Info("Finished connecting to Postgresql")
}

func connectToBadger() {
	logrus.Info("Attempting to connect to Badger")

	ops := badger.DefaultOptions("")
	ops.Logger = nil
	var err error
	Badger, err = badger.Open(ops.WithInMemory(true))
	if err != nil {
		log.Fatal(err)
	}

	logrus.Info("Finished connecting to Badger")
}

func Sha256(password string) string {
	data := []byte(password)
	encrypted := fmt.Sprintf("%x", sha256.Sum256(data))
	return encrypted
}

func AddAccount(account *Account) {
	Postgres.Create(account)
}

func DeleteAccount(account *Account) {
	Postgres.Delete(account)
}

func GetAccountFollowingIDs(id *int64) []int64 {
	var following []int64
	Postgres.Model(&Account{}).Where("id = ?", id).Pluck("following", &following)
	return following
}
