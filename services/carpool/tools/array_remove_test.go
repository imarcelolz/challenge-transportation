package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayRemove(t *testing.T) {
	t.Run("removes an item from an array", func(t *testing.T) {
		target := []int{1, 2, 3, 4, 5}
		result := ArrayRemove(target, 3)

		assert.Equal(t, len(result), 4)
		assert.Equal(t, result[2], 4)
	})

	t.Run("when item is not found, returns the original array", func(t *testing.T) {
		target := []int{1, 2, 3, 4, 5}
		result := ArrayRemove(target, 6)

		assert.Equal(t, len(result), 5)
		assert.Equal(t, result[2], 3)
	})
}
