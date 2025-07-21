package database

import (
	"os"
	"testing"

	"github.com/echo-webkom/ludo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func createTestDB(t *testing.T) *Database {
	tmpfile, err := os.CreateTemp("", "testdb-*.sqlite")
	assert.NoError(t, err)

	db := NewSQLite(tmpfile.Name())

	t.Cleanup(func() {
		os.Remove(tmpfile.Name())
	})

	// Auto migrate schema
	err = db.db.AutoMigrate(&model.Board{}, &model.Item{}, &model.User{})
	assert.NoError(t, err)

	return db
}

func TestUserMethods(t *testing.T) {
	db := createTestDB(t)

	// Create user
	user := model.User{DisplayName: "Test User", GithubUsername: "testuser"}
	id, err := db.NewUser(user)
	assert.NoError(t, err)

	// Get user
	u, err := db.User(id)
	assert.NoError(t, err)
	assert.Equal(t, "Test User", u.DisplayName)

	// Get all users
	users, err := db.Users()
	assert.NoError(t, err)
	assert.Len(t, users, 1)

	// Delete user
	err = db.DeleteUser(id)
	assert.NoError(t, err)

	// Confirm user deleted
	_, err = db.User(id)
	assert.Error(t, err)
}

func TestBoardMethods(t *testing.T) {
	db := createTestDB(t)

	// Create board
	board := model.Board{Title: "Board 1", RepoURL: "https://github.com/example/repo"}
	_, err := db.NewBoard(board)
	assert.NoError(t, err)

	// Get boards
	boards, err := db.Boards()
	assert.NoError(t, err)
	assert.Len(t, boards, 1)

	// Get by ID
	got, err := db.Board(boards[0].ID)
	assert.NoError(t, err)
	assert.Equal(t, "Board 1", got.Title)

	// Update board
	got.Title = "Updated"
	err = db.UpdateBoard(got.ID, got)
	assert.NoError(t, err)

	// Get updated board
	got2, err := db.Board(got.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated", got2.Title)

	// Delete board
	err = db.DeleteBoard(got.ID)
	assert.NoError(t, err)
}

func TestItemMethods(t *testing.T) {
	db := createTestDB(t)

	// Create board
	board := model.Board{Title: "Board", RepoURL: "url"}
	id, err := db.NewBoard(board)
	assert.NoError(t, err)
	boards, _ := db.Boards()

	// Create item
	item := model.Item{BoardID: boards[0].ID, Title: "Item 1", Status: model.StatusBacklog}
	itemID, err := db.NewItem(id, item)
	assert.NoError(t, err)

	// Get item
	got, err := db.Item(itemID)
	assert.NoError(t, err)
	assert.Equal(t, "Item 1", got.Title)

	// Change status
	err = db.ChangeItemStatus(itemID, model.StatusInProgress)
	assert.NoError(t, err)

	// Get items with status
	items, err := db.GetAllItemsWithStatus(boards[0].ID, model.StatusInProgress)
	assert.NoError(t, err)
	assert.Len(t, items, 1)

	// Set/Get item data
	err = db.SetItemData(itemID, "some-data")
	assert.NoError(t, err)

	data, err := db.GetItemData(itemID)
	assert.NoError(t, err)
	assert.Equal(t, "some-data", data)

	// Delete item
	assert.NoError(t, db.DeleteItem(itemID))
}

func TestBoardUsers(t *testing.T) {
	db := createTestDB(t)

	// Create user and board
	user := model.User{DisplayName: "A"}
	userId, _ := db.NewUser(user)
	board := model.Board{Title: "Board"}

	_, err := db.NewBoard(board)
	assert.NoError(t, err)

	boards, _ := db.Boards()

	// Add user to board
	assert.NoError(t, db.AddUserToBoard(boards[0].ID, userId))

	// Get board users
	users, err := db.BoardUsers(boards[0].ID)
	assert.NoError(t, err)
	assert.Len(t, users, 1)

	// Remove user from board
	assert.NoError(t, db.RemoveUserFromBoard(boards[0].ID, userId))

	// Check no users left
	users, err = db.BoardUsers(boards[0].ID)
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestGetBoardItemsByStatus(t *testing.T) {
	db := createTestDB(t)

	board := model.Board{Title: "Board"}
	_, err := db.NewBoard(board)
	assert.NoError(t, err)

	boards, _ := db.Boards()

	id := boards[0].ID

	// Create multiple items with different statuses
	db.NewItem(id, model.Item{BoardID: boards[0].ID, Title: "Item 1", Status: model.StatusBacklog})
	db.NewItem(id, model.Item{BoardID: boards[0].ID, Title: "Item 2", Status: model.StatusInProgress})
	db.NewItem(id, model.Item{BoardID: boards[0].ID, Title: "Item 3", Status: model.StatusBacklog})

	// Fetch items with StatusBacklog
	items, err := db.BoardItemsWithStatus(boards[0].ID, model.StatusBacklog)
	assert.NoError(t, err)
	assert.Len(t, items, 2)
}
