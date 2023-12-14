package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexAdd(t *testing.T) {
	index := createTarget()

	expected := newTestStruct(5, 6)
	index.Add(expected)

	assert.Equal(t, index.Get(5), expected)
	assert.Equal(t, index.SecondaryIndex(6)[0], expected)
}

func TestIndexRemove(t *testing.T) {
	t.Run("removes the item from the index", func(t *testing.T) {
		index := createTarget()

		expected := newTestStruct(5, 6)
		index.Add(expected)

		index.Remove(expected)

		assert.Nil(t, index.Get(5))
		assert.Empty(t, index.SecondaryIndex(6))
	})

	t.Run("does not remove the item from the index when it does not exist", func(t *testing.T) {
		assert.NotPanics(t, func() {
			createTarget().Remove(newTestStruct(5, 6))
		})
	})
}

func TestIndexGet(t *testing.T) {
	t.Run("returns the indexed item", func(t *testing.T) {
		index := createTarget()
		expected := newTestStruct(1, 2)

		index.Add(expected)

		assert.Equal(t, index.Get(1), expected)
	})

	t.Run("returns nil when the index does not exist", func(t *testing.T) {
		assert.Nil(t, createTarget().Get(1))
	})
}

func TestSecondaryIndex(t *testing.T) {
	t.Run("returns the indexed items", func(t *testing.T) {
		index := createTarget()
		expected := newTestStruct(1, 2)

		index.Add(expected)

		assert.Equal(t, index.SecondaryIndex(2), []*testStruct{expected})
	})

	t.Run("returns an empty slice when the index does not exist", func(t *testing.T) {
		assert.Equal(t, createTarget().SecondaryIndex(1), []*testStruct{})
	})
}

func TestKeys(t *testing.T) {
	index := createTarget()

	index.Add(newTestStruct(2, 2))
	index.Add(newTestStruct(1, 3))

	assert.Equal(t, index.Keys(), []int{1, 2})
}

func TestSecondaryKeys(t *testing.T) {
	index := createTarget()

	index.Add(newTestStruct(1, 4))
	index.Add(newTestStruct(2, 3))

	assert.Equal(t, index.SecondaryKeys(), []int{3, 4})
}

type testStruct struct {
	Id     int
	Length int
}

func newTestStruct(id int, length int) *testStruct {
	return &testStruct{Id: id, Length: length}
}

func createTarget() *Index[*testStruct, int, int] {
	index := NewIndex[*testStruct, int, int](
		func(i *testStruct) int { return i.Id },
		func(i *testStruct) int { return i.Length },
	)

	return index
}
