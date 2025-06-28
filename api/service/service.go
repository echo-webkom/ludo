package service

import (
	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/database"
)

type Service struct {
	config *config.Config
	db     database.Database
}

func New(config *config.Config) *Service {
	return &Service{
		config,
		database.NewTursoDB(config),
	}
}

// Get all users across all boards.
func (s *Service) GetAllUsers() (users []database.User, err error) {
	return users, err
}

// Create new user from object. Returns id of created user.
func (s *Service) CreateUser(user database.User) (id uint, err error) {
	return id, err
}

// Get a user by its ID.
func (s *Service) GetUserById(id uint) (user database.User, err error) {
	return user, nil
}

// Update a user. All values are updated to those given.
func (s *Service) UpdateUser(id uint, user database.User) error {
	return nil
}

// Delete user and remove it from all boards.
func (s *Service) DeleteUser(id uint) error {
	return nil
}

// Get a list of all boards.
func (s *Service) GetAllBoards() (boards []database.Board, err error) {
	return boards, err
}

// Create a new board from object. Returns ID of created board.
func (s *Service) CreateBoard(board database.Board) (id uint, err error) {
	return id, err
}

// Gets a board by its ID.
func (s *Service) GetBoardById(id uint) (board database.Board, err error) {
	return board, err
}

// Update a board. All values are updated to those given.
func (s *Service) UpdateBoard(id uint, board database.Board) error {
	return nil
}

// Delete a board.
func (s *Service) DeleteBoard(id uint) error {
	return nil
}

// Get all users in the board.
func (s *Service) GetUsersInBoard(id uint) (users []database.User, err error) {
	return users, err
}

// Get all items in the board.
func (s *Service) GetItemsInBoard(id uint) (items []database.Item, err error) {
	return items, err
}

// Add an existing user to a board.
func (s *Service) AddUserToBoard(boardId, userId uint) error {
	return nil
}

// Remove an existing user from a board.
func (s *Service) RemoveUserFromBoard(boardId, userId uint) error {
	return nil
}

// Get all items in given list in board.
func (s *Service) GetAllItemsInList(boardId, listId uint) (items []database.Item, err error) {
	return items, err
}

// Get all items across all boards.
func (s *Service) GetAllItems() (items []database.Item, err error) {
	return items, err
}

// Create new item from object. Returns id of created item.
func (s *Service) CreateItem(item database.Item) (id uint, err error) {
	return id, err
}

func (s *Service) GetItemById(id uint) (item database.Item, err error) {
	return item, err
}

// Update an existing item.
func (s *Service) UpdateItem(item database.Item) error {
	return nil
}

func (s *Service) DeleteItem(id uint) error {
	return nil
}

// Move an item to another list.
func (s *Service) MoveItem(id, list uint) error {
	return nil
}
