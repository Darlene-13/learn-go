package main

import "fmt"

func main() {

	var username string = "Darlene"
	var password string = "12345678"
	fmt.Println("Authentication Basic", username+": "+password)
}

// Two strings can be concatenated using + sign.
// Go uses a garbage collector like java therefore it has an automated memory management.
// Difference between java and go is that go does not use a virtual machine, whilst it is clear that using a virtual machine creates a memory overhead.
// But go uses a runtime within every single binary in which it is used.

// Go runtime is basically used to clean up unused memory.
