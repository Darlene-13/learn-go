package main

import "fmt"

// Multiple return values
func getNumbers() (int, int) {
	return 1, 2
}

// / Variables are names that hold data which can change when the program is running
func main() {
	var number int = 10
	fmt.Println("Number of jobs: ", number)

	// VARIABLE DECLARATION WITH TYPE INFERENCE
	var myVariable int = 23
	var myVariable1 string = "Darlene"
	var myVariable2 float64 = 23.45

	fmt.Println(myVariable, myVariable1, myVariable2)

	//ZERO VALUE VARIABLES
	var myVariable3 int
	var myVariable4 string
	var myVariable5 float64
	var myVariable6 bool

	fmt.Printf("My variable %d\n", myVariable3)
	fmt.Printf("My variable %d\n", myVariable4)

	fmt.Printf("My variable %d\n", myVariable5)
	fmt.Printf("My variable %d\n", myVariable6)

	var a, b, c = 2, "Engineer", 67.56
	fmt.Println(a, b, c)

	e, f := getNumbers()
	fmt.Println(e)
	fmt.Println(f)

	w, d := getValues()
	fmt.Println(w)
	fmt.Println(d)
}

/*
RULES OF NAMING VARIABLES
1. Variable names must start with a letter or underscore
2. Names can contain a letter or digits or underscores
3. They can't start with a digit, and they are case-sensitive
4. Keywords can't be used as variable names, and they have no strict limit on the length but it is recommended to have 4-15 chars
5.
*/

// Variables can be declared using the var keyword
// Type or expression can be omitted and the variable will be assigned zero, false for booleans and "" for strings and nil for pointers or reference types

// Declaring multiple variables of different types
//Multiple variables of different types can be declared in a single line
// Short variables: Local variables that are declared and initialized in functions using short variable declaration
//Multiple variables of the same type can be declared using :=
// ie: myvar1, myvar2, myvar3 := 3,4,5

// Multiple variables of different types can be declared using :=
// ie: myvar1, myvar2, myvar3 := 3, "Darlene", 8.992
// Multiple return values using short declaration
func getValues() (int, string) {
	return 7, "Golang"
}
