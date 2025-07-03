package database

import (
	"os"
	"testing"

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
	err = db.db.AutoMigrate(&Board{}, &Item{}, &User{})
	assert.NoError(t, err)

	return db
}

func TestUserMethods(t *testing.T) {
	db := createTestDB(t)

	// Create user
	user := User{DisplayName: "Test User", GithubUsername: "testuser"}
	id, err := db.CreateUser(user)
	assert.NoError(t, err)

	// Get user
	u, err := db.GetUserById(id)
	assert.NoError(t, err)
	assert.Equal(t, "Test User", u.DisplayName)

	// Get all users
	users, err := db.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 1)

	// Delete user
	err = db.DeleteUserById(id)
	assert.NoError(t, err)

	// Confirm user deleted
	_, err = db.GetUserById(id)
	assert.Error(t, err)
}

func TestBoardMethods(t *testing.T) {
	db := createTestDB(t)

	// Create board
	board := Board{Title: "Board 1", RepoURL: "https://github.com/example/repo"}
	err := db.CreateBoard(board)
	assert.NoError(t, err)

	// Get boards
	boards, err := db.GetAllBoards()
	assert.NoError(t, err)
	assert.Len(t, boards, 1)

	// Get by ID
	got, err := db.GetBoardById(boards[0].ID)
	assert.NoError(t, err)
	assert.Equal(t, "Board 1", got.Title)

	// Update board
	got.Title = "Updated"
	err = db.UpdateBoard(got, got.ID)
	assert.NoError(t, err)

	// Get updated board
	got2, err := db.GetBoardById(got.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated", got2.Title)

	// Delete board
	err = db.DeleteBoard(got.ID)
	assert.NoError(t, err)
}

func TestItemMethods(t *testing.T) {
	db := createTestDB(t)

	// Create board
	board := Board{Title: "Board", RepoURL: "url"}
	assert.NoError(t, db.CreateBoard(board))
	boards, _ := db.GetAllBoards()

	// Create item
	item := Item{BoardID: boards[0].ID, Title: "Item 1", Status: StatusBacklog}
	itemID, err := db.CreateItem(item)
	assert.NoError(t, err)

	// Get item
	got, err := db.GetItemById(itemID)
	assert.NoError(t, err)
	assert.Equal(t, "Item 1", got.Title)

	// Change status
	err = db.ChangeItemStatus(itemID, StatusInProgress)
	assert.NoError(t, err)

	// Get items with status
	items, err := db.GetAllItemsWithStatus(boards[0].ID, StatusInProgress)
	assert.NoError(t, err)
	assert.Len(t, items, 1)

	// Set/Get item data
	err = db.SetItemData(itemID, "some-data")
	assert.NoError(t, err)

	data, err := db.GetItemData(itemID)
	assert.NoError(t, err)
	assert.Equal(t, "some-data", data)

	// Delete item
	assert.NoError(t, db.DeleteItemByID(itemID))
}

func TestBoardUsers(t *testing.T) {
	db := createTestDB(t)

	// Create user and board
	user := User{DisplayName: "A"}
	userId, _ := db.CreateUser(user)
	board := Board{Title: "Board"}
	assert.NoError(t, db.CreateBoard(board))
	boards, _ := db.GetAllBoards()

	// Add user to board
	assert.NoError(t, db.AddUserToBoard(boards[0].ID, userId))

	// Get board users
	users, err := db.GetBoardUsers(boards[0].ID)
	assert.NoError(t, err)
	assert.Len(t, users, 1)

	// Remove user from board
	assert.NoError(t, db.RemoveUserFromBoard(boards[0].ID, userId))

	// Check no users left
	users, err = db.GetBoardUsers(boards[0].ID)
	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestGetBoardItemsByStatus(t *testing.T) {
	db := createTestDB(t)

	board := Board{Title: "Board"}
	assert.NoError(t, db.CreateBoard(board))
	boards, _ := db.GetAllBoards()

	// Create multiple items with different statuses
	db.CreateItem(Item{BoardID: boards[0].ID, Title: "Item 1", Status: StatusBacklog})
	db.CreateItem(Item{BoardID: boards[0].ID, Title: "Item 2", Status: StatusInProgress})
	db.CreateItem(Item{BoardID: boards[0].ID, Title: "Item 3", Status: StatusBacklog})

	// Fetch items with StatusBacklog
	items, err := db.GetBoardItemsByStatus(boards[0].ID, StatusBacklog)
	assert.NoError(t, err)
	assert.Len(t, items, 2)
}
