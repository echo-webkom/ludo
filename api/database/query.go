package database

import (
	"errors"
)

func (db *Database) GetUserById(id uint) (User, error) {
	var user User
	if res := db.db.First(&user, id); res.Error != nil {
		return User{}, errors.New("could not find usre with id")
	}
	return user, nil
}

func (db *Database) CreateUser(user User) (uint, error) {
	if res := db.db.Create(&user); res.Error != nil {
		return 0, errors.New("could not create a new user")
	}
	return user.ID, nil
}

func (db *Database) DeleteUserById(id uint) error {
	if res := db.db.Delete(&User{}, id); res.Error != nil {
		return errors.New("could not delete user")
	}
	return nil
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

func (db *Database) CreateItem(item Item) (id uint, err error) {
	if res := db.db.Create(&item); res.Error != nil {
		return id, errors.New("could not create	item")
	}
	return item.ID, nil
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

func (db *Database) GetAllItemsWithStatus(boardId uint, list Status) ([]Item, error) {
	var items []Item
	if res := db.db.Find(&items, "status = ?", list); res.Error != nil {
		return nil, errors.New("could not find items with status")
	}
	return items, nil
}

func (db *Database) ChangeItemStatus(id uint, list Status) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("status", list); res.Error != nil {
		return errors.New("could not set item status")
	}
	return nil
}

func (db *Database) ChangeItemTitle(id uint, title string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("title", title); res.Error != nil {
		return errors.New("could not change item title")
	}
	return nil
}

func (db *Database) ChangeItemDescription(id uint, description string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("description", description); res.Error != nil {
		return errors.New("could not change item description")
	}
	return nil
}

func (db *Database) SetItemData(id uint, data string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("data", data); res.Error != nil {
		return errors.New("could not set data field")
	}
	return nil
}

func (db *Database) GetItemData(id uint) (data string, err error) {
	var item Item
	if res := db.db.First(&item, id); res.Error != nil {
		return data, errors.New("item not found")
	}
	return item.Data, nil
}
