package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/echo-webkom/ludo/pkg/model"
)

type apiService struct {
	url string
}

// Create new LudoService using the Ludo REST API as backend.
func NewApiService(apiUrl string) LudoService {
	return &apiService{
		url: apiUrl,
	}
}

// Send request to API using config URL. Method should be "GET", "POST" etc.
//
// Endpoint is the target endpoint, eg. "/users/12482", notice path values
// should already be formatted. Body is the request body, a reference to the
// object you want to encode as json, nil for empty body. Response is a
// reference to the response type, or nil for no response.
func (s *apiService) request(method, endpoint string, body, response any) error {
	var bodyReader io.Reader
	if body != nil {
		bodyJson, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bodyJson)
	}

	url, err := url.JoinPath(s.url, endpoint)
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

func (s *apiService) Users() (users []model.User, err error) {
	err = s.request("GET", "/users", nil, &users)
	return users, err
}

func (s *apiService) NewUser(user model.User) (id uint, err error) {
	var idRes model.ID
	err = s.request("POST", "/users", &user, &idRes)
	return idRes.ID, err
}

func (s *apiService) User(id uint) (user model.User, err error) {
	err = s.request("GET", fmt.Sprintf("/users/%d", id), nil, &user)
	return user, err
}

func (s *apiService) UpdateUser(id uint, user model.User) error {
	return s.request("PATCH", fmt.Sprintf("/users/%d", id), &user, nil)
}

func (s *apiService) DeleteUser(id uint) error {
	return s.request("DELETE", fmt.Sprintf("/users/%d", id), nil, nil)
}

func (s *apiService) Boards() (boards []model.Board, err error) {
	err = s.request("GET", "/boards", nil, &boards)
	return boards, err
}

func (s *apiService) NewBoard(board model.Board) (id uint, err error) {
	var idRes model.ID
	err = s.request("POST", "/boards", &board, &idRes)
	return idRes.ID, err
}

func (s *apiService) Board(id uint) (board model.Board, err error) {
	err = s.request("GET", fmt.Sprintf("/boards/%d", id), nil, &board)
	return board, err
}

func (s *apiService) UpdateBoard(id uint, board model.Board) error {
	return s.request("PATCH", fmt.Sprintf("/boards/%d", id), &board, nil)
}

func (s *apiService) DeleteBoard(id uint) error {
	return s.request("DELETE", fmt.Sprintf("/boards/%d", id), nil, nil)
}

func (s *apiService) BoardUsers(id uint) (users []model.User, err error) {
	err = s.request("GET", fmt.Sprintf("/boards/%d/users", id), nil, &users)
	return users, err
}

func (s *apiService) BoardItems(id uint) (items []model.Item, err error) {
	err = s.request("GET", fmt.Sprintf("/boards/%d/items", id), nil, &items)
	return items, err
}

func (s *apiService) AddUserToBoard(boardId, userId uint) error {
	return s.request("POST", fmt.Sprintf("/boards/%d/users/%d", boardId, userId), nil, nil)
}

func (s *apiService) RemoveUserFromBoard(boardId, userId uint) error {
	return s.request("DELETE", fmt.Sprintf("/boards/%d/users/%d", boardId, userId), nil, nil)
}

func (s *apiService) BoardItemsWithStatus(boardId uint, status model.Status) (items []model.Item, err error) {
	err = s.request("GET", fmt.Sprintf("/boards/%d/lists/%d/items", boardId, status), nil, &items)
	return items, err
}

func (s *apiService) Items() (items []model.Item, err error) {
	err = s.request("GET", "/items", nil, &items)
	return items, err
}

func (s *apiService) NewItem(boardId uint, item model.Item) (id uint, err error) {
	var idRes model.ID
	err = s.request("POST", "/items", &item, &idRes)
	return idRes.ID, err
}

func (s *apiService) Item(id uint) (item model.Item, err error) {
	err = s.request("GET", fmt.Sprintf("/items/%d", id), nil, &item)
	return item, err
}

func (s *apiService) UpdateItem(id uint, item model.Item) error {
	return s.request("PATCH", fmt.Sprintf("/items/%d", id), &item, nil)
}

func (s *apiService) DeleteItem(id uint) error {
	return s.request("DELETE", fmt.Sprintf("/items/%d", id), nil, nil)
}

// TODO: implement set and get item data

func (s *apiService) SetItemData(itemId uint, data string) error {
	return nil
}

func (s *apiService) ItemData(itemId uint) (data string, err error) {
	return data, err
}
