package service

import "github.com/echo-webkom/ludo/pkg/model"

type LudoService interface {
	// Get a board by ID.
	Board(id uint) (model.Board, error)
	// Get all boards.
	Boards() ([]model.Board, error)
	// Update the board. Replaces all fields with those given.
	UpdateBoard(id uint, board model.Board) error
	// Delete the board and all its items.
	DeleteBoard(id uint) error
	// Create a new board. Returns the ID.
	NewBoard(board model.Board) (uint, error)

	// Get a user by ID.
	User(id uint) (model.User, error)
	// Get all users.
	Users() ([]model.User, error)
	// Update the user. Replaces all fields with those given.
	UpdateUser(id uint, user model.User) error
	// Delete the user and remove it from all boards.
	DeleteUser(id uint) error
	// Create a new user. Returns the ID.
	NewUser(user model.User) (uint, error)

	// Get all users assigned to a board.
	BoardUsers(boardID uint) ([]model.User, error)
	// Add an existing user to a board.
	AddUserToBoard(boardID uint, userID uint) error
	// Remove a user from a board.
	RemoveUserFromBoard(boardID uint, userID uint) error

	// Get an item in a board by ID.
	Item(itemID uint) (model.Item, error)
	// Update an item. Replaces all fields with those given.
	UpdateItem(id uint, item model.Item) error
	// Delete an item.
	DeleteItem(itemID uint) error
	// Create a new item in a board. Returns the ID.
	NewItem(boardID uint, item model.Item) (uint, error)
	// Get all items in a board.
	BoardItems(boardID uint) ([]model.Item, error)
	// Get all items in a board with a given status.
	BoardItemsWithStatus(boardID uint, status model.Status) ([]model.Item, error)

	// Get this item's data segment.
	ItemData(itemID uint) (string, error)
	// Set this item's data segment.
	SetItemData(itemID uint, data string) error
}
