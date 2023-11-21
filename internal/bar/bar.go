package bar

import "internal/foo"

func Bar() {
	println("Bar")

	foo.Foo()
}
