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

	r.Patch("/{id}/set/{status}", func(w http.ResponseWriter, r *http.Request) {
		var item database.Item
		if err := getJSON(r, &item); err != nil {
			http.Error(w, "bad request data", http.StatusBadRequest)
			return
		}

		id, err1 := strconv.Atoi(r.PathValue("id"))
		status, err2 := strconv.Atoi(r.PathValue("status"))
		if err1 != nil || err2 != nil {
			http.Error(w, "bad value", http.StatusBadRequest)
			return
		}

		if err := db.ChangeItemStatus(uint(id), database.Status(status)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	return r
}
