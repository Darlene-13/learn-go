package main

import "fmt"

// Go provides a single looping construct, the for loop which can be used in multiple ways to implement iteration
// It can behave like a traditional loop, a while loop, infinite loop and can also iterate over collection like arrays, slices, strings, maps, and channels using the range.

func main() {

	// Traditional for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		fmt.Printf("My name is Darlene Wendy! ")
	}

	traditionalLoop()
	infiniteLoop()
	forLoopAsWhileLoop()
	forRangeLoop()
	iterateOverString()
	iterateOverMaps()
	iterateOverChannels()
}

// Types of loops in go
// TRADITIONAL FOR LOOP IN GO
func traditionalLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// Initialize i with zero continue looping as long as the current number i is still less than 10
// Time complexity is O(n) while space is O(1)

// /INFINITE FOR LOOP
func infiniteLoop() {
	count := 1
	for {
		fmt.Println("Count: ", count)
		count++
		if count > 5 {
			break
		}
	}
}

/**
Create an infinite loop without any expressions
Print the value of count within each iteration
Increment count after every iteration
Use break statement to terminate the loop when the count  becomes greater than 5
*/

// FOR LOOP AS A WHILE LOOP
func forLoopAsWhileLoop() {
	sum := 0
	num := 1

	for num <= 5 {
		sum += num
		num++
	}

	fmt.Println("Sum: ", sum)
}

/**
Initialize sum with 0 and num with 1
use a for loop (while) with only a condition
*/

// FOR RANGE LOOP
func forRangeLoop() {
	numbers := []int{1, 2, 3, 4, 5}

	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %v\n", i, v)
	}
}

/**
Create a slice named numbers
Use a for range to iterate over the slice
Index stores the position of each element
Values stores the element of the corresponding numbers in the same position
Print both the index and the value during each iteration
*/

// ITERATE OVER STRING USING A FOR RANGE
func iterateOverString() {
	str := "Darlene"

	for i, c := range str {
		fmt.Printf("Index: %d, Value: %v\n", i, string(c))
	}
}

/**
Create a string str named : Darlene
use a for range to iterate over the string
index stores the starting byte index of each character and char stores the corresponding Unicode character
*/

// ITERATE OVER MAPS USING A FOR RANGE
func iterateOverMaps() {
	students := map[int]string{
		1: "Darlene",
		2: "Wendy",
		3: "Nasimiyu",
	}

	for k, v := range students {
		fmt.Printf("ID: %d, Name: %s\n", k, v)
	}
}

/**
Create a map named students that takes an integer as the key and string as values
use a for range to iterate over the map
Key stores the current map key
values store the corresponding value

*/

// ITERATE OVER CHANNELS USING A FOR RANGE
func iterateOverChannels() {
	ch := make(chan int, 3)
	go func() {
		ch <- 1 // This means inserting 1 to the channel
		ch <- 2
		ch <- 3

		close(ch)
	}()

	for value := range ch {
		fmt.Printf("Value: %d\n", value)
	}
}

/**
Create an integer channel named ch that has a capacity of 3
Use a go routine to send values to the channel
Close the channel after sending the values to the channel
Use a for range to receive and print each value
// The loop automatically terminates when the channel closes.
*/
