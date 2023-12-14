package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestCars(s *testing.T) {
	router := NewRouter()

	s.Run("returns 200 OK", func(t *testing.T) {
		apitest.New().Handler(router).Put("/cars").
			JSON(`[{ "id": 1, "seats": 10 },{ "id": 2, "seats": 15 }]`).
			Expect(t).
			Status(http.StatusOK).
			Assert(expectEmptyBody).
			End()
	})

	s.Run("returns 400 Bad Request with broken body", func(t *testing.T) {
		apitest.New().Handler(router).Put("/cars").
			JSON(`[{ "id": 1, "seats": 10 },{ "id": 2, "seats": 15 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("returns 400 Bad Request with empty body", func(t *testing.T) {
		apitest.New().Handler(router).Put("/cars").
			JSON("").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("returns 400 Bad Request with wrong mime type", func(t *testing.T) {
		apitest.New().Handler(router).Put("/cars").
			Body(`[{ "id": 1, "seats": 10 },{ "id": 2, "seats": 15 }]`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("refuses cars with duplicated IDs", func(t *testing.T) {
		apitest.New().Handler(router).Put("/cars").
			JSON(`[{ "id": 1, "seats": 10 },{ "id": 1, "seats": 15 }]`).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(expectEmptyBody).
			End()
	})
}

func TestDropoff(s *testing.T) {
	router := NewRouter()
	addTestData(s, router)

	s.Run("returns 204 No Content", func(t *testing.T) {
		apitest.New().Handler(router).Post("/dropoff").
			FormData("ID", "1").
			Expect(t).
			Status(http.StatusNoContent).
			Assert(expectEmptyBody).
			End()
	})

	s.Run("returns 404 Not found when group is not found", func(t *testing.T) {
		apitest.New().Handler(router).Post("/dropoff").
			FormData("ID", "5").
			Expect(t).
			Status(http.StatusNotFound).
			Assert(expectEmptyBody).
			End()
	})

	s.Run("returns 400 Bad Request with broken body", func(t *testing.T) {
		apitest.New().Handler(router).Post("/dropoff").
			FormData("ID", "").
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(expectEmptyBody).
			End()
	})

	s.Run("returns 400 Bad Request with empty body", func(t *testing.T) {
		apitest.New().Handler(router).Post("/dropoff").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("returns 400 Bad Request with wrong mime type", func(t *testing.T) {
		apitest.New().Handler(router).Post("/dropoff").
			JSON(`{ "id": 1 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
}

func TestJourney(s *testing.T) {
	router := NewRouter()

	s.Run("returns 201 Accepted", func(t *testing.T) {
		apitest.New().Handler(router).Post("/journey").
			JSON(`{ "id": 1, "people": 2 }`).
			Expect(t).
			Status(http.StatusAccepted).
			End()
	})

	s.Run("returns 400 Bad Request with broken body", func(t *testing.T) {
		apitest.New().Handler(router).Post("/journey").
			JSON(`{ "id": 1, "people": 2`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("returns 400 Bad Request with empty body", func(t *testing.T) {
		apitest.New().Handler(router).Post("/journey").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("returns 400 Bad Request with wrong mime type", func(t *testing.T) {
		apitest.New().Handler(router).Post("/journey").
			Body(`{ "id": 1, "people": 2 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	s.Run("returns 500 when the group already exists", func(t *testing.T) {
		addTestData(t, router)

		apitest.New().Handler(router).Post("/journey").
			JSON(`{ "id": 1, "people": 2 }`).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(expectEmptyBody).
			End()
	})
}

func TestLocate(s *testing.T) {
	router := NewRouter()
	addTestData(s, router)

	s.Run("returns 200 OK", func(t *testing.T) {
		expectedResponse := jsonpath.Chain().Equal("id", float64(1)).Equal("seats", float64(10)).End()

		apitest.New().Handler(router).Post("/locate").
			FormData("ID", "1").
			Expect(t).
			Status(http.StatusOK).
			Assert(expectedResponse).
			End()
	})

	s.Run("returns 204 No Content when the group is waiting for a car", func(t *testing.T) {
		apitest.New().Handler(router).Post("/locate").
			FormData("ID", "3").
			Expect(t).
			Status(http.StatusNoContent).
			Assert(expectEmptyBody).
			End()
	})

	s.Run("returns 404 Not found when group is not found", func(t *testing.T) {
		apitest.New().Handler(router).Post("/locate").
			FormData("ID", "10").
			Expect(t).
			Status(http.StatusNotFound).
			Assert(expectEmptyBody).
			End()
	})

	s.Run("returns 400 Bad Request without ID", func(t *testing.T) {
		apitest.New().Handler(router).Post("/locate").
			FormData("ID", "").
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(expectEmptyBody).
			End()

		apitest.New().Handler(router).Post("/locate").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
}

func TestStatus(s *testing.T) {
	router := NewRouter()

	s.Run("returns 200 OK", func(t *testing.T) {
		apitest.New().Handler(router).Get("/status").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

func addTestData(t *testing.T, router *mux.Router) {
	apitest.New().Handler(router).Put("/cars").
		JSON(`[{ "id": 1, "seats": 10 },{ "id": 2, "seats": 15 }]`).
		Expect(t).
		Status(http.StatusOK).
		End()

	createJourney := func(id int, people int) {
		apitest.New().Handler(router).Post("/journey").
			JSON(fmt.Sprintf(`{ "id": %d, "people": %d }`, id, people)).
			Expect(t).
			Status(http.StatusAccepted).
			End()
	}

	createJourney(1, 2)
	createJourney(2, 15)
	createJourney(3, 50)
}

func expectEmptyBody(res *http.Response, req *http.Request) error {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return errors.New("cannot read body")
	}

	if len(body) != 0 {
		return fmt.Errorf("I expect an empty body as response, got: %s", body)
	}

	return nil
}
