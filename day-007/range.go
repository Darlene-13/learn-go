package main

import "fmt"

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

// Range Syntax makes it easier to iterate over elements like slice
// for ELEMENT, INDEX := range SLICE
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badword := range badWords {
			if word == badword {
				return i
			}
		}
	}

	return -1
}
