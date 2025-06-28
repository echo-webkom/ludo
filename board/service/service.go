package service

import "github.com/echo-webkom/ludo/api/database"

type LudoService struct {
}

func New() *LudoService {
	return &LudoService{}
}

func (l *LudoService) GetAllUsers() (users []database.User, err error) {
	return users, err
}

func (l *LudoService) CreateUser(user database.User) (id uint, err error) {
	return id, err
}

func (l *LudoService) GetUserById(id uint) (user database.User, err error) {
	return user, err
}

func (l *LudoService) UpdateUser(id uint, user database.User) error {
	return nil
}

func (l *LudoService) DeleteUser(id uint) error {
	return nil
}

func (l *LudoService) GetAllBoards() (boards []database.Board, err error) {
	return boards, err
}

func (l *LudoService) CreateBoard(board database.Board) error {
	return nil
}

func (l *LudoService) GetBoardById(id uint) (board database.Board, err error) {
	return board, err
}

func (l *LudoService) UpdateBoard(id uint, board database.Board) error {
	return nil
}

func (l *LudoService) DeleteBoard(id uint) error {
	return nil
}

func (l *LudoService) GetAllUsersInBoard(id uint) (users []database.User, err error) {
	return users, err
}

func (l *LudoService) GetAllItemsInBoard(id uint) (items []database.Item, err error) {
	return items, err
}

func (l *LudoService) AddUserToBoard(boardId, userId uint) error {
	return nil
}

func (l *LudoService) RemoveUserFromBoard(boardId, userId uint) error {
	return nil
}

func (l *LudoService) GetAllItemsInBoardList(boardId, list uint) (items []database.Item, err error) {
	return items, err
}

func (l *LudoService) GetAllItems() (items []database.Item, err error) {
	return items, err
}

func (l *LudoService) CreateItem(item database.Item) (id uint, err error) {
	return id, err
}

func (l *LudoService) GetItemById(id uint) (item database.Item, err error) {
	return item, err
}

func (l *LudoService) UpdateItem(id uint, item database.Item) error {
	return nil
}

func (l *LudoService) DeleteItem(id uint) error {
	return nil
}

func (l *LudoService) MoveItemToList(itemId, list uint) error {
	return nil
}
