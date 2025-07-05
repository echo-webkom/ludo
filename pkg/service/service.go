package service

import "github.com/echo-webkom/ludo/pkg/model"

type LudoService interface {
	// Get a board by ID.
	GetBoard(id uint) (model.Board, error)
	// Update the board. Replaces all fields with those given.
	UpdateBoard(board model.Board) error
	// Delete the board and all its items.
	DeleteBoard(id uint) error
	// Create a new board. Returns the ID.
	NewBoard(board model.Board) (uint, error)
	// Get all boards.
	GetAllBoards() ([]model.Board, error)

	// Get a user by ID.
	GetUser(id uint) (model.User, error)
	// Update the user. Replaces all fields with those given.
	UpdateUser(user model.User) error
	// Delete the user and remove it from all boards.
	DeleteUser(id uint) error
	// Create a new user. Returns the ID.
	NewUser(user model.User) (uint, error)
	// Get all users.
	GetAllUsers() ([]model.User, error)

	// Get all users assigned to a board.
	GetBoardUsers(boardID uint) ([]model.User, error)
	// Add an existing user to a board.
	AddUserToBoard(boardID uint, userID uint) error
	// Remove a user from a board.
	RemoveUserFromBoard(boardID uint, userID uint) error

	// Get an item in a board by ID.
	GetItem(boardID uint, itemID uint) (model.Item, error)
	// Update an item. Replaces all fields with those given.
	UpdateItem(boardID uint, item model.Item) error
	// Delete an item.
	DeleteItem(boardID uint, itemID uint) error
	// Create a new item in a board. Returns the ID.
	NewItem(boardID uint, item model.Item) (uint, error)
	// Get all items in a board.
	GetBoardItems(boardID uint) ([]model.Item, error)
	// Get all items in a board with a given status.
	GetBoardItemsWithStatus(boardID uint, status model.Status) ([]model.Item, error)

	// Get this item's data segment.
	GetItemData(boardID uint, itemID uint) (string, error)
	// Set this item's data segment.
	SetItemData(boardID uint, itemID uint, data string) error
}
