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

func usersRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// service.GetAllUser()
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
		}

		var user database.User
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}
		// service.CreateUser(User)
	})

	router.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")
		userId, err := strconv.Atoi(userIdString)

		if err != nil && userId < 0 {
			http.Error(w, "not a valid user id", http.StatusBadRequest)
			return
		}

		// service.GetUserById(userId)

	})

	router.Delete("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")

		userId, err := strconv.Atoi(userIdString)
		if err != nil && userId < 0 {
			http.Error(w, "not a valid user id", http.StatusBadRequest)
			return
		}
		// service.DeleteUserById(userId)
	})

	router.Patch("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")

		userId, err := strconv.Atoi(userIdString)
		if err != nil && userId < 0 {
			http.Error(w, "not a valid user id", http.StatusBadRequest)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
		}

		var user database.User
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}
		// service.updateUser(user)

	})
	return router
}


func itemRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// services.getAllItems()
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll((r.Body))
		if err != nil {
			http.Error(w, "invalid item data", http.StatusBadGateway)
		}

		var item database.Item
		if err := json.Unmarshal(b, &item); err != nil {
			http.Error(w, "invalid item data", http.StatusBadRequest)
			return
		}

		// service.createItem(item)
	})

	router.Delete("/{itemId}", func(w http.ResponseWriter, r *http.Request) {
		
	})


	return router
}


