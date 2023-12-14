package tools

func ArrayRemove[T comparable](index []T, item T) []T {
	for idx, val := range index {
		if val != item {
			continue
		}

		return append(index[:idx], index[idx+1:]...)
	}

	return index
}
