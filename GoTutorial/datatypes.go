package main

import "fmt"

//Data types specify the type of data a valid go variable can hold

// FOUR CATEGORIES OF DATA TYPES
//1. Basic: Numbers, strings, booleans
//2. Aggregate types: Array and structures
//3. Reference types:Pointers, slices, map, functions, channels
//4. Interface
// Here we will discuss the basic types.

// Integers in Go language are both signed and unsigned are available and in different sizes signed represented by int and the other by uint
// Both int and uint contain same size, either 32 or 64 bit.

// rune is a synonym of int32 and also represent Unicode code points
// uintptr: unsigned integer type, with undefined width but can hold all the bits of a pointer value
// Signed numbers can store both positive and negative while unsigned they can only store positive or non negative values

func main() {
	var X uint8 = 225
	fmt.Println(X, X-3)

	var Y int16 = 32687
	fmt.Println(Y, Y-3)

	// Examples of arithmetic operations
	var Z int16 = 150
	var W int16 = 250
	fmt.Printf("Adition of %d + %d = %d\n", Z, W)
	fmt.Printf("Substraction of %d - %d = %d\n", W, Z)
	fmt.Printf("Division of %d / %d = %d\n", W, Z)
	fmt.Printf("Multiplication of %d / %d = %d\n", W, Z)
	fmt.Printf("Modulous of %d / %d = %d\n", W, Z)

	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)

	fmt.Printf("The type of a %T and "+"the type of  is %T\n", a, b)

	// Boolean
	str1 := "Darlene"
	str2 := "Wendy"

	result1 := str1 == str2
	result2 := str1 != str2

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of result 1 is %T\n", result1)
	fmt.Printf("The type of result 2 is %T\n", result2)

	//Strings
	str := "I am a software Engineer"

	fmt.Println(str)

	fmt.Printf("The type of string  is %T\n", str)
	fmt.Printf("The length of string is %d\n", len(str))

	// String concatenation
	var stra = "Darlene"
	var strb = "Wendy"

	fmt.Println("New string: ", stra+strb)
}

// floating point numbers can be divided to float32 and float64
// complex numbers can also be divided to complex64 and complex 128
// complex 64 are complex numbers that contain float32 as the real and imaginary component
// complex 128 are complex numbers that contain float64 as the real and imaginary number.
// boolean: They rep only one bit of information: either true or false.
// They can't be implicit or explicitly be converted to any other type of code
// Strings: The string daa type represents a string of Unicode code points
// In other words we could say that a string is a sequence of immutable bytes, once created can't be changes, but they can be concataned using the + operator

//When to use (var) and :=
// Use := when declaring and assigning inside a function
// use var: when declaring without assigning a value
// Use var also when explicitly specifying the type, declaring  a package level variable or one that will be assigned later