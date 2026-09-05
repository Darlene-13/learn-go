package main

import "fmt"

// TYPE ASSERTION
func getExpenseReport(e expense1) (string, float64) {

	em, ok := e.(email1)
	if ok {
		return em.toAddress, em.cost()
	}

	sm, ok := e.(sms)
	if ok {
		return sm.toPhoneNumber, sm.cost()
	}

	return "", 0.0
}

// STRUCTS

type email1 struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct {
}

// INTERFACES

type expense1 interface {
	cost() float64
}

type printer1 interface {
	print()
}

// METHODS

func (e email1) cost() float64 {

	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}

	return 0.01 * float64(len(e.body))
}

func (e email1) print() {
	fmt.Println(e.body)
}

func (sm sms) cost() float64 {

	if !sm.isSubscribed {
		return 0.05 * float64(len(sm.body))
	}

	return 0.01 * float64(len(sm.body))
}

func (sm sms) print() {
	fmt.Println(sm.body)
}

func (i invalid) cost() float64 {
	return 0.0
}

// HELPER FUNCTIONS

func getEmail1(e expense1, p printer1) {

	fmt.Println("=============================================")

	fmt.Printf(
		"Printing with cost %.2f ........\n",
		e.cost(),
	)

	p.print()
}

func getSms(sm expense1, s printer1) {

	fmt.Println("=============================================")

	fmt.Printf(
		"Printing with cost %.2f ........\n",
		sm.cost(),
	)

	s.print()
}

// TYPE SWITCH

func getEmailOrSms(e expense1) {

	address, cost := getExpenseReport(e)

	switch e.(type) {

	case email1:

		fmt.Printf(
			"Report: The email is going to %s and it will cost: %.2f\n",
			address,
			cost,
		)

		fmt.Println("=================================================================")

	case sms:

		fmt.Printf(
			"Report: The sms is going to %s and it will cost: %.2f\n",
			address,
			cost,
		)

		fmt.Println("=================================================================")

	default:

		fmt.Println("Report: Invalid expense")

		fmt.Println("=================================================================")
	}
}

// MAIN

func main() {

	// EMAIL 1

	e1 := email1{
		isSubscribed: true,
		body:         "Hello there!",
	}

	getEmail1(e1, e1)

	// EMAIL 2

	e1 = email1{
		isSubscribed: false,
		body:         "I want my money back",
	}

	getEmail1(e1, e1)

	// EMAIL 3

	e1 = email1{
		isSubscribed: true,
		body:         "I will send it to you tomorrow",
	}

	getEmail1(e1, e1)

	// SMS 1

	sm := sms{
		isSubscribed:  true,
		toPhoneNumber: "070000000",
		body:          "Where are you?",
	}

	getSms(sm, sm)

	// SMS 2

	sm = sms{
		isSubscribed:  true,
		toPhoneNumber: "070000000",
		body:          "I am at the Club",
	}

	getSms(sm, sm)

	// SMS 3
	sm = sms{
		isSubscribed:  false,
		toPhoneNumber: "070000000",
		body:          "Okay please come how now?",
	}

	getSms(sm, sm)

	// EMAIL REPORT 1
	getEmailOrSms(email1{
		isSubscribed: true,
		body:         "Hello baby!",
		toAddress:    "janedoe@gmail.com",
	})

	// EMAIL REPORT 2

	getEmailOrSms(email1{
		isSubscribed: false,
		body:         "Hello there, my friend!",
		toAddress:    "johndoe@gmail.com",
	})

	// SMS REPORT 1
	getEmailOrSms(sms{
		isSubscribed:  true,
		body:          "What would you want to have for dinner?",
		toPhoneNumber: "0790909990",
	})

	// SMS REPORT 2
	getEmailOrSms(sms{
		isSubscribed:  false,
		body:          "I would love to have some chicken soup and cassava?",
		toPhoneNumber: "0790909990",
	})

	// INVALID EXPENSE
	getEmailOrSms(invalid{})
}
