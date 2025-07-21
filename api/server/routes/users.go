package routes

import (
	"net/http"

	"github.com/echo-webkom/ludo/api/rest"
	"github.com/echo-webkom/ludo/pkg/model"
	"github.com/echo-webkom/ludo/pkg/service"
	"github.com/go-chi/chi/v5"
)

func UsersHandler(s service.LudoService) chi.Router {
	r := chi.NewRouter()

	// Get all users
	r.Get("/", rest.Handler(func(r rest.Request) int {
		if users, err := s.Users(); err == nil {
			return r.RespondJSON(&users)
		}
		return http.StatusInternalServerError
	}))

	// Create user
	r.Post("/", rest.Handler(func(r rest.Request) int {
		var user model.User
		if err := r.ParseJSON(&user); err != nil {
			return http.StatusBadRequest
		}

		id, err := s.NewUser(user)
		if err != nil {
			return http.StatusInternalServerError
		}

		return r.RespondJSON(&model.ID{ID: id})
	}))

	r.Route("/{id}", func(r chi.Router) {
		r.Use(idMiddleware)

		// Get user by id
		r.Get("/", rest.Handler(func(r rest.Request) int {
			if user, err := s.User(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&user)
			}
			return http.StatusInternalServerError
		}))

		// Update user
		r.Patch("/", rest.Handler(func(r rest.Request) int {
			var user model.User
			if err := r.ParseJSON(&user); err != nil {
				return http.StatusBadRequest
			}

			if err := s.UpdateUser(r.ContextValue(idKey).(uint), user); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete user by id
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := s.DeleteUser(r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))
	})

	return r
}
