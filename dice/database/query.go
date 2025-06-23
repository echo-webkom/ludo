package database

import (
	"errors"
	"log"
)

func (db *Database) GetUserById(id uint) (User, error) {
	var user User
	if res := db.db.First(&user, id); res.Error != nil {
		return User{}, errors.New("could not find usre with id")
	}
	return user, nil
}

func (db *Database) CreateUser(name string) (User,	error) {
	user := User{Name: name}
	if res := db.db.Create(&user); res.Error != nil {
		return User{}, errors.New("could not create a new user")
	}
	return user, nil
}

func (db *Database) GetAllUsers() ([]User, error) {
	var users []User
	if res := db.db.Find(&users); res.Error != nil {
		return nil, errors.New("courl not get all userts")
	}
	return users, nil
}


func (db *Database) GetItemById(id uint) (Item, error) {
	var item Item
	if res := db.db.First(&item, id); res.Error != nil {
		return Item{}, errors.New("could not get item")
	}
	return item, nil
}

func (db *Database) GetAllItems() ([]Item, error) {
	var items []Item
	if res := db.db.Find(&items); res.Error != nil {
		return nil, errors.New("could not get all items")
	}
	return items, nil 
}

func (db *Database) DeleteItemByID(id uint) error {
	if res := db.db.Delete(&Item{}, id); res.Error != nil {
		return errors.New("could not delete item")
	}	
	return nil 
}

func (db *Database) DeleteUserById(id uint) error {
	if res := db.db.Delete(&User{}, id); res.Error != nil {
		return errors.New("could not delete user")
	}
	return nil
}


func (db *Database) GetAllItemsFromLits(list List) ([]Item, error) {
	var items []Item
	if res := db.db.Find(&items, "list = ?", list); res.Error != nil {
		return nil, errors.New("could not find all items from list")
	}
	return items, nil
}

func (db *Database) MoveItemToList(id uint, list List) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("list", list); res.Error != nil {
		return errors.New("could not move item")
	}
	return nil
}

func (db *Database) ChangeItemTitle(id uint, title string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("title", title); res.Error != nil {
		return errors.New("could not change item title")
	}
	return nil
}


func (db *Database) ChangeItemDesctription(id uint, description string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("description",description); res.Error != nil {
		return errors.New("could not change item description")
	}
	return nil
}
