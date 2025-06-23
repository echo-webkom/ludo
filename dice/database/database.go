package database

import (
	"log"

	"github.com/echo-webkom/ludo/dice/config"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func New(config *config.Config) *Database {
	if config.IsDev {
		db, err := gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{})
		if err != nil {
			log.Fatal("could not load local database", err)
		}
		if err := db.AutoMigrate(&Item{}); err != nil {
			log.Fatalf("migration: %v", err)
		}

		return &Database{db}
	}

	db, err := gorm.Open(loadRemoteDB(config), &gorm.Config{})
	if err != nil {
		log.Fatal("could not load remote database", err)
	}

	if err := db.AutoMigrate(&Item{}, &User{}, &Repo{}); err != nil {
		log.Fatalf("migration: %v", err)
	}
	return &Database{db}
}

func (db *Database) Close() error {
	raw, err := db.db.DB()
	if err != nil {
		return err
	}
	return raw.Close()
}

func loadRemoteDB(config *config.Config) gorm.Dialector {
	url := config.DatabaseURL + config.DatabaseToken
	return sqlite.Open(url)
}
