package database

import (
	"os"
	"testing"
)

func assert(t *testing.T, v bool, msg string, args ...any) {
	if !v {
		t.Fatalf(msg, args...)
	}
}

func assertNoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestDatabase(t *testing.T) {
	filename := "test.db"
	db := NewSQLite(filename)
	defer os.Remove(filename)

	t.Run("Create and get user", func(t *testing.T) {
		id, err := db.CreateUser(User{
			DisplayName:    "John",
			GithubUsername: "johnpork_",
		})
		assertNoErr(t, err)

		user, err := db.GetUserById(id)
		assertNoErr(t, err)

		assert(t, user.ID == id, "id does not match")
		assert(t, user.DisplayName == "John", "incorrect display name, got %s", user.DisplayName)
		assert(t, user.GithubUsername == "johnpork_", "incorrect github name, got %s", user.GithubUsername)
	})

	t.Run("Create board and add items", func(t *testing.T) {
		userID, err := db.CreateUser(User{
			DisplayName:    "Alice",
			GithubUsername: "alicehub",
		})
		assertNoErr(t, err)

		board := Board{}
		assertNoErr(t, db.db.Create(&board).Error)

		item := Item{
			BoardID:         board.ID,
			RepoName:        "example-repo",
			RepoURL:         "https://github.com/example/repo",
			Title:           "Initial Task",
			Description:     "Setup project",
			ConnectedBranch: "main",
			Status:          StatusInProgress,
			CreatorID:       userID,
			AssigneeID:      userID,
		}
		_, err = db.CreateItem(item)
		assertNoErr(t, err)

		items, err := db.GetAllItems()
		assertNoErr(t, err)
		assert(t, len(items) == 1, "expected 1 item, got %d", len(items))

		itemFromDB, err := db.GetItemById(items[0].ID)
		assertNoErr(t, err)
		assert(t, itemFromDB.Title == "Initial Task", "unexpected title, got %s", itemFromDB.Title)
	})

	t.Run("Get all users", func(t *testing.T) {
		users, err := db.GetAllUsers()
		assertNoErr(t, err)
		assert(t, len(users) >= 2, "expected at least 2 users, got %d", len(users))
	})

	t.Run("Get all items with status", func(t *testing.T) {
		board := Board{}
		assertNoErr(t, db.db.Create(&board).Error)

		item := Item{
			BoardID: board.ID,
			Title:   "Status Test",
			Status:  StatusInReview,
		}
		_, err := db.CreateItem(item)
		assertNoErr(t, err)

		items, err := db.GetAllItemsWithStatus(board.ID, StatusInReview)
		assertNoErr(t, err)
		assert(t, len(items) > 0, "expected at least 1 item with status 2")
	})

	t.Run("Set item status", func(t *testing.T) {
		items, _ := db.GetAllItems()
		item := items[0]

		err := db.ChangeItemStatus(item.ID, StatusClosed)
		assertNoErr(t, err)

		movedItem, err := db.GetItemById(item.ID)
		assertNoErr(t, err)
		assert(t, movedItem.Status == StatusClosed, "item was not set to status 3")
	})

	t.Run("Change item title", func(t *testing.T) {
		items, _ := db.GetAllItems()
		item := items[0]

		err := db.ChangeItemTitle(item.ID, "Updated Title")
		assertNoErr(t, err)

		updatedItem, err := db.GetItemById(item.ID)
		assertNoErr(t, err)
		assert(t, updatedItem.Title == "Updated Title", "title not updated correctly")
	})

	t.Run("Change item description", func(t *testing.T) {
		items, _ := db.GetAllItems()
		item := items[0]

		err := db.ChangeItemDescription(item.ID, "New Description")
		assertNoErr(t, err)

		updatedItem, err := db.GetItemById(item.ID)
		assertNoErr(t, err)
		assert(t, updatedItem.Description == "New Description", "description not updated correctly")
	})

	t.Run("Delete item", func(t *testing.T) {
		id, err := db.CreateItem(Item{})
		assertNoErr(t, err)

		err = db.DeleteItemByID(id)
		assertNoErr(t, err)

		_, err = db.GetItemById(id)
		assert(t, err != nil, "item should have been deleted")
	})

	t.Run("Delete user", func(t *testing.T) {
		id, err := db.CreateUser(User{})
		assertNoErr(t, err)

		err = db.DeleteUserById(id)
		assertNoErr(t, err)

		_, err = db.GetUserById(id)
		assert(t, err != nil, "user should have been deleted")
	})

	t.Run("Get and set item data", func(t *testing.T) {
		id, err := db.CreateItem(Item{})
		assertNoErr(t, err)

		str := "Hello World"
		assertNoErr(t, db.SetItemData(id, str))

		data, err := db.GetItemData(id)
		assertNoErr(t, err)
		assert(t, data == str, "data didnt match expected value")
	})
}
