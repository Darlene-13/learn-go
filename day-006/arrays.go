package main

import "errors"

// An array in go
var arr []int

func getMessageWithRetries() [3]string {
	return [3]string{
		"Click here to sign up",
		"Pretty please click here",
		"We beg you to sign up"}
}

func main() {

}

// In go arrays have fixed sizes, slices do not have fixed sizes.
//Slices in go: Dynamically size flexible arrays, they are built on top of the array (memory management)
//[] int

// primes := [6]int{2, 3, 5, 7, 11, 13}
// mySlice := primes[1:4]
const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[2:], nil
	}
	if plan == planFree {
		return allMessages[2:], nil
	}
	return nil, errors.New("Unsupported plan")
}

// Ram is a mapping of address to data
//Slices and arrays they are made of are stored in contiguous memory
//Slice is an address in memory where the array starts.
//Arrays have fied size to avoid overwriting data in memory.
// mySlice := make([] int, 5, 10)
// 10: Capacity, length of the underlying array
// 5, length
//We have a built in length and capacity (cap) for slices
