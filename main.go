package main

import (
	"net/http"
	"os"

	"github.com/imarcelolz/transportation-challenge/api"
)

func main() {
	http.ListenAndServe(port(), api.NewRouter())
}

func port() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9091"
	}

	println("Api Started - Listening on port http://127.0.0.1:" + port)
	return ":" + port
}
