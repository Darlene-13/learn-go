package main

import "fmt"

// Decision-making in go is similar to what happens in real life.
// This  basically involves a piece of code that gets executed  when a given condition is fulfilled, also termed as control flow statements.
// They control the flow of execution of a program on certain conditions.
// They are used to cause the flow of execution to advance and branch based on the changes to the state of the program

// Decision-making statements are the "if statements"
// It is the simplest decision-making statement as it is used to determine whether a certain statement or block of statements will be executted or not that is if a certain condition is true otherwise not.

func main() {
	var v int = 700

	// Has a time complexity and space complexity of O(1)
	if v < 1000 {
		fmt.Printf("v is less than 1000\n")
	}

	fmt.Printf("Value of v is:  %d\n", v)

	// if else statements
	mainZ()
	// Nested if else statements
	mainW()

	// Multiple if else if
	mainT()
}

//If else statement
// The else statement can be used to execute something else if the "if" is false

func mainZ() {
	var v int = 1200
	if v < 1000 {
		fmt.Printf("v is less than 1000\n")

	} else {
		fmt.Printf("v is greater than 1000\n")
	}
}

// NESTED IF
// In go language nested if is an if statement that is the target of another if or else statement
// ie: if condition1 { executes when condition1 is true ... if condition2 }

func mainW() {
	var v1 int = 400
	var v2 int = 700

	if v1 == 400 {
		fmt.Printf("v1 is equal to 400\n")
		if v2 == 700 {
			fmt.Printf("Value of v1 is 400 and v2 is 700\n")
			fmt.Printf("v2 is equal to 700\n")

		}
	}
}

// If else if ladder
// A user can decide amongst multiple conditions.
// The if statements are executed top to bottom as sson as one of the conditions controlling the if is true, the statement associated with that if is executed and the rest of the ladder is bypassed, otherwise as long as none of the if is true then the final else will be executed.

func mainT() {
	var v1 int = 600

	if v1 == 100 {
		fmt.Printf("Value of v1 is 100\n")
	} else if v1 == 400 {
		fmt.Printf("Value of v1 is 400\n")
	} else if v1 == 800 {
		fmt.Printf("Value of v1 is 800\n")
	} else {
		fmt.Print("None of the values is matching\n")
	}

}
