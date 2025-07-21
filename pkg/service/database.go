package service

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

// Implements LudoService
type Database struct {
	db *gorm.DB
}

func newDatabase(db *gorm.DB) *Database {
	if err := db.AutoMigrate(&model.Board{}, &model.Item{}, &model.User{}); err != nil {
		log.Fatalf("migration: %v", err)
	}
	return &Database{db}
}

func NewTursoService(url, token string) *Database {
	dbUrl := url + token
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return newDatabase(db)
}

func NewSQLiteService(filename string) *Database {
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

func (db *Database) User(id uint) (model.User, error) {
	var user model.User
	if err := db.db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (db *Database) NewUser(user model.User) (uint, error) {
	if err := db.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (db *Database) DeleteUser(id uint) error {
	if err := db.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) Users() ([]model.User, error) {
	var users []model.User
	if err := db.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (db *Database) Item(id uint) (model.Item, error) {
	var item model.Item
	if err := db.db.First(&item, id).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (db *Database) NewItem(boardId uint, item model.Item) (id uint, err error) {
	item.BoardID = boardId
	if err := db.db.Create(&item).Error; err != nil {
		return id, err
	}
	return item.ID, nil
}

func (db *Database) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	if err := db.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Database) DeleteItem(id uint) error {
	if err := db.db.Delete(&model.Item{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) GetAllItemsWithStatus(boardId uint, list model.Status) ([]model.Item, error) {
	var items []model.Item
	if err := db.db.Where("board_id = ? AND status = ?", boardId, list).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Database) ChangeItemStatus(id uint, list model.Status) error {
	if err := db.db.Model(&model.Item{}).Where("id = ?", id).Update("status", list).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) SetItemData(id uint, data string) error {
	if err := db.db.Model(&model.Item{}).Where("id = ?", id).Update("data", data).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) GetItemData(id uint) (data string, err error) {
	var item model.Item
	if err := db.db.First(&item, id).Error; err != nil {
		return data, err
	}
	return item.Data, nil
}

func (db *Database) Boards() (boards []model.Board, err error) {
	if err := db.db.Find(&boards).Error; err != nil {
		return boards, err
	}
	return boards, err
}

func (db *Database) NewBoard(board model.Board) (id uint, err error) {
	if err := db.db.Create(&board).Error; err != nil {
		return id, err
	}
	return board.ID, nil
}

func (db *Database) Board(id uint) (board model.Board, err error) {
	if err := db.db.Find(&board, id).Error; err != nil {
		return board, err
	}
	return board, nil
}

func (db *Database) UpdateBoard(id uint, board model.Board) error {
	if err := db.db.Model(&model.Board{}).Where("id = ?", id).Updates(board).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteBoard(id uint) error {
	if err := db.db.Delete(&model.Board{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) BoardUsers(boardId uint) ([]model.User, error) {
	var board model.Board
	if err := d.db.Preload("Users").First(&board, boardId).Error; err != nil {
		return nil, err
	}
	return board.Users, nil
}

func (d *Database) BoardItems(boardId uint) ([]model.Item, error) {
	var items []model.Item
	if err := d.db.Where("board_id = ?", boardId).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (d *Database) AddUserToBoard(boardId, userId uint) error {
	var board model.Board
	if err := d.db.First(&board, boardId).Error; err != nil {
		return err
	}
	var user model.User
	if err := d.db.First(&user, userId).Error; err != nil {
		return err
	}
	return d.db.Model(&board).Association("Users").Append(&user)
}

func (d *Database) RemoveUserFromBoard(boardId, userId uint) error {
	var board model.Board
	if err := d.db.First(&board, boardId).Error; err != nil {
		return err
	}
	var user model.User
	if err := d.db.First(&user, userId).Error; err != nil {
		return err
	}
	return d.db.Model(&board).Association("Users").Delete(&user)
}

func (d *Database) BoardItemsWithStatus(boardId uint, status model.Status) ([]model.Item, error) {
	var items []model.Item
	if err := d.db.Where("board_id = ? AND status = ?", boardId, status).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Database) UpdateItem(id uint, item model.Item) error {
	if err := db.db.Model(&model.Item{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUser(id uint, user model.User) error {
	if err := db.db.Model(&model.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) ItemData(id uint) (string, error) {
	return "", nil
}
