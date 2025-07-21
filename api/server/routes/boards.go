package routes

import (
	"net/http"
	"strconv"

	"github.com/echo-webkom/ludo/api/rest"
	"github.com/echo-webkom/ludo/pkg/model"
	"github.com/echo-webkom/ludo/pkg/service"
	"github.com/go-chi/chi/v5"
)

func BoardsHandler(s service.LudoService) chi.Router {
	r := chi.NewRouter()

	// Get all boards
	r.Get("/", rest.Handler(func(r rest.Request) int {
		if boards, err := s.Boards(); err == nil {
			return r.RespondJSON(&boards)
		}
		return http.StatusInternalServerError
	}))

	// Create board
	r.Post("/", rest.Handler(func(r rest.Request) int {
		var board model.Board
		if err := r.ParseJSON(&board); err != nil {
			return http.StatusBadRequest
		}

		id, err := s.NewBoard(board)
		if err != nil {
			return http.StatusInternalServerError
		}

		return r.RespondJSON(&model.ID{ID: id})
	}))

	r.Route("/{id}", func(r chi.Router) {
		r.Use(idMiddleware)

		// Get board by id
		r.Get("/", rest.Handler(func(r rest.Request) int {
			if board, err := s.Board(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&board)
			}
			return http.StatusInternalServerError
		}))

		// Update board
		r.Patch("/", rest.Handler(func(r rest.Request) int {
			var board model.Board
			if err := r.ParseJSON(&board); err != nil {
				return http.StatusBadRequest
			}

			if err := s.UpdateBoard(r.ContextValue(idKey).(uint), board); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete baord
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := s.DeleteBoard(r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))

		// Get all items in board
		r.Get("/items", rest.Handler(func(r rest.Request) int {
			if items, err := s.BoardItems(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&items)
			}
			return http.StatusInternalServerError
		}))

		// Create item
		r.Post("/items", rest.Handler(func(r rest.Request) int {
			var item model.Item
			if err := r.ParseJSON(&item); err != nil {
				return http.StatusBadRequest
			}

			boardId := r.ContextValue(idKey).(uint)

			// Check if board exists
			if _, err := s.Board(boardId); err != nil {
				return http.StatusNotFound
			}

			id, err := s.NewItem(boardId, item)
			if err != nil {
				return http.StatusInternalServerError
			}

			return r.RespondJSON(&model.ID{ID: id})
		}))

		// Get all items with status
		r.Get("/status/{status}/items", rest.Handler(func(r rest.Request) int {
			status, err := strconv.Atoi(r.R.PathValue("status"))
			if err != nil {
				return http.StatusBadRequest
			}

			items, err := s.BoardItemsWithStatus(r.ContextValue(idKey).(uint), model.Status(status))
			if err != nil {
				return http.StatusInternalServerError
			}

			return r.RespondJSON(&items)
		}))

		// Get all users in board
		r.Get("/users", rest.Handler(func(r rest.Request) int {
			if users, err := s.BoardUsers(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&users)
			}
			return http.StatusInternalServerError
		}))

		// Add user to board
		r.Post("/users/{userId}", rest.Handler(func(r rest.Request) int {
			userId, err := strconv.Atoi(r.R.PathValue("userId"))
			if err != nil {
				return http.StatusBadRequest
			}

			if err := s.AddUserToBoard(r.ContextValue(idKey).(uint), uint(userId)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Remove user from board
		r.Delete("/users/{userId}", rest.Handler(func(r rest.Request) int {
			userId, err := strconv.Atoi(r.R.PathValue("userId"))
			if err != nil {
				return http.StatusBadRequest
			}

			if err := s.RemoveUserFromBoard(r.ContextValue(idKey).(uint), uint(userId)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))
	})

	return r
}
