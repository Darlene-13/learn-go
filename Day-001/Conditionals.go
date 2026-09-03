package main

import "fmt"

func main() {

	conditional()
}

func conditional() {

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length: ", messageLen, "and the maximum length is", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
