package main

import "fmt"

// Array: fixed length seq that holds elements of a specific type, unlike slices they have a constant type determined when the array is declared.
// Copying one array to another is required but requires them to be of the same length at least and type.

func main() {
	//Basic array
	var source = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Source array: ", source)

	var source3 [4]int
	fmt.Println(source3)

	// Assignment by looping
	mainA()

	mainB()

}

// Copying an array can happen through the following ways:
//1. Using a for loop 2. Direct copying of the array 3. Using pointers if arrays are large

// FOR LOOP COPYING
//1. Using a for loop: There is not built in copy function, so we loop over the array and do the assignment one by one

var mySource = [5]int{1, 2, 3, 4, 5}

func mainA() {

	var myDestination [5]int
	// Manually copying the array
	for i := 0; i < len(mySource); i++ {
		myDestination[i] = mySource[i]
	}
	fmt.Println("My source array: ", mySource)
	fmt.Println("My destination array: ", myDestination)

}

// DIRECT ASSIGNMENT (This only works with arrays and not slices"
func mainB() {
	var myDestination [5]int = mySource

	fmt.Println("Copying by direct assignment which only works for arrays and not slices")
	fmt.Println("My source array: ", mySource)
	fmt.Println("My destination array: ", myDestination)
}

// USING POINTERS FOR CASES OF LARGE ARRAYS
func mainC() {

	var myDestination *[5]int = &mySource

	fmt.Println("Copying using pointers and references and this happens for large arrays")
	fmt.Println("My source array: ", mySource)
	fmt.Println("My destination array: ", myDestination)
}
