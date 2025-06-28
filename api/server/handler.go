package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/service"
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

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}

func usersRouter(service *service.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := service.GetAllUsers()
		if err != nil {
			http.Error(w, "failed to get all users", http.StatusNotFound)
			return
		}
		JSON(w, users)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		var user database.User
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}
		id, err := service.CreateUser(user)
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

		user, err := service.GetUserById(uint(userId))
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

		if err := service.DeleteUser(uint(userId)); err != nil {
			http.Error(w, "not valid user id", http.StatusNotFound)
			return
		}
	})

	router.Patch("/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userIdString := r.PathValue("userId")

		userId, err := strconv.Atoi(userIdString)
		if err != nil || userId < 0 {
			http.Error(w, "not a valid user id", http.StatusBadRequest)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		var user database.User
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, "invalid user data", http.StatusBadRequest)
			return
		}

		service.UpdateUser(uint(userId), user)

	})
	return router
}

func itemRouter(service *service.Service) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		items, err := service.GetAllItems()
		if err != nil {
			http.Error(w, "could not find items", http.StatusNotFound)
			return
		}
		JSON(w, items)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll((r.Body))
		if err != nil {
			http.Error(w, "invalid item data", http.StatusBadGateway)
			return
		}

		var item database.Item
		if err := json.Unmarshal(b, &item); err != nil {
			http.Error(w, "invalid item data", http.StatusBadRequest)
			return
		}

		itemId, err := service.CreateItem(item)
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}

		res := ResponseId{itemId}
		JSON(w, res)	
	})

	router.Delete("/{itemId}", func(w http.ResponseWriter, r *http.Request) {
		itemIdString := r.PathValue("itemId")
		itemId, err := strconv.Atoi(itemIdString)
		if err != nil || itemId < 0{
			http.Error(w, "not valid item id", http.StatusBadRequest)
		}
		
		if err := service.DeleteItem(uint(itemId)); err != nil {
			http.Error(w, "could not delet item", http.StatusBadRequest)
		}
		


	})

	return router
}
