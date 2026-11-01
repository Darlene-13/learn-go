package main

import "fmt"

func calculateAverage(arr [6]int, size int) int {
	var sum int
	for _, value := range arr {
		sum += value
	}

	return sum / size
}

func main() {

	scores := [6]int{1, 2, 3, 4, 5, 6}
	calculateAverage(scores, len(scores))
	average := calculateAverage(scores, len(scores))
	fmt.Printf("The average is: %d\n", average)

	modifyArray(&scores)
	fmt.Printf("Modify array is: %d\n", scores)
	incrementArray(&scores)
	fmt.Printf("Incremented array is: %d\n", scores)
}

/*
In this example the calculate Average function takes an array of integers and its sizes as arguments, computes the average of array elements and returns it.
The main function initializes the array called scores and calls the average calculation function

*/

/**
KEY POINTS
1. Array Declaration: Use the syntax var arrayName [size] type
2. In go arrays are passed to functions by values meaning the function receives a copy of the array, changes made inside the function do not affect the original array.
3. If we want to modify the original array we should pass a pointer to it.

*/

func modifyArray(arr *[6]int) {
	for i := range arr {
		arr[i] += 10 // increment each element by 10
	}
}

func incrementArray(arr *[6]int) {
	for i := range arr {
		arr[i] += 1
	}
}
