package server

import (
	"io"
	"net/http"
	"strconv"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/rest"
	"github.com/go-chi/chi/v5"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}
}

func usersHandler(db *database.Database) chi.Router {
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

func itemsHandler(db *database.Database) chi.Router {
	r := chi.NewRouter()

	// Get all items
	r.Get("/", rest.Handler(func(r rest.Request) int {
		if items, err := db.GetAllItems(); err == nil {
			return r.RespondJSON(&items)
		}
		return http.StatusInternalServerError
	}))

	// Create item
	r.Post("/", rest.Handler(func(r rest.Request) int {
		var item database.Item
		if err := r.ParseJSON(&item); err != nil {
			return http.StatusBadRequest
		}

		id, err := db.CreateItem(item)
		if err != nil {
			return http.StatusInternalServerError
		}

		return r.RespondJSON(&database.ID{ID: id})
	}))

	r.Route("/{id}", func(r chi.Router) {
		r.Use(idMiddleware)

		// Get item by id
		r.Get("/", rest.Handler(func(r rest.Request) int {
			if item, err := db.GetItemById(r.ContextValue("id").(uint)); err == nil {
				return r.RespondJSON(&item)
			}
			return http.StatusInternalServerError
		}))

		// Update item
		r.Patch("/", rest.Handler(func(r rest.Request) int {
			var item database.Item
			if err := r.ParseJSON(&item); err != nil {
				return http.StatusBadRequest
			}

			if err := db.UpdateItem(item, r.ContextValue("id").(uint)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete item by id
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := db.DeleteItemByID(r.ContextValue("id").(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))

		// Get item data
		r.Get("/data", rest.Handler(func(r rest.Request) int {
			if item, err := db.GetItemById(r.ContextValue("id").(uint)); err == nil {
				return r.RespondString(item.Data)
			}
			return http.StatusInternalServerError
		}))

		// Set item data
		r.Patch("/data", rest.Handler(func(r rest.Request) int {
			data, err := io.ReadAll(r.R.Body)
			if err != nil {
				return http.StatusInternalServerError
			}

			id := r.ContextValue("id").(uint)
			if err := db.SetItemData(id, string(data)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))
	})

	return r
}

func boardsHandler(db *database.Database) chi.Router {
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
			if board, err := db.GetBoardById(r.ContextValue("id").(uint)); err == nil {
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

			if err := db.UpdateBoard(board, r.ContextValue("id").(uint)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete baord
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := db.DeleteBoard(r.ContextValue("id").(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))

		// Get all items in board
		r.Get("/items", rest.Handler(func(r rest.Request) int {
			if items, err := db.GetBoardItems(r.ContextValue("id").(uint)); err == nil {
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

			items, err := db.GetBoardItemsByStatus(r.ContextValue("id").(uint), database.Status(status))
			if err != nil {
				return http.StatusInternalServerError
			}

			return r.RespondJSON(&items)
		}))

		r.Route("/users", func(r chi.Router) {
			// Get all users in board
			r.Get("/", rest.Handler(func(r rest.Request) int {
				if users, err := db.GetBoardUsers(r.ContextValue("id").(uint)); err == nil {
					return r.RespondJSON(&users)
				}
				return http.StatusInternalServerError
			}))

			// Add user to board
			r.Post("/{userId}", rest.Handler(func(r rest.Request) int {
				userId, err := strconv.Atoi(r.R.PathValue("userId"))
				if err != nil {
					return http.StatusBadRequest
				}

				if err := db.AddUserToBoard(r.ContextValue("id").(uint), uint(userId)); err != nil {
					return http.StatusInternalServerError
				}

				return http.StatusOK
			}))

			// Remove user from board
			r.Delete("/{userId}", rest.Handler(func(r rest.Request) int {
				userId, err := strconv.Atoi(r.R.PathValue("userId"))
				if err != nil {
					return http.StatusBadRequest
				}

				if err := db.RemoveUserFromBoard(r.ContextValue("id").(uint), uint(userId)); err != nil {
					return http.StatusInternalServerError
				}

				return http.StatusOK
			}))
		})
	})

	return r
}
