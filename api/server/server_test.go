package server_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/server"
	"github.com/echo-webkom/ludo/pkg/model"
)

func setupTestServer(t *testing.T) (*server.Server, func()) {
	tempFile, err := os.CreateTemp("", "testdb-*.sqlite")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	db := database.NewSQLite(tempFile.Name())

	cfg := &config.Config{Port: ":8080"}
	srv := server.New(cfg, db)

	cleanup := func() {
		db.Close()
		os.Remove(tempFile.Name())
	}

	return srv, cleanup
}

func makeRequest(t *testing.T, handler http.Handler, method, path string, body any) *httptest.ResponseRecorder {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}
		bodyReader = bytes.NewReader(b)
	}
	req := httptest.NewRequest(method, path, bodyReader)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}

func TestUserEndpoints(t *testing.T) {
	s, cleanup := setupTestServer(t)
	defer cleanup()

	// Create user
	user := model.User{DisplayName: "Test User"}
	resp := makeRequest(t, s, http.MethodPost, "/users", &user)
	if resp.Code != http.StatusOK {
		t.Fatalf("CreateUser failed, code %d", resp.Code)
	}
	var created model.ID
	json.NewDecoder(resp.Body).Decode(&created)

	// Get user
	resp = makeRequest(t, s, http.MethodGet, "/users/"+strconv.Itoa(int(created.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetUserById failed, code %d", resp.Code)
	}

	// Get all users
	resp = makeRequest(t, s, http.MethodGet, "/users", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetAllUsers failed, code %d", resp.Code)
	}

	// Delete user
	resp = makeRequest(t, s, http.MethodDelete, "/users/"+strconv.Itoa(int(created.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("DeleteUser failed, code %d", resp.Code)
	}
}

func TestItemEndpoints(t *testing.T) {
	s, cleanup := setupTestServer(t)
	defer cleanup()

	// Create item
	item := model.Item{Data: "Test Data"}
	resp := makeRequest(t, s, http.MethodPost, "/items", &item)
	if resp.Code != http.StatusOK {
		t.Fatalf("CreateItem failed, code %d", resp.Code)
	}
	var created model.ID
	json.NewDecoder(resp.Body).Decode(&created)

	// Get item
	resp = makeRequest(t, s, http.MethodGet, "/items/"+strconv.Itoa(int(created.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetItemById failed, code %d", resp.Code)
	}

	// Get all items
	resp = makeRequest(t, s, http.MethodGet, "/items", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetAllItems failed, code %d", resp.Code)
	}

	// Get item data
	resp = makeRequest(t, s, http.MethodGet, "/items/"+strconv.Itoa(int(created.ID))+"/data", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetItemData failed, code %d", resp.Code)
	}

	// Update item data
	resp = makeRequest(t, s, http.MethodPatch, "/items/"+strconv.Itoa(int(created.ID))+"/data", bytes.NewBufferString("Updated Data"))
	if resp.Code != http.StatusOK {
		t.Fatalf("SetItemData failed, code %d", resp.Code)
	}

	// Delete item
	resp = makeRequest(t, s, http.MethodDelete, "/items/"+strconv.Itoa(int(created.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("DeleteItem failed, code %d", resp.Code)
	}
}

func TestBoardEndpoints(t *testing.T) {
	s, cleanup := setupTestServer(t)
	defer cleanup()

	// Create board
	board := model.Board{Title: "Test Board"}
	resp := makeRequest(t, s, http.MethodPost, "/boards", &board)
	if resp.Code != http.StatusOK {
		t.Fatalf("CreateBoard failed, code %d", resp.Code)
	}
	var createdBoard model.ID
	json.NewDecoder(resp.Body).Decode(&createdBoard)

	// Get board
	resp = makeRequest(t, s, http.MethodGet, "/boards/"+strconv.Itoa(int(createdBoard.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetBoard failed, code %d", resp.Code)
	}

	// Get all boards
	resp = makeRequest(t, s, http.MethodGet, "/boards", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetAllBoards failed, code %d", resp.Code)
	}

	// Add user to board (assuming user creation tested before)
	user := model.User{DisplayName: "Board User"}
	resp = makeRequest(t, s, http.MethodPost, "/users", &user)
	if resp.Code != http.StatusOK {
		t.Fatalf("CreateUser failed, code %d", resp.Code)
	}
	var userID model.ID
	json.NewDecoder(resp.Body).Decode(&userID)

	resp = makeRequest(t, s, http.MethodPost, "/boards/"+strconv.Itoa(int(createdBoard.ID))+"/users/"+strconv.Itoa(int(userID.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("AddUserToBoard failed, code %d", resp.Code)
	}

	// Get board users
	resp = makeRequest(t, s, http.MethodGet, "/boards/"+strconv.Itoa(int(createdBoard.ID))+"/users", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetBoardUsers failed, code %d", resp.Code)
	}

	// Get board items
	resp = makeRequest(t, s, http.MethodGet, "/boards/"+strconv.Itoa(int(createdBoard.ID))+"/items", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetBoardItems failed, code %d", resp.Code)
	}

	// Get board items by status
	resp = makeRequest(t, s, http.MethodGet, "/boards/"+strconv.Itoa(int(createdBoard.ID))+"/status/1/items", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("GetBoardItemsByStatus failed, code %d", resp.Code)
	}

	// Remove user from board
	resp = makeRequest(t, s, http.MethodDelete, "/boards/"+strconv.Itoa(int(createdBoard.ID))+"/users/"+strconv.Itoa(int(userID.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("RemoveUserFromBoard failed, code %d", resp.Code)
	}

	// Delete board
	resp = makeRequest(t, s, http.MethodDelete, "/boards/"+strconv.Itoa(int(createdBoard.ID)), nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("DeleteBoard failed, code %d", resp.Code)
	}
}
