package main

import "fmt"

// Arrays in golang are a fixed length sequence that is used to store homogenous elements in memory.
// Due to their fixed length behaviors Slices are much more preferred.
// Once is allowed to store zero or more than zero elements
// The elements are indexed using the []index elements

// CREATING AND ACCESSING AN ARRAY
// Arrays are created using the var keyword of a particular type with name and size
var names [10]string

//Arrays are mutable
//var firstNames[0] = "Gopher"
// Store the name Gopher as index 0 or the first element of the firstName array

// ACCESSING THE ELEMENTS
//  can be accessed using the index value or using the for loop
// Array type is 1d and unchangeable and allows duplicate

// DECLARATION USING SHORT HAND DECLARATION
// Arrays can be declared using the shorthand variable declaration
//secondNames := [2]string{
//	"Gopher","Darlene",
//}

func main() {
	ourNames := [4]string{
		"Darlene",
		"Wendy",
		"Stacy",
		"Juma",
	}

	fmt.Println("The following are our names: ")

	for i := 0; i < len(ourNames); i++ {
		fmt.Println(ourNames[i])
	}

	main1()
	main3()
	main4()
	main5()
}

// MULTI DIMENSIONAL ARRAY
// MD arrays are arrays of the arrays of the same type
// They can be created using the var or the shorthand declaration.

func main1() {

	// Using the shorthand declaration
	studentNames := [3][3]string{
		{"Darlene", "Wendy", "Nasimiyu"},
		{"Stacy", "Phanice", "Juma"},
	}

	fmt.Println("The student names: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(studentNames[i][j])
		}
	}

	// Using the var keyword
	var students [3][3]int
	students[0][0] = 100
	students[0][1] = 200
	students[0][2] = 300
	students[1][0] = 400
	students[1][1] = 500
	students[1][2] = 600

	fmt.Println("Student marks are: ")
	for p := 0; p < len(students)-1; p++ {
		for q := 0; q < len(students[p]); q++ {
			fmt.Println(students[p][q])
		}
	}
}

// If an arrays does not initialize explicitly then the default value of the array is 0.
// You can find the length using len
//If ellipsis becomes visible at the place of length then the length of the array is determined by the initialized elements.

func main3() {
	myArray := [...]string{
		"A", "B", "C", "D", "E", "F", "G",
	}

	fmt.Println(myArray)
	fmt.Printf("Length of array: %d\n", len(myArray))

}

// In an array you are allowed to iterate over the range of the elements of the array as shown below
func main4() {

	myNumbers := [3]int{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Println(myNumbers[i])
	}
}

//An array is of value type and not a reference type
// So when an array is assigned to a new var, changes made in the new var are for the new var and do not affect the original array

func main5() {
	myNumbers := [...]int{1, 2, 3}
	fmt.Printf("Length of array: %d\n", len(myNumbers))
	fmt.Println("Original array before: ", myNumbers)

	//Assigning the original array myNumbers to a new one my floats
	myFloats := myNumbers
	fmt.Println("New array before: ", myFloats)

	// Change the values at the first and second indexes
	myFloats[0] = 100
	myFloats[1] = 200

	fmt.Println("The new array values after : ", myFloats)

	fmt.Println("Original array after: ", myNumbers)

}

func main6() {

	arr0 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}

	fmt.Println(arr0, arr2, arr3)
	fmt.Println(arr0 == arr2, arr3 == arr3)
}

// Two arrays can be compared using the == operator
