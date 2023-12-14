package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/imarcelolz/transportation-challenge/api/middleware"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	controller := NewPoolController()

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.HttpErrorMiddleware)
	router.HandleFunc("/status", status).Methods("GET")

	jsonApi := middleware.NewMimeTypeMiddleware("application/json")
	router.HandleFunc("/cars", jsonApi(controller.cars)).Methods("PUT")
	router.HandleFunc("/journey", jsonApi(controller.journey)).Methods("POST")

	formEncoded := middleware.NewMimeTypeMiddleware("application/x-www-form-urlencoded")
	router.HandleFunc("/dropoff", formEncoded(controller.dropoff)).Methods("POST")
	router.HandleFunc("/locate", formEncoded(controller.locate)).Methods("POST")

	return router
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
