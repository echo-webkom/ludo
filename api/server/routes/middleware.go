package routes

import (
	"context"
	"net/http"
	"strconv"
)

type ctxId string

const idKey ctxId = "id"

// Require 'id' path value to be integer. Adds to requests context. Responds
// with bad request and doesnt forward request if id is invalid.
func idMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), idKey, uint(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
