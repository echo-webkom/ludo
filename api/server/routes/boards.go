package routes

import (
	"net/http"
	"strconv"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/rest"
	"github.com/go-chi/chi/v5"
)

func BoardsHandler(db *database.Database) chi.Router {
	r := chi.NewRouter()

	// Get all boards
	r.Get("/", rest.Handler(func(r rest.Request) int {
		if boards, err := db.GetAllBoards(); err == nil {
			return r.RespondJSON(&boards)
		}
		return http.StatusInternalServerError
	}))

	// Create board
	r.Post("/", rest.Handler(func(r rest.Request) int {
		var board database.Board
		if err := r.ParseJSON(&board); err != nil {
			return http.StatusBadRequest
		}

		id, err := db.CreateBoard(board)
		if err != nil {
			return http.StatusInternalServerError
		}

		return r.RespondJSON(&database.ID{ID: id})
	}))

	r.Route("/{id}", func(r chi.Router) {
		r.Use(idMiddleware)

		// Get board by id
		r.Get("/", rest.Handler(func(r rest.Request) int {
			if board, err := db.GetBoardById(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&board)
			}
			return http.StatusInternalServerError
		}))

		// Update board
		r.Patch("/", rest.Handler(func(r rest.Request) int {
			var board database.Board
			if err := r.ParseJSON(&board); err != nil {
				return http.StatusBadRequest
			}

			if err := db.UpdateBoard(board, r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete baord
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := db.DeleteBoard(r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))

		// Get all items in board
		r.Get("/items", rest.Handler(func(r rest.Request) int {
			if items, err := db.GetBoardItems(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&items)
			}
			return http.StatusInternalServerError
		}))

		// Get all items with status
		r.Get("/status/{status}/items", rest.Handler(func(r rest.Request) int {
			status, err := strconv.Atoi(r.R.PathValue("status"))
			if err != nil {
				return http.StatusBadRequest
			}

			items, err := db.GetBoardItemsByStatus(r.ContextValue(idKey).(uint), database.Status(status))
			if err != nil {
				return http.StatusInternalServerError
			}

			return r.RespondJSON(&items)
		}))

		// Get all users in board
		r.Get("/users", rest.Handler(func(r rest.Request) int {
			if users, err := db.GetBoardUsers(r.ContextValue(idKey).(uint)); err == nil {
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

			if err := db.AddUserToBoard(r.ContextValue(idKey).(uint), uint(userId)); err != nil {
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

			if err := db.RemoveUserFromBoard(r.ContextValue(idKey).(uint), uint(userId)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))
	})

	return r
}
