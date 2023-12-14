package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/imarcelolz/transportation-challenge/api/middleware"
)

func decodeJSON[T any](body io.ReadCloser) T {
	var target T

	err := json.NewDecoder(body).Decode(&target)

	if err != nil {
		panic(&middleware.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
	}

	return target
}

func encodeJSON[T any](w io.Writer, data T) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(&middleware.HTTPError{Code: http.StatusInternalServerError, Message: "Cannot serialize response"})
	}
}

func respondToJson[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	encodeJSON[T](w, data)
}
