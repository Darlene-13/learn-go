package main

import "fmt"

// Keywords are reserved words with pre-defined meanings
// They are used to define the structure and behavior of a program such as declaring variables, creating functions, controlling program flow and importing packages
// They therefore can't be used as identifiers for variables, functions, constants or other program elements

func main() {
	var name = "Darlene"
	fmt.Printf("Hello, %s!\n", name)
	const language = "Go"
	for i := 1; i <= 3; i++ {
		if i == 2 {
			fmt.Printf("Hello, %s!\n", language)
		}
	}
	// Var is the keyword while age is the identifier.
	var age = 25
	fmt.Println(age)
}

// Difference between keyword and Identifiers
// Keywords are reserved by a language while identifiers are set by the programmer
// Keywords can't be used as names, identifiers are used to name program elements

// Package defines package names, import, includes external packages to the program, func is used to define a function and var used to create a variable.

// List of go keywords.
// break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto, if, import
// interface, map, package, range, return, select, struct, switch, type, var

// Defer keyword in go: Delays the execution of a function until the sounding function finishes everything.
