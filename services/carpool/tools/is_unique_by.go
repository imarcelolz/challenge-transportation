package tools

func IsUniqueBy[T any, I comparable](collection []T, key func(T) I) bool {
	existing := make(map[I]bool)

	for _, item := range collection {
		keyValue := key(item)

		if existing[keyValue] {
			return false
		}

		existing[keyValue] = true
	}

	return true
}
