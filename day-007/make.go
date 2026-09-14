package main

import (
	"fmt"
)

func main() {
	//mySlice := make([]int, 5, 10)
	mySlice := make([]int, 5)
	_ = mySlice
}

//We can use make to create a new slice instead of looking at the underlying array
// 10 = Capacity
// Length len() returns the length of the slice while Capacity cap returns capacity
// Resizing a slice can be computationally expensive
// If we can know the slice size ahead of time so that go routines do not allocate the size over and over again

// Create a function called getMessageCost that takes a slice of messages which is a string and returns a float64 which is the cost
func getMessageCost(messages []string) []float64 {
	// Empty slice that holds the costs and the slice is of the length(messages)
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}

	return costs
}

func text(messages []string) {
	costs := getMessageCost(messages)
	fmt.Println("Messages")
	for i := 0; i < len(messages); i++ {
		fmt.Printf("- %v\n", messages[i])
	}

	fmt.Println("Costs")
	for i := 0; i < len(costs); i++ {

	}
}
