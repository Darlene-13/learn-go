package main

import "fmt"

func main() {

	messageFromDoris := []string{
		"You doing anything later ?",
		"Did you get my last message ?",
		"Don't leave me hanging.",
		"Please respond I am lonely!",
	}

	numMessages := float64(len(messageFromDoris))

	costPerMessage := .02

	totalCost := numMessages * costPerMessage

	fmt.Printf("Doris spent $%.2f on text messages today\n", totalCost)
}

// Go executes slower than rust.
// Go is a compiled language.
// Compilation is taking human-readable code and changing it to machine code or binary.
// main.go -> go.build -> main.exe
//Compiling is faster at runtime because it is done upfront.
// Compile time: When we are actually changing our source code file .go to something executable
// Run time: When we are actually running or executing our code
