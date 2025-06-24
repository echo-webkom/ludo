package database

import (
	"errors"
)

type DatabaseMethods interface {
	GetUserById(id uint) (User, error)
	CreateUser(name string) error
	DeleteUserById(id uint) error
	GetAllUsers() ([]User, error)
	GetItemById(id uint) (Item, error)
	CreateItem(title, description string, repo Repo, creator User, conectedBranch string, list List)
	GetAllItems() ([]Item, error)
	DeleteItemByID(id uint) error
	GetAllItemsFromLits(list List) ([]Item, error)
	MoveItemToList(id uint, list List) error
	ChangeItemTitle(id uint, title string) error
	ChangeItemDescription(id uint, description string) error
	CreateNewRepo(name, url string) error
	DeleteRepoById(id uint) error
}

func (db *Database) GetUserById(id uint) (User, error) {
	var user User
	if res := db.db.First(&user, id); res.Error != nil {
		return User{}, errors.New("could not find usre with id")
	}
	return user, nil
}

func (db *Database) CreateUser(name string) error {
	if res := db.db.Create(User{Name: name}); res.Error != nil {
		return errors.New("could not create a new user")
	}
	return nil
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

func (db *Database) CreateItem(title, description string, repo Repo, creator User, connectedBranch string, list List) error {
	item := Item{Title: title, Description: description, Repo: repo, Creator: creator, ConnectedBranch: connectedBranch, List: list, Assignee: User{}}
	if res := db.db.Create(&item); res.Error != nil {
		return errors.New("could not create	item")
	}
	return nil
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

func (db *Database) ChangeItemDescription(id uint, description string) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("description", description); res.Error != nil {
		return errors.New("could not change item description")
	}
	return nil
}

func (db *Database) CreateNewRepo(name, url string) error {
	repo := Repo{Name: name, URL: url}
	if res := db.db.Create(&repo); res.Error != nil {
		return errors.New("could not create new user")
	}
	return nil
}

func (db *Database) DeleteRepoById(id uint) error {
	if res := db.db.Delete(&Repo{}, id); res.Error != nil {
		return errors.New("could not delet the repo")
	}
	return nil
}
