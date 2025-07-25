package routes

import (
	"io"
	"net/http"

	"github.com/echo-webkom/ludo/api/rest"
	"github.com/echo-webkom/ludo/pkg/model"
	"github.com/echo-webkom/ludo/pkg/service"
	"github.com/go-chi/chi/v5"
)

func ItemsHandler(s service.LudoService) chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Use(idMiddleware)

		// Get item by id
		r.Get("/", rest.Handler(func(r rest.Request) int {
			if item, err := s.Item(r.ContextValue(idKey).(uint)); err == nil {
				return r.RespondJSON(&item)
			}
			return http.StatusInternalServerError
		}))

		// Update item
		r.Patch("/", rest.Handler(func(r rest.Request) int {
			var item model.Item
			if err := r.ParseJSON(&item); err != nil {
				return http.StatusBadRequest
			}

			if err := s.UpdateItem(r.ContextValue(idKey).(uint), item); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))

		// Delete item by id
		r.Delete("/", rest.Handler(func(r rest.Request) int {
			if err := s.DeleteItem(r.ContextValue(idKey).(uint)); err != nil {
				return http.StatusInternalServerError
			}
			return http.StatusOK
		}))

		// Get item data
		r.Get("/data", rest.Handler(func(r rest.Request) int {
			if item, err := s.Item(r.ContextValue(idKey).(uint)); err == nil {
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
			if err := s.SetItemData(id, string(data)); err != nil {
				return http.StatusInternalServerError
			}

			return http.StatusOK
		}))
	})

	return r
}
