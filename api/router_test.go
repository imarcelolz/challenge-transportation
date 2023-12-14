package api

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

func TestRouterHandler(t *testing.T) {
	t.Run("PUT /cars", assertRoute(NewRouter(), "PUT", "/cars"))
	t.Run("GET /status", assertRoute(NewRouter(), "GET", "/status"))
	t.Run("POST /dropoff", assertRoute(NewRouter(), "POST", "/dropoff"))
	t.Run("POST /journey", assertRoute(NewRouter(), "POST", "/journey"))
	t.Run("POST /locate", assertRoute(NewRouter(), "POST", "/locate"))
}

func assertRoute(router *mux.Router, method string, path string) func(t *testing.T) {
	return func(t *testing.T) {
		routeMatch := &mux.RouteMatch{}

		assert.True(t, router.Match(&http.Request{Method: method, URL: &url.URL{Path: path}}, routeMatch))
	}
}
