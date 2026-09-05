package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type recipientName struct {
	firstName string
	lastName  string
}
type birthdayMessage struct {
	birthdayTime time.Time
	recipientName
}

func (br birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s %s, it is your birthday on %s", br.firstName, br.lastName, br.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s report is ready. So far you have sent %v number of reports", sr.reportName, sr.numberOfSends)
}

// Test function
func test(msg message) {
	sendMessage(msg)
	fmt.Println("==================================")
}

// Main function
func main() {
	test(sendingReport{
		reportName:    "Darlene's birthday report",
		numberOfSends: 25,
	})

	test(birthdayMessage{
		birthdayTime: time.Date(2003, 01, 21, 0, 0, 0, 0, time.UTC),
		recipientName: recipientName{
			firstName: "Darlene",
			lastName:  "Wendy",
		},
	})

	test(sendingReport{
		reportName:    "Wendy's birthday report",
		numberOfSends: 23,
	})

	test(birthdayMessage{
		birthdayTime: time.Date(2003, 01, 21, 0, 0, 0, 0, time.UTC),
		recipientName: recipientName{
			firstName: "Wendy",
			lastName:  "Darlene",
		},
	})
}
