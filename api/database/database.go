package database

type Database interface {
	GetAllUsers() ([]User, error)
	// Create new user, returns its ID.
	CreateUser(user User) (id uint, err error)
	GetUserById(id uint) (User, error)
	// Delete a user and remove it from all boards.
	DeleteUserById(id uint) error
	GetAllItems() ([]Item, error)
	// Create new item, returns its ID.
	CreateItem(item Item) (uint, error)
	GetItemById(id uint) (Item, error)
	DeleteItemByID(id uint) error
	GetAllItemsFromList(boardId, list uint) ([]Item, error)
	MoveItemToList(id uint, list uint) error
	UpdateUser(id uint, user User) error
	UpdateItem(id uint, item Item) error
}
