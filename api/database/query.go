package database

func (db *Database) GetUserById(id uint) (User, error) {
	var user User
	if err := db.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (db *Database) CreateUser(user User) (uint, error) {
	if err := db.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (db *Database) DeleteUserById(id uint) error {
	if err := db.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) GetAllUsers() ([]User, error) {
	var users []User
	if err := db.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (db *Database) GetItemById(id uint) (Item, error) {
	var item Item
	if err := db.db.First(&item, id).Error; err != nil {
		return Item{}, err
	}
	return item, nil
}

func (db *Database) CreateItem(item Item) (id uint, err error) {
	if err := db.db.Create(&item).Error; err != nil {
		return id, err
	}
	return item.ID, nil
}

func (db *Database) GetAllItems() ([]Item, error) {
	var items []Item
	if err := db.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Database) DeleteItemByID(id uint) error {
	if err := db.db.Delete(&Item{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) GetAllItemsWithStatus(boardId uint, list Status) ([]Item, error) {
	var items []Item
	if err := db.db.Where("board_id = ? AND status = ?", boardId, list).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Database) ChangeItemStatus(id uint, list Status) error {
	if err := db.db.Model(&Item{}).Where("id = ?", id).Update("status", list).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) SetItemData(id uint, data string) error {
	if err := db.db.Model(&Item{}).Where("id = ?", id).Update("data", data).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) GetItemData(id uint) (data string, err error) {
	var item Item
	if err := db.db.First(&item, id).Error; err != nil {
		return data, err
	}
	return item.Data, nil
}

func (db *Database) GetAllBoards() (boards []Board, err error) {
	if err := db.db.Find(&boards).Error; err != nil {
		return boards, err
	}
	return boards, err
}

func (db *Database) CreateBoard(board Board) (id uint, err error) {
	if err := db.db.Create(&board).Error; err != nil {
		return id, err
	}
	return board.ID, nil
}

func (db *Database) GetBoardById(id uint) (board Board, err error) {
	if err := db.db.Find(&board, id).Error; err != nil {
		return board, err
	}
	return board, nil
}

func (db *Database) UpdateBoard(board Board, id uint) error {
	if err := db.db.Model(&Board{}).Where("id = ?", id).Updates(board).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteBoard(id uint) error {
	if err := db.db.Delete(&Board{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) GetBoardUsers(boardId uint) ([]User, error) {
	var board Board
	if err := d.db.Preload("Users").First(&board, boardId).Error; err != nil {
		return nil, err
	}
	return board.Users, nil
}

func (d *Database) GetBoardItems(boardId uint) ([]Item, error) {
	var items []Item
	if err := d.db.Where("board_id = ?", boardId).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (d *Database) AddUserToBoard(boardId, userId uint) error {
	var board Board
	if err := d.db.First(&board, boardId).Error; err != nil {
		return err
	}
	var user User
	if err := d.db.First(&user, userId).Error; err != nil {
		return err
	}
	return d.db.Model(&board).Association("Users").Append(&user)
}

func (d *Database) RemoveUserFromBoard(boardId, userId uint) error {
	var board Board
	if err := d.db.First(&board, boardId).Error; err != nil {
		return err
	}
	var user User
	if err := d.db.First(&user, userId).Error; err != nil {
		return err
	}
	return d.db.Model(&board).Association("Users").Delete(&user)
}

func (d *Database) GetBoardItemsByStatus(boardId uint, status Status) ([]Item, error) {
	var items []Item
	if err := d.db.Where("board_id = ? AND status = ?", boardId, status).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
