package database

type Database interface {
	// Username must be identical to GitHub username.
	CreateUser(displayName, githubUsername string) error
	GetUserById(id uint) (User, error)
	DeleteUserById(id uint) error
	GetAllUsers() ([]User, error)

	GetItemById(id uint) (Item, error)
	CreateItem(item Item) error
	GetAllItems() ([]Item, error)
	DeleteItemByID(id uint) error
	GetAllItemsFromLits(list uint) ([]Item, error)
	MoveItemToList(id uint, list uint) error
}
