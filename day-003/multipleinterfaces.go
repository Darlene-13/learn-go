package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
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
type email struct {
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

// INTERFACES TO BE IMPLEMENTED
type expense interface {
	cost() float64
}

type printer interface {
	print()
}

// FUNCTIONS
func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body)) //Go is strictly type so we need to cast the int to a float first
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
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

// HELPER FUNCTION
func getEmail(e expense, p printer) {
	fmt.Println("=============================================")
	fmt.Printf("Printing with cost %2f ........\n", e.cost())
	p.print()
}

func getSms(sm expense, s printer) {
	fmt.Println("=============================================")
	fmt.Printf("Printing with cost %2f ........\n", sm.cost())
	s.print()
}

func getEmailOrSms(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email is going to %s and it will cost: %2f\n", address, cost)
		fmt.Println("=================================================================")
	case sms:
		fmt.Printf("Report: The sms is going to %s and it will cost: %2f\n", address, cost)
		fmt.Println("================================================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("================================================================")
	}
}

// Main function
func main() {
	e := email{
		isSubscribed: true,
		body:         "Hello there!",
	}

	getEmail(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}

	getEmail(e, e)
	e = email{
		isSubscribed: true,
		body:         "I will send it to you tomorrow",
	}
	getEmail(e, e)

	sm := sms{
		isSubscribed:  true,
		toPhoneNumber: "070000000",
		body:          "Where are you?",
	}

	getSms(sm, sm)

	sm = sms{
		isSubscribed:  true,
		toPhoneNumber: "070000000",
		body:          "I am at the Club",
	}

	getSms(sm, sm)

	sm = sms{
		isSubscribed:  false,
		toPhoneNumber: "070000000",
		body:          "Okay please come how now?",
	}

	getSms(sm, sm)

	getEmailOrSms(email{
		isSubscribed: true,
		body:         "Hello baby!",
		toAddress:    "janedoe@gmail.com",
	})

	getEmailOrSms(email{
		isSubscribed: false,
		body:         "Hello there, my friend!",
		toAddress:    "johndoe@gmail.com",
	})

	getEmailOrSms(sms{
		isSubscribed:  true,
		body:          "What would you want to have for dinner?",
		toPhoneNumber: "0790909990",
	})

	getEmailOrSms(sms{
		isSubscribed:  false,
		body:          "I would love to have some chicken soup and cassava?",
		toPhoneNumber: "0790909990",
	})

	getEmailOrSms(invalid{})
}
