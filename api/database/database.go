package database

type Database interface {
	GetUserById(id uint) (User, error)
	UpdateUser(id uint, user User) error
	CreateUser(user User) (uint, error)
	DeleteUserById(id uint) error
	GetAllUsers() ([]User, error)
	
	GetItemById(id uint) (Item, error)
	CreateItem(item Item) (uint, error)
	GetAllItems() ([]Item, error)
	DeleteItemByID(id uint) error
	GetAllItemsFromList(boardId, list uint) ([]Item, error)
	MoveItemToList(id uint, list uint) error
	UpdateItem(id uint, item Item) error
	
	GetAllBoards() ([]Board, error)
	CreateBoard(board Board) (uint, error)
	GetBoardById(id uint) (Board, error)
	UpdateBoard(id uint, board Board) error
	DeleteBoardById(id uint) error

	GetBoardUsers(boardId uint) ([]User, error)
	GetBoardItems(boardId uint) ([]Item, error)
	AddUserToBoard(boardId, userId uint) error
	RemoveUserFromBoard(boardId, userId uint) error
}
