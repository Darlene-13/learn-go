package main

/**
 Golang is an open source programming language developed by Google and released in 2009
Widely used for developing: Web application, cloud services, networking tools, api system software.
It has built in concurrency support using Go routines and Channels for running multiple tasks efficiently.
*/
import (
	"fmt"
	"time"
)

// Take for example: One waiter receives an order, another waiter serves a different table and the other one prepares for delivery
func main() {
	fmt.Println("Hello World")

	var lats = [2]float64{-1.098899, 1.989044}
	fmt.Println(lats)

	// Take different orders
	// The go keyword starts a goroutine
	// These three go routines can be processed concurrently: Separately yes but they tend to appear as if they are happening at the same time.

	// Real World Applications of this:
	// 1. Processing multiple API requests
	// 2. Sending emails in the background
	// 3. Handling multiple MQTT messages
	// 4. Running background jobs.
	go processOrder(1)
	go processOrder(2)
	go processOrder(3)

	time.Sleep(3 * time.Second)

	// Create a channel
	ch := make(chan string)
	// Start a go routine
	// We start a worker that waits for a message
	go worker(ch)
	// send a message process payment through the channel
	ch <- "Process payment"

	// ch <- "My name is Darlene" : this is used to send to a channel
	// value := <-ch Receive from a channel
}

// Example of a go routine
func processOrder(orderId int) {
	fmt.Printf("Processing order %d\n", orderId)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished processing order %d\n", orderId)
}

// Since we had the three go routines
// The three go routines communicate through or using a channel.
// The channel is a communication mechanism that allows go routines to send and receive values.
// Think of it as a conveyor belt, a go routine puts the work in the belt and the other receives it.

// I am defining a function called a worker that takes one param called ch and that parameter is a channel or strings
func worker(ch chan string) {
	message := <-ch
	fmt.Print("Worker received a message", message)
}
