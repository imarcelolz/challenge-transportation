package middleware

import (
	"net/http"
)

func HttpErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(*HTTPError); ok {
				http.Error(w, err.Message, err.Code)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
