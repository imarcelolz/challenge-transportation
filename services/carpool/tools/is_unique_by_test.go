package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUniqueBey(t *testing.T) {
	t.Run("returns true if there are no duplicated items", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}

		assert.True(t, IsUniqueBy(input, func(i int) int { return i }))
	})

	t.Run("returns false when there are duplicated items", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 5}

		assert.False(t, IsUniqueBy(input, func(i int) int { return i }))
	})
}
