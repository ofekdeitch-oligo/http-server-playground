package utils

func Find[T any](items []T, predicate func(T) bool) (bool, T) {
	for _, item := range items {
		if predicate(item) {
			return true, item
		}
	}

	return false, NoValue[T]()
}
