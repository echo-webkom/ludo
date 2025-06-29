package database

import (
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func newDatabase(db *gorm.DB) *Database {
	if err := db.AutoMigrate(&Board{}, &Item{}, &User{}); err != nil {
		log.Fatalf("migration: %v", err)
	}
	return &Database{db}
}

func NewTurso(url, token string) *Database {
	dbUrl := url + token
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return newDatabase(db)
}

func NewSQLite(filename string) *Database {
	if !strings.HasPrefix(filename, ":file") {
		filename = "file:" + filename
	}

	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return newDatabase(db)
}

func (db *Database) Close() error {
	raw, err := db.db.DB()
	if err != nil {
		return err
	}
	return raw.Close()
}
