package main

import (
	"fmt"
)

const MessageToDisplay = "Hello world"

func main() {
	fmt.Printf("Welcome!, %s\n", MessageToDisplay)

	a, b, c := numbers()

	fmt.Printf("a: %d b: %d, c: %d\n", a, b, c)
}

func numbers() (a, b, c int) {
	return 1, 2, 3
}
