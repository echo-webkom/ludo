package routes

import (
	"io"
	"net/http"

	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/rest"
	"github.com/go-chi/chi/v5"
)

func ItemsHandler(db *database.Database) chi.Router {
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
			if item, err := db.GetItemById(r.ContextValue(idKey).(uint)); err == nil {
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

			if err := db.UpdateItem(item, r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete item by id
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := db.DeleteItemByID(r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))

		// Get item data
		r.Get("/data", rest.Handler(func(r rest.Request) int {
			if item, err := db.GetItemById(r.ContextValue(idKey).(uint)); err == nil {
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

			id := r.ContextValue(idKey).(uint)
			if err := db.SetItemData(id, string(data)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))
	})

	return r
}
