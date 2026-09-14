package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		total = +num

	}
	return total
}

func main() {
	total := sum(1, 2, 3) // Go gives the function a slice of size 2 with 3 integers
	fmt.Println(total)
}
