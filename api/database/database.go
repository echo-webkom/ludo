package database

import (
	"log"
	"strings"

	"github.com/echo-webkom/ludo/pkg/model"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db *gorm.DB
}

func newDatabase(db *gorm.DB) *Database {
	if err := db.AutoMigrate(&model.Board{}, &model.Item{}, &model.User{}); err != nil {
		log.Fatalf("migration: %v", err)
	}
	return &Database{db}
}

func NewTurso(url, token string) *Database {
	dbUrl := url + token
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return newDatabase(db)
}

func NewSQLite(filename string) *Database {
	if !strings.HasPrefix(filename, "file:") {
		filename = "file:" + filename
	}

	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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
