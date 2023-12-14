package api

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestDecodeJSON(t *testing.T) {
	t.Run("decodes a json body", func(t *testing.T) {
		body := `{"id": 1, "name": "test"}`
		target := decodeJSON[testStruct](io.NopCloser(strings.NewReader(body)))

		assert.Equal(t, target.Id, 1)
		assert.Equal(t, target.Name, "test")
	})

	t.Run("panics if the body is not a valid JSON", func(t *testing.T) {
		body := `{"id": 1, "name": "`

		assert.Panics(t, func() { decodeJSON[testStruct](io.NopCloser(strings.NewReader(body))) })
	})
}

func TestEncodeJSON(t *testing.T) {
	t.Run("encodes a json body", func(t *testing.T) {
		target := testStruct{Id: 1, Name: "test"}
		var b strings.Builder

		encodeJSON(&b, target)

		assert.JSONEq(t, b.String(), `{"id":1,"name":"test"}`)
	})

	t.Run("panics if the body is not a valid JSON", func(t *testing.T) {
		target := make(chan int)
		var b strings.Builder

		assert.Panics(t, func() { encodeJSON(&b, target) })
	})
}
