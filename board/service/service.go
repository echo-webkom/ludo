package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/board/config"
)

type LudoService struct {
	config *config.Config
}

func New(config *config.Config) *LudoService {
	return &LudoService{
		config: config,
	}
}

// Send request to API using config URL. Method should be "GET", "POST" etc.
//
// Endpoint is the target endpoint, eg. "/users/12482", notice path values
// should already be formatted. Body is the request body, a reference to the
// object you want to encode as json, nil for empty body. Response is a
// reference to the response type, or nil for no response.
func (l *LudoService) request(method, endpoint string, body, response any) error {
	var bodyReader io.Reader
	if body != nil {
		bodyJson, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bodyJson)
	}

	url, err := url.JoinPath(l.config.ApiUrl, endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK && response != nil {
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
	var idRes database.ID
	err = l.request("POST", "/users", &user, &idRes)
	return idRes.ID, err
}

func (l *LudoService) GetUserById(id uint) (user database.User, err error) {
	err = l.request("GET", fmt.Sprintf("/users/%d", id), nil, &user)
	return user, err
}

func (l *LudoService) UpdateUser(id uint, user database.User) error {
	return l.request("PATCH", fmt.Sprintf("/users/%d", id), &user, nil)
}

func (l *LudoService) DeleteUser(id uint) error {
	return l.request("DELETE", fmt.Sprintf("/users/%d", id), nil, nil)
}

func (l *LudoService) GetAllBoards() (boards []database.Board, err error) {
	err = l.request("GET", "/boards", nil, &boards)
	return boards, err
}

func (l *LudoService) CreateBoard(board database.Board) (id uint, err error) {
	var idRes database.ID
	err = l.request("POST", "/boards", &board, &idRes)
	return idRes.ID, err
}

func (l *LudoService) GetBoardById(id uint) (board database.Board, err error) {
	err = l.request("GET", fmt.Sprintf("/boards/%d", id), nil, &board)
	return board, err
}

func (l *LudoService) UpdateBoard(id uint, board database.Board) error {
	return l.request("PATCH", fmt.Sprintf("/boards/%d", id), &board, nil)
}

func (l *LudoService) DeleteBoard(id uint) error {
	return l.request("DELETE", fmt.Sprintf("/boards/%d", id), nil, nil)
}

func (l *LudoService) GetAllUsersInBoard(id uint) (users []database.User, err error) {
	err = l.request("GET", fmt.Sprintf("/boards/%d/users", id), nil, &users)
	return users, err
}

func (l *LudoService) GetAllItemsInBoard(id uint) (items []database.Item, err error) {
	err = l.request("GET", fmt.Sprintf("/boards/%d/items", id), nil, &items)
	return items, err
}

func (l *LudoService) AddUserToBoard(boardId, userId uint) error {
	return l.request("POST", fmt.Sprintf("/boards/%d/users/%d", boardId, userId), nil, nil)
}

func (l *LudoService) RemoveUserFromBoard(boardId, userId uint) error {
	return l.request("DELETE", fmt.Sprintf("/boards/%d/users/%d", boardId, userId), nil, nil)
}

func (l *LudoService) GetAllItemsInBoardList(boardId, list uint) (items []database.Item, err error) {
	err = l.request("GET", fmt.Sprintf("/boards/%d/lists/%d/items", boardId, list), nil, &items)
	return items, err
}

func (l *LudoService) GetAllItems() (items []database.Item, err error) {
	err = l.request("GET", "/items", nil, &items)
	return items, err
}

func (l *LudoService) CreateItem(item database.Item) (id uint, err error) {
	var idRes database.ID
	err = l.request("POST", "/items", &item, &idRes)
	return idRes.ID, err
}

func (l *LudoService) GetItemById(id uint) (item database.Item, err error) {
	err = l.request("GET", fmt.Sprintf("/items/%d", id), nil, &item)
	return item, err
}

func (l *LudoService) UpdateItem(id uint, item database.Item) error {
	return l.request("PATCH", fmt.Sprintf("/items/%d", id), &item, nil)
}

func (l *LudoService) DeleteItem(id uint) error {
	return l.request("DELETE", fmt.Sprintf("/items/%d", id), nil, nil)
}

func (l *LudoService) MoveItemToList(itemId, list uint) error {
	return l.request("PATCH", fmt.Sprintf("/items/%d/move/%d", itemId, list), nil, nil)
}
