package main

func noValue[V any]() V {
	var result V
	return result
}
