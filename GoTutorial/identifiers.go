package main

import "fmt"

// What is an identifier?
// Identifiers are names given to program elements such as variables, functions, constant packages and types
// They help identify and access different parts of a Go program

/*
TYPES OF IDENTIFIERS
1. Variables names: age, name, totalMarks
2. Function names: main, display, calculateSum
3. Constant names: PI, MAX_SIZE, sitName
4. Package names: main, fmt, math
5. Type names: Student, Employee, Person

*/

const siteName = "Anyango's site"

type Student struct {
	Name string
}

func display() {
	fmt.Println("Hello World")
}

func classAttendance(student Student) {
	fmt.Printf("Student Name: %s\n", student.Name)
}

func main() {
	var age = 20
	display()
	fmt.Println("Age", age)

	main3()
}

// In the above code: Age is a variable identifier, display is a function identifier, siteName is a constant identifier,
// main and fmt are package identifiers, and student is a type identifier.

// A valid go identifier must start with letter A - Z or an underscore
// It can contain, letters, digits and underscores
// Can't start with a digit, and they are case-sensitive and go keywords can't be used as an identifier

// BLANK IDENTIFIER
// The _underscore identifier also known as the blank identifier is used when a value is required syntactically but doesn't need to be stored.

func main3() {

	// 10 is assigned to _ and then ignored.
	_, b := 10, 20
	fmt.Println(b)
}

// EXPORTED IDENTIFIERS
// They begin with an uppercase letter and are called exported identifiers, they can be accessed from other packages



var TotalMarks = 500
// Since TotalMarks starts with a capital letter to mean it can be exported

//NON EXPORTED IDENTIFIER
var totalMarks = 500
// Since totalMarks begins with a small letter it can not be exported and used with other packages.
