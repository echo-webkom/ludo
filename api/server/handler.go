package server

import (
	"encoding/json"
	"net/http"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/go-chi/chi/v5"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}

func JSON(w http.ResponseWriter, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func usersHandler(db *database.Database) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := db.GetAllUsers()
		if err != nil {
			http.Error(w, "failed to get users", http.StatusInternalServerError)
			return
		}

		JSON(w, &users)
	})

	return r
}
