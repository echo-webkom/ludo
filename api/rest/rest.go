package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type HandlerFunc func(r Request) int

// Request wraps http.ResponseWriter and http.Request and adds multiple common
// utilities for parsing request data, writing encoded responses etc.
type Request struct {
	W http.ResponseWriter
	R *http.Request
}

// Handler wraps http.HandlerFunc and provides a Request object with common utility functions.
func Handler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := f(Request{
			W: w,
			R: r,
		})

		if code != http.StatusOK {
			w.WriteHeader(code)
		}
	}
}

// RespondJSON writes a json response by encoding v. Returns status 200, or 500 on error.
func (r Request) RespondJSON(v any) int {
	defer r.R.Body.Close()
	if err := json.NewEncoder(r.W).Encode(v); err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func (r Request) RespondString(s string) int {
	if _, err := r.W.Write([]byte(s)); err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

// Parses data from request body as json into v.
func (r Request) ParseJSON(v any) error {
	defer r.R.Body.Close()
	return json.NewDecoder(r.R.Body).Decode(v)
}

func (r Request) PathUint(name string) (uint, error) {
	n, err := strconv.Atoi(r.R.PathValue(name))
	return uint(n), err
}

func (r Request) ContextValue(key any) any {
	return r.R.Context().Value(key)
}
