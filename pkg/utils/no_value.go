package utils

func NoValue[V any]() V {
	var result V
	return result
}
