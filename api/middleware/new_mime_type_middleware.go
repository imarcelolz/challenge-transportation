package middleware

import (
	"fmt"
	"net/http"
)

type HttpHandler = func(http.ResponseWriter, *http.Request)

func NewMimeTypeMiddleware(mimeType string) func(next HttpHandler) HttpHandler {
	return func(next HttpHandler) HttpHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Type") != mimeType {
				http.Error(w, fmt.Sprintf("Server requires %s mime type", mimeType), http.StatusBadRequest)
				return
			}

			next(w, r)
		}
	}
}
