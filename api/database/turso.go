package database

import (
	"log"

	"github.com/echo-webkom/ludo/api/config"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type tursoDB struct {
	db *gorm.DB
}

func NewTursoDB(config *config.Config) Database {
	if config.IsDev {
		db, err := gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{})
		if err != nil {
			log.Fatal("could not load local database", err)
		}
		if err := db.AutoMigrate(&Item{}); err != nil {
			log.Fatalf("migration: %v", err)
		}

		return &tursoDB{db}
	}

	db, err := gorm.Open(loadRemoteDB(config), &gorm.Config{})
	if err != nil {
		log.Fatal("could not load remote database", err)
	}

	if err := db.AutoMigrate(&Item{}, &User{}, &Board{}); err != nil {
		log.Fatalf("migration: %v", err)
	}
	return &tursoDB{db}
}

func (db *tursoDB) Close() error {
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
