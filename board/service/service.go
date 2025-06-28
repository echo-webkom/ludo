package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/board/config"
)

type LudoService struct {
	config *config.Config
}

func New() *LudoService {
	return &LudoService{}
}

// Send request to API using config URL. Method should be "GET", "POST" etc. Endpoint is the
// target endpoint, eg. "/users/12482", notice path values should already be formatted. body
// is the request body, a reference to the object you want to encode as json, nil for empty
// body. response is a reference to the response type, or nil for no response.
func (l *LudoService) request(method string, endpoint string, body any, response any) error {
	var bodyReader *bytes.Reader
	if body != nil {
		bodyJson, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bodyJson)
	}

	path := path.Join(l.config.ApiUrl, endpoint)
	req, err := http.NewRequest(method, path, bodyReader)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if response != nil {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, response)
	}

	return err
}

func (l *LudoService) GetAllUsers() (users []database.User, err error) {
	err = l.request("GET", "/users", nil, &users)
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
	err = l.request("GET", "/boards", nil, &boards)
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
	err = l.request("GET", "/items", nil, &items)
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
