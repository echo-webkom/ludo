package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/go-chi/chi/v5"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}

func respondJSON(w http.ResponseWriter, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getJSON(r *http.Request, v any) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	return json.Unmarshal(b, v)
}

func usersHandler(db *database.Database) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetAllUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &users)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var user database.User
		if err := getJSON(r, &user); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		id, err := db.CreateUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &database.ID{ID: id})
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}

		user, err := db.GetUserById(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &user)
	})

	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		var user database.User
		if err := getJSON(r, &user); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		// err := db.UpdateUser()
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}

		if err := db.DeleteUserById(uint(id)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	return r
}

func itemsHandler(db *database.Database) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		items, err := db.GetAllItems()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &items)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var item database.Item
		if err := getJSON(r, &item); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		id, err := db.CreateItem(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &database.ID{ID: id})
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}

		item, err := db.GetItemById(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &item)
	})

	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		var item database.Item
		if err := getJSON(r, &item); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		// err := db.UpdateItem()
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}

		if err := db.DeleteItemByID(uint(id)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	r.Get("/{id}/data", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}

		item, err := db.GetItemById(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(item.Data))
	})

	r.Patch("/{id}/data", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := db.SetItemData(uint(id), string(data)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	return r
}

func boardsHandler(db *database.Database) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		boards, err := db.GetAllBoards()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		respondJSON(w, &boards)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var board database.Board
		if err := getJSON(r, &board); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		id, err := db.CreateBoard(board)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &database.ID{ID: id})
	})

	r.Get("/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		boardId, err := strconv.Atoi(r.PathValue("boardId"))
		if err != nil {
			http.Error(w, "bad boardId", http.StatusBadRequest)
			return
		}

		board, err := db.GetBoardById(uint(boardId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &board)
	})

	r.Patch("/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		boardId, err := strconv.Atoi(r.PathValue("boardId"))
		if err != nil {
			http.Error(w, "bad boardId", http.StatusBadRequest)
			return
		}

		var board database.Board
		if err := getJSON(r, &board); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		if err := db.UpdateBoard(board, uint(boardId)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	r.Delete("/{boardId}", func(w http.ResponseWriter, r *http.Request) {
		boardId, err := strconv.Atoi(r.PathValue("boardId"))
		if err != nil {
			http.Error(w, "bad boardId", http.StatusBadRequest)
			return
		}

		if err := db.DeleteBoard(uint(boardId)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	r.Get("/{boardId}/users", func(w http.ResponseWriter, r *http.Request) {
		boardId, err := strconv.Atoi(r.PathValue("boardId"))
		if err != nil {
			http.Error(w, "bad boardId", http.StatusBadRequest)
			return
		}

		users, err := db.GetBoardUsers(uint(boardId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &users)
	})

	r.Get("/{boardId}/items", func(w http.ResponseWriter, r *http.Request) {
		boardId, err := strconv.Atoi(r.PathValue("boardId"))
		if err != nil {
			http.Error(w, "bad boardId", http.StatusBadRequest)
			return
		}

		items, err := db.GetBoardItems(uint(boardId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &items)
	})

	r.Post("/{boardId}/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		boardId, err1 := strconv.Atoi(r.PathValue("boardId"))
		userId, err2 := strconv.Atoi(r.PathValue("userId"))
		if err1 != nil || err2 != nil {
			http.Error(w, "bad id(s)", http.StatusBadRequest)
			return
		}

		if err := db.AddUserToBoard(uint(boardId), uint(userId)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	r.Delete("/{boardId}/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		boardId, err1 := strconv.Atoi(r.PathValue("boardId"))
		userId, err2 := strconv.Atoi(r.PathValue("userId"))
		if err1 != nil || err2 != nil {
			http.Error(w, "bad id(s)", http.StatusBadRequest)
			return
		}

		if err := db.RemoveUserFromBoard(uint(boardId), uint(userId)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	r.Get("/{boardId}/status/{status}/items", func(w http.ResponseWriter, r *http.Request) {
		boardId, err := strconv.Atoi(r.PathValue("boardId"))
		if err != nil {
			http.Error(w, "bad boardId", http.StatusBadRequest)
			return
		}

		status, err := strconv.Atoi(r.PathValue("status"))
		if err != nil {
			http.Error(w, "bad status", http.StatusBadRequest)
			return
		}

		items, err := db.GetBoardItemsByStatus(uint(boardId), database.Status(status))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respondJSON(w, &items)
	})

	return r
}
