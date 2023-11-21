package main

import "fmt"

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hello Annonymous"
	}

	return fmt.Sprintf("Hello %s", name)
}
