package index

import (
	"sort"

	"github.com/imarcelolz/transportation-challenge/services/carpool/tools"
)

type keyConstraint interface{ int | string }
type keyBuilder[T comparable, I keyConstraint] func(T) I

type Index[T comparable, P keyConstraint, S keyConstraint] struct {
	primaryIndex   map[P]T
	secondaryIndex map[S][]T

	primaryKeyBuilder   keyBuilder[T, P]
	secondaryKeyBuilder keyBuilder[T, S]
}

func NewIndex[T comparable, P keyConstraint, S keyConstraint](primaryKeyBuilder keyBuilder[T, P], secondaryKeyBuilder keyBuilder[T, S]) *Index[T, P, S] {
	return &Index[T, P, S]{
		primaryIndex:        make(map[P]T),
		secondaryIndex:      make(map[S][]T),
		primaryKeyBuilder:   primaryKeyBuilder,
		secondaryKeyBuilder: secondaryKeyBuilder,
	}
}

func (i *Index[T, P, S]) Add(item T) {
	pk, sk := i.primaryKeyBuilder(item), i.secondaryKeyBuilder(item)

	secondaryIndex := i.SecondaryIndex(sk)

	i.primaryIndex[pk] = item
	i.secondaryIndex[sk] = append(secondaryIndex, item)
}

func (i *Index[T, P, S]) Remove(item T) {
	pk, sk := i.primaryKeyBuilder(item), i.secondaryKeyBuilder(item)

	secondaryIndex := i.SecondaryIndex(sk)

	delete(i.primaryIndex, pk)
	i.secondaryIndex[sk] = tools.ArrayRemove(secondaryIndex, item)
}

func (i *Index[T, P, S]) Get(index P) T {
	return i.primaryIndex[index]
}

func (i *Index[T, P, S]) SecondaryIndex(index S) []T {
	if values, ok := i.secondaryIndex[index]; ok {
		return values
	}

	return []T{}
}

func (i *Index[T, P, S]) Keys() []P {
	keys := make([]P, 0, len(i.primaryIndex))

	for key := range i.primaryIndex {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func (i *Index[T, P, S]) SecondaryKeys() []S {
	keys := make([]S, 0, len(i.secondaryIndex))

	for key := range i.secondaryIndex {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}
