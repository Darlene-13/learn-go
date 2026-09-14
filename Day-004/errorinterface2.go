package main

import "fmt"

// user error struct with the user name that implements the error interface
// With this we could store structured data.
type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("User Error: %v", e.name)
}

// A function called cansendtouser which takes in one argument called username that is a string and returns a boolean
// WHAT I LEARNT: It does not need a whole user struct to get to know it just needs the username .......lol
// If it were canSendToUser(user)  we would be passing the whole user struct
func canSendToUser(userName string) bool {
	// Give me a username of type string
	allowedUsers := map[string]bool{
		"Darlene": true,
		"Stacy":   true,
		"John":    true,
	}
	return allowedUsers[userName]
}

// Define a function called sends that accepts two parameters msg and username both of which are strings and returns an error
func sendSMS(msg, userName string) error {
	if !canSendToUser(userName) {
		return userError{name: userName}
	}

	fmt.Printf("SMS sent to %s: %s\n", userName, msg)

	return nil
}

func main() {
	err := sendSMS("Hello World", "John")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("SMS Sent successfully")

	fmt.Println("============================================")

	err = sendSMS(
		"Hey, your package has arrived",
		"John",
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("SMS successfully sent!")

}
