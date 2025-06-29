package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/go-chi/chi/v5"
)

type ResponseId struct {
	ID uint `json:"id"`
}

func JSON(w http.ResponseWriter, data any) {
	resp, err := json.Marshal(&data)
	if err != nil {
		http.Error(w, "parse error", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func getId(w http.ResponseWriter, id string) uint {
	itemId, err := strconv.Atoi(id)
	if err != nil || itemId < 0 {
		http.Error(w, "not valid item id", http.StatusBadRequest)
		return 0
	}

	return uint(itemId)
}

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}

func usersRouter(db database.Database) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetAllUsers()
		if err != nil {
			http.Error(w, "failed to get all users", http.StatusNotFound)
			return
		}
		JSON(w, users)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		var user database.User
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}

		id, err := db.CreateUser(user)
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		res := ResponseId{id}
		JSON(w, res)
	})

	router.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")
		userId, err := strconv.Atoi(userIdString)

		if err != nil || userId < 0 {
			http.Error(w, "not a valid user id", http.StatusBadRequest)
			return
		}

		user, err := db.GetUserById(uint(userId))
		if err != nil {
			http.Error(w, "could not find user", http.StatusNotFound)
			return
		}
		JSON(w, user)
	})

	router.Delete("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")

		userId, err := strconv.Atoi(userIdString)
		if err != nil || userId < 0 {
			http.Error(w, "not a valid user id", http.StatusBadRequest)
			return
		}

		if err := db.DeleteUserById(uint(userId)); err != nil {
			http.Error(w, "not valid user id", http.StatusNotFound)
			return
		}
	})

	router.Patch("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")

		userId := getId(w, userIdString)

		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		var user database.User
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}

		if err := db.UpdateUser(userId, user); err != nil {
			http.Error(w, "could not update user", http.StatusInternalServerError)
			return
		}

	})
	return router
}

func itemRouter(db database.Database) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		items, err := db.GetAllItems()
		if err != nil {
			http.Error(w, "could not find items", http.StatusNotFound)
			return
		}
		JSON(w, items)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "invalid item data", http.StatusBadGateway)
			return
		}

		var item database.Item
		if err := json.Unmarshal(b, &item); err != nil {
			http.Error(w, "invalid item data", http.StatusBadRequest)
			return
		}

		itemId, err := db.CreateItem(item)
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		res := ResponseId{itemId}
		JSON(w, res)
	})

	router.Get("/{itemId}", func(w http.ResponseWriter, r *http.Request) {
		itemIdString := r.PathValue("itemId")
		itemId := getId(w, itemIdString)

		item, err := db.GetItemById(uint(itemId))
		if err != nil {
			http.Error(w, "could not find item", http.StatusBadRequest)
			return
		}
		JSON(w, item)
	})

	router.Patch("/{itemId}", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Delete("/{itemId}", func(w http.ResponseWriter, r *http.Request) {
		itemIdString := r.PathValue("itemId")
		itemId := getId(w, itemIdString)

		if err := db.DeleteItemByID(uint(itemId)); err != nil {
			http.Error(w, "could not delet item", http.StatusBadRequest)
			return
		}
	})

	router.Patch("/{itemId}", func(w http.ResponseWriter, r *http.Request) {
		itemIdString := r.PathValue("itemId")
		itemId := getId(w, itemIdString)

		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "could not find user", http.StatusInternalServerError)
		}

		var item database.Item
		if err := json.Unmarshal(b, &item); err != nil {
			http.Error(w, "invalid item data", http.StatusBadRequest)
		}

		if err := db.UpdateItem(itemId, item); err != nil {
			http.Error(w, "could not update item", http.StatusInternalServerError)
		}
	})

	router.Patch("/{itemId}/move/{listId}", func(w http.ResponseWriter, r *http.Request) {
		itemIdString := r.PathValue("itemId")
		listIdString := r.PathValue("listId")

		itemId := getId(w, itemIdString)
		listId := getId(w, listIdString)

		if err := db.MoveItemToList(itemId, listId); err != nil {
			http.Error(w, "could not move item to list", http.StatusInternalServerError)
		}
	})

	return router
}

func boardsRouter(db database.Database) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		boards, err := db.GetAllBoards()
		if err != nil {
			http.Error(w, "failed to get all boards", http.StatusNotFound)
		}

		JSON(w, boards)

	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}
		
		var board database.Board
		if err := json.Unmarshal(b, &board); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}


		id, err := db.CreateBoard(board)
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		res := ResponseId{id}
		JSON(w, res)

	})

	router.Get("/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		boardIdString := r.PathValue("boardId")
		boardId := getId(w, boardIdString)	

		board, err := db.GetBoardById(boardId)
		if err != nil {
			http.Error(w, "failed to get board by id", http.StatusNotFound)
		}

		JSON(w, board)
	})

	router.Patch("/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		boardIdString := r.PathValue("boardId")
		boardId := getId(w, boardIdString)


		b, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "could not find user", http.StatusInternalServerError)
		}

		var board database.Board
		if err := json.Unmarshal(b, &board); err != nil {
			http.Error(w, "invalid item data", http.StatusBadRequest)
		}

		if err := db.UpdateBoard(boardId, board); err != nil {
			http.Error(w, "could not update user", http.StatusInternalServerError)
		}
	})

	router.Delete("/{boardId}", func(w http.ResponseWriter, r *http.Request) {

		boardIdString := r.PathValue("boardId")
		boardId := getId(w, boardIdString)

		if err := db.DeleteBoardById(boardId); err != nil {
			http.Error(w, "could not delete", http.StatusBadRequest)
		}
	})

	router.Get("{boardId}/users", func(w http.ResponseWriter, r *http.Request) {
		
	})

	router.Get("/{boardId}/items", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Get("/{boardId}/users/{userId}", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Delete("/{boardId}/users/{userId}", func(w http.ResponseWriter, r *http.Request) {

	})

	router.Get("/{boardId}/lists/{listId}/items", func(w http.ResponseWriter, r *http.Request) {

	})

	return router
}
