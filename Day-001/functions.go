package main

import "fmt"

// It takes zero inputs
func main() {
	fmt.Println("Darlene,", "Happy Birthday!")
	fmt.Println("Stacy,", "Happy Valentines Baby!")
	fmt.Println("Go", "is actually fantastic")

	result := sub(10, 8)

	//fmt.Printf("Subtracting %v from %v is equal to: %v\n", 10, 8, sub(10, 8)
	fmt.Printf("Subtracting %v from %v is equal to: %v\n", 10, 8, result)
}

func concat(s1, s2 string) string {
	return s1 + s2
}

func sub(num1 int, num2 int) int {
	return num1 - num2
}

// In go variables are passed by values and not by reference.
