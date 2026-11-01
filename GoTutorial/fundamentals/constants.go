package main

import "fmt"

// Constants are fixed ie: they can't be modified once declared, there can be any basic data type of type constant like integer, float, string.
// They are declared as variables only that we use the const keyword

const secondName = "Wendy"

func main() {
	const firstName = "Darlene"
	fmt.Printf("Hello, %s!\n", firstName)

	fmt.Printf("Hello, %s!\n", secondName)

	//String constants
	const myName = "Nasimiyu"
	var herName = "Stacy"
	//concat strings
	var ourNames = myName + " " + herName
	fmt.Println(ourNames)

	//compare the strings
	fmt.Println(myName == herName)
	fmt.Println(myName < herName)

	//Boolean constants
	//Similar to string constants, but it only has two untyped: true and false.
	const trueConstant = true
	// Type definition using type keyword
	type myBool bool
	var defaultBool = trueConstant
	var customBool myBool = trueConstant
	// defaultBool = customBool is not allowed
	fmt.Println(customBool)
	fmt.Println(defaultBool)

}

// TYPED AND UNTYPED NUMERIC CONSTANTS
// Untyped and types numeric constants
// Typed constants work like immutable variables with the same type and types constants work like literals can interoperate with similar types.

//Types of Constants in go.
// Numeric constants, String literals and Boolean constants.
// Numeric could be : integer, floats and complex
// Integer constants: can have a suffix that is a combination of U(uppercase) and L lowercase for unsigned and long respectively
// It could be decimal, octal or hexadecimal constant and can store a max or 64-bit integer
// floating types have an integer and decimal parts, fractional or exponential part
// While represented using the decimal and exponential form they must include (decimal, exponential or both) and (integer, fractional or both) respectively.
// String literals: "" and ''
//String literals are comparable using operators like: ==, !=
