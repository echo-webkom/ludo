package routes

import (
	"net/http"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/rest"
	"github.com/go-chi/chi/v5"
)

func UsersHandler(db *database.Database) chi.Router {
	r := chi.NewRouter()

	// Get all users
	r.Get("/", rest.Handler(func(r rest.Request) int {
		if users, err := db.GetAllUsers(); err == nil {
			return r.RespondJSON(&users)
		}
		return http.StatusInternalServerError
	}))

	// Create user
	r.Post("/", rest.Handler(func(r rest.Request) int {
		var user database.User
		if err := r.ParseJSON(&user); err != nil {
			return http.StatusBadRequest
		}

		id, err := db.CreateUser(user)
		if err != nil {
			return http.StatusInternalServerError
		}

		return r.RespondJSON(&database.ID{ID: id})
	}))

	r.Route("/{id}", func(r chi.Router) {
		r.Use(idMiddleware)

		// Get user by id
		r.Get("/", rest.Handler(func(r rest.Request) int {
			if user, err := db.GetUserById(r.ContextValue("id").(uint)); err == nil {
				return r.RespondJSON(&user)
			}
			return http.StatusInternalServerError
		}))

		// Update user
		r.Patch("/", rest.Handler(func(r rest.Request) int {
			var user database.User
			if err := r.ParseJSON(&user); err != nil {
				return http.StatusBadRequest
			}

			if err := db.UpdateUser(user, r.ContextValue("id").(uint)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete user by id
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := db.DeleteUserById(r.ContextValue("id").(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))
	})

	return r
}
