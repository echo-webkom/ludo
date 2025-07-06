package database

import (
	"errors"
)

func (db *tursoDB) GetUserById(id uint) (User, error) {
	var user User
	if res := db.db.First(&user, id); res.Error != nil {
		return User{}, errors.New("could not find user with id")
	}
	return user, nil
}

func (db *tursoDB) UpdateUser(id uint, user User) error {
	if res := db.db.Model(&User{}).Where("id = ?", id).Updates(user); res.Error != nil {
		return errors.New("could not update user")
	}
	return nil
}

func (db *tursoDB) CreateUser(user User) (uint, error) {
	if res := db.db.Create(&user); res.Error != nil {
		return 0, errors.New("could not create a new user")
	}
	return user.ID, nil
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
		return nil, errors.New("could not get all users")
	}
	return users, nil
}

func (db *tursoDB) GetItemById(id uint) (Item, error) {
	var item Item
	if res := db.db.Preload("Creator").Preload("Assignee").First(&item, id); res.Error != nil {
		return Item{}, errors.New("could not get item")
	}
	return item, nil
}

func (db *tursoDB) CreateItem(item Item) (uint, error) {
	if res := db.db.Create(&item); res.Error != nil {
		return 0, errors.New("could not create item")
	}
	return item.ID, nil
}

func (db *tursoDB) GetAllItems() ([]Item, error) {
	var items []Item
	if res := db.db.Preload("Creator").Preload("Assignee").Find(&items); res.Error != nil {
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

func (db *tursoDB) GetAllItemsFromList(boardId, list uint) ([]Item, error) {
	var board Board
	if res := db.db.Preload("Items", "list = ?", list).First(&board, boardId); res.Error != nil {
		return nil, errors.New("could not find board")
	}
	
	for i := range board.Items {
		db.db.Preload("Creator").Preload("Assignee").First(&board.Items[i], board.Items[i].ID)
	}
	
	return board.Items, nil
}

func (db *tursoDB) MoveItemToList(id uint, list uint) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Update("list", list); res.Error != nil {
		return errors.New("could not move item")
	}
	return nil
}

func (db *tursoDB) UpdateItem(id uint, item Item) error {
	if res := db.db.Model(&Item{}).Where("id = ?", id).Updates(item); res.Error != nil {
		return errors.New("could not update item")
	}
	return nil
}

func (db *tursoDB) GetAllBoards() ([]Board, error) {
	var boards []Board
	if res := db.db.Preload("Items").Preload("Users").Find(&boards); res.Error != nil {
		return nil, errors.New("could not find boards")
	}
	return boards, nil
}

func (db *tursoDB) CreateBoard(board Board) (uint, error) {
	if res := db.db.Create(&board); res.Error != nil {
		return 0, errors.New("could not create a new board")
	}
	return board.ID, nil
}

func (db *tursoDB) GetBoardById(id uint) (Board, error) {
	var board Board
	if res := db.db.Preload("Items").Preload("Users").First(&board, id); res.Error != nil {
		return Board{}, errors.New("could not find board")
	}
	return board, nil
}

func (db *tursoDB) UpdateBoard(id uint, board Board) error {
	if res := db.db.Model(&Board{}).Where("id = ?", id).Updates(board); res.Error != nil {
		return errors.New("could not update board")
	}
	return nil
}

func (db *tursoDB) DeleteBoardById(id uint) error {
	if res := db.db.Delete(&Board{}, id); res.Error != nil {
		return errors.New("could not delete board")
	}
	return nil
}

func (db *tursoDB) GetBoardUsers(boardId uint) ([]User, error) {
	var board Board
	if res := db.db.Preload("Users").First(&board, boardId); res.Error != nil {
		return nil, errors.New("could not find board")
	}
	return board.Users, nil
}

func (db *tursoDB) GetBoardItems(boardId uint) ([]Item, error) {
	var board Board
	if res := db.db.Preload("Items").Preload("Items.Creator").Preload("Items.Assignee").First(&board, boardId); res.Error != nil {
		return nil, errors.New("could not find board")
	}
	return board.Items, nil
}

func (db *tursoDB) AddUserToBoard(boardId, userId uint) error {
	var board Board
	if res := db.db.First(&board, boardId); res.Error != nil {
		return errors.New("could not find board")
	}
	
	var user User
	if res := db.db.First(&user, userId); res.Error != nil {
		return errors.New("could not find user")
	}
	
	if res := db.db.Model(&board).Association("Users").Append(&user); res != nil {
		return errors.New("could not add user to board")
	}
	
	return nil
}

func (db *tursoDB) RemoveUserFromBoard(boardId, userId uint) error {
	var board Board
	if res := db.db.First(&board, boardId); res.Error != nil {
		return errors.New("could not find board")
	}
	
	var user User
	if res := db.db.First(&user, userId); res.Error != nil {
		return errors.New("could not find user")
	}
	
	if res := db.db.Model(&board).Association("Users").Delete(&user); res != nil {
		return errors.New("could not remove user from board")
	}
	
	return nil
}