package service

import "github.com/echo-webkom/ludo/pkg/model"

type LudoService interface {
	// Get the board service.
	Boards() Boards
	// Get the user service.
	Users() Users
}

type Users interface {
	// Get all users.
	All() ([]model.User, error)
	// Create a new user. Returns the ID.
	New(user model.User) (uint, error)
	// Get a user by ID.
	Id(id uint) User
}

type User interface {
	// Get the user data.
	Get() (model.User, error)
	// Update the user. Replaces all fields with those given.
	Update(user model.User) error
	// Delete the user and remove it from all boards.
	Delete() error
}

type Boards interface {
	// Get all boards.
	All() ([]model.Board, error)
	// Create a new board. Returns the ID.
	New(board model.Board) (uint, error)
	// Get a board by ID.
	Id(id uint) Board
}

type Board interface {
	// Get board data.
	Get() (model.Board, error)
	// Update the board. Replaces all fields with those given.
	Update(board model.Board) error
	// Delete the board and all its items.
	Delete() error
	// Get the items service for this board.
	Items() Items
	// Get all users assigned to this board.
	Users() ([]model.User, error)
	// Add an existing user to this board.
	AddUser(userId uint) error
	// Remove a user from this board.
	RemoveUser(userId uint) error
}

type Items interface {
	// Get all items in this board.
	All() ([]model.Item, error)
	// Create new item in this board.
	New(item model.Item) (uint, error)
	// Get all items with a given status in this board.
	WithStatus(list uint) ([]model.Item, error)
	// Get an item by ID.
	Id(id uint) Item
}

type Item interface {
	// Get item data.
	Get() (model.Item, error)
	// Update this item. Replaces all fields with those given.
	Update(item model.Item) error
	// Delete this item.
	Delete() error
	// Get this items data segment.
	Data() (string, error)
	// Set this items data segment.
	SetData(data string) error
}
