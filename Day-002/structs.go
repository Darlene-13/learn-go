package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
	recipient   user
}

type user struct {
	username string
	email    string
}

func test(m messageToSend) {
	fmt.Printf(
		"Sending message: '%s' to: %v. Your username is %s and your email is %s\n",
		m.message,
		m.phoneNumber,
		m.recipient.username,
		m.recipient.email,
	)

	fmt.Println("=============================================================================================================================================")
}

func canSendMessage(mToSend messageToSend) bool {
	// Guard clauses
	if mToSend.recipient.username == "" {
		return false
	}

	if mToSend.recipient.email == "" {
		return false
	}

	if mToSend.phoneNumber == 0 {
		return false
	}

	return true
}

func main() {

	// First message
	test(messageToSend{
		phoneNumber: 12345678,
		message:     "Thanks for signing up",
		recipient: user{
			username: "Darlene Wendy",
			email:    "darlenewendie@gmail.com",
		},
	})

	// Second message
	test(messageToSend{
		phoneNumber: 87654321,
		message:     "Welcome to Atee Cakes and Bakes",
		recipient: user{
			username: "Stacy Phanice",
			email:    "stacyjuma000@gmail.com",
		},
	})

	// Third message - missing phone number
	test(messageToSend{
		message: "Thanks for signing up",
		recipient: user{
			username: "Darlene Wendy",
			email:    "darlenewendie@gmail.com",
		},
	})

	// Testing the canSendMessage function
	cansend := canSendMessage(messageToSend{
		phoneNumber: 12345678,
		message:     "Thanks for signing up",
		recipient: user{
			email: "darlenewendie@gmail.com",
		},
	})

	canSend := canSendMessage(messageToSend{
		phoneNumber: 12345678,
		message:     "Hello",
		recipient: user{
			username: "Darlene",
			email:    "darlene@example.com",
		},
	})

	fmt.Println("Can send:", canSend)
	fmt.Println("Can send:", cansend)

	car1 := car{
		Make:   "Toyota",
		Model:  "Corolla",
		Width:  10,
		Height: 12,
	}

	// Another struct

	car2 := car{
		Make:   "Toyota",
		Model:  "Corolla",
		Width:  10,
		Height: 12,
	}

	fmt.Println("The first car is: ", car1.Make)
	fmt.Println("The second car is: ", car2.Make)

	myCar := struct {
		Make  string
		Color string
	}{
		Make:  "Corolla",
		Color: "White",
	}

	fmt.Println("This type of car uses anonymous struct that can only be instantiated once ", myCar)
}

// NOTES ON STRUCTS

// A struct is a custom data type that groups related values together.
// Structs allow us to represent something that has multiple related properties.

// Positional struct literals are possible, but using named fields is preferred.

// We can create an empty struct value and assign values to its fields later.

// We can have structs inside other structs.

// Structs can be passed as function arguments.

// Structs can also be returned from functions.

type car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel wheel
	BackWheel  wheel
}

type wheel struct {
	Radius   int
	Material string
}

// Creating a struct using named fields

// Accessing struct fields
// To access struct fields, we use the . operator.

// Example:
// car1.Make
// car1.Model
// car1.Height

// Struct literals
// There are several ways to create structs.
// The most preferred way is using named fields.

// Anonymous struct
// An anonymous struct is useful when we only need the struct once
// and don't need to define a reusable named type.

// EMBEDDED STRUCTS
type mCar struct {
	model string
	color string
}

type truck struct {
	car
	bedsize int
}

// The major difference between an embedded and nested struct is how we access the variables in the struct that has been embedded.
