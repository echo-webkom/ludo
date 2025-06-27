package database

import (
	"errors"
)

func (db *tursoDB) GetUserById(id uint) (User, error) {
	var user User
	if res := db.db.First(&user, id); res.Error != nil {
		return User{}, errors.New("could not find usre with id")
	}
	return user, nil
}

func (db *tursoDB) CreateUser(name string) error {
	if res := db.db.Create(User{Name: name}); res.Error != nil {
		return errors.New("could not create a new user")
	}
	return nil
}

func (db *tursoDB) DeleteUserById(id uint) error {
	if res := db.db.Delete(&User{}, id); res.Error != nil {
		return errors.New("could not delete user")
	}
	return nil
}

func (db *tursoDB) GetAllUsers() ([]User, error) {
	var users []User
	if res := db.db.Find(&users); res.Error != nil {
		return nil, errors.New("courl not get all userts")
	}
	return users, nil
}

func (db *tursoDB) GetItemById(id uint) (Item, error) {
	var item Item
	if res := db.db.First(&item, id); res.Error != nil {
		return Item{}, errors.New("could not get item")
	}
	return item, nil
}

func (db *tursoDB) CreateItem(item Item) error {
	if res := db.db.Create(&item); res.Error != nil {
		return errors.New("could not create	item")
	}
	return nil
}

func (db *tursoDB) GetAllItems() ([]Item, error) {
	var items []Item
	if res := db.db.Find(&items); res.Error != nil {
		return nil, errors.New("could not get all items")
	}
	return items, nil
}

func (db *tursoDB) DeleteItemByID(id uint) error {
	if res := db.db.Delete(&Item{}, id); res.Error != nil {
		return errors.New("could not delete item")
	}
	return nil
}

func (db *tursoDB) GetAllItemsFromLits(list List) ([]Item, error) {
	var items []Item
	if res := db.db.Find(&items, "list = ?", list); res.Error != nil {
		return nil, errors.New("could not find all items from list")
	}
	return items, nil
}

func (db *tursoDB) MoveItemToList(id uint, list List) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("list", list); res.Error != nil {
		return errors.New("could not move item")
	}
	return nil
}

func (db *tursoDB) ChangeItemTitle(id uint, title string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("title", title); res.Error != nil {
		return errors.New("could not change item title")
	}
	return nil
}

func (db *tursoDB) ChangeItemDescription(id uint, description string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("description", description); res.Error != nil {
		return errors.New("could not change item description")
	}
	return nil
}

func (db *tursoDB) CreateNewRepo(name, url string) error {
	repo := Repo{Name: name, URL: url}
	if res := db.db.Create(&repo); res.Error != nil {
		return errors.New("could not create new user")
	}
	return nil
}

func (db *tursoDB) DeleteRepoById(id uint) error {
	if res := db.db.Delete(&Repo{}, id); res.Error != nil {
		return errors.New("could not delet the repo")
	}
	return nil
}
