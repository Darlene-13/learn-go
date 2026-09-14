package main

import "fmt"

// Spread operator is the opposite of the variadic operator
// It allows us to pass a slice to a variadic function

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)
}
