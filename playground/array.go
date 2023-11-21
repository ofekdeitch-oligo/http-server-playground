package main

func sum(arr []int) int {
	return sumBy(arr, func(value int) int {
		return value
	})
}

func sumBy[T any](arr []T, getValue func(T) int) int {
	result := 0

	for i := 0; i < len(arr); i++ {
		result += getValue(arr[i])
	}

	return result
}

func max(arr []int) int {
	return maxBy(arr, func(value int) int { return value })
}

func maxBy[T any](arr []T, getValue func(T) int) int {
	result := getValue(arr[0])

	for i := 0; i < len(arr); i++ {
		current := getValue(arr[i])

		if result < current {
			result = current
		}
	}

	return result
}

func Map[T any, V any](arr []T, getValue func(T) V) []V {
	mapped := make([]V, len(arr))

	for i := 0; i < len(arr); i++ {
		mapped[i] = getValue(arr[i])
	}

	return mapped
}
