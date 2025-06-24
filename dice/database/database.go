package database

type Database interface {
	GetUserById(id uint) (User, error)
	DeleteUserById(id uint) error
	GetAllUsers() ([]User, error)

	GetItemById(id uint) (Item, error)
	CreateItem(item Item) error
	GetAllItems() ([]Item, error)
	DeleteItemByID(id uint) error
	GetAllItemsFromLits(list List) ([]Item, error)
	MoveItemToList(id uint, list List) error
	ChangeItemTitle(id uint, title string) error
	ChangeItemDescription(id uint, description string) error

	CreateNewRepo(name, url string) error
	DeleteRepoById(id uint) error

	// Username must be identical to GitHub username.
	CreateUser(username string) error
}
