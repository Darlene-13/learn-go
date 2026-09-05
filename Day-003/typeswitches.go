// Type switch makes it easy to do type assertions in series
// It is similar to regular switch case but then a bit different in that the case is the type and not the value
package main

import "fmt"

// TYPE ASSERTION
func getExpenseReport2(e expense2) (string, float64) {
	switch v := e.(type) {
	case email2:
		return v.toAddress, v.cost()
	case sms2:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

// STRUCTS

type email2 struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms2 struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid2 struct{}

// INTERFACES

type expense2 interface {
	cost() float64
}

type printer2 interface {
	print()
}

// METHODS

func (e email2) cost() float64 {

	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}

	return 0.01 * float64(len(e.body))
}

func (e email2) print() {
	fmt.Println(e.body)
}

func (sm sms2) cost() float64 {

	if !sm.isSubscribed {
		return 0.05 * float64(len(sm.body))
	}

	return 0.01 * float64(len(sm.body))
}

func (sm sms2) print() {
	fmt.Println(sm.body)
}

func (i invalid2) cost() float64 {
	return 0.0
}

// HELPER FUNCTIONS

func getEmail2(e expense2, p printer2) {

	fmt.Println("=============================================")

	fmt.Printf(
		"Printing with cost %.2f ........\n",
		e.cost(),
	)

	p.print()
}

func getSms2(sm expense2, s printer2) {

	fmt.Println("=============================================")

	fmt.Printf(
		"Printing with cost %.2f ........\n",
		sm.cost(),
	)

	s.print()
}

// TYPE SWITCH

func getEmailOrSms2(e expense2) {

	address, cost := getExpenseReport2(e)

	switch e.(type) {

	case email2:

		fmt.Printf(
			"Report: The email is going to %s and it will cost: %.2f\n",
			address,
			cost,
		)

		fmt.Println("=================================================================")

	case sms2:

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

	e2 := email2{
		isSubscribed: true,
		body:         "Hello there!",
	}

	getEmail2(e2, e2)

	// EMAIL 2

	e2 = email2{
		isSubscribed: false,
		body:         "I want my money back",
	}

	getEmail2(e2, e2)

	// EMAIL 3

	e2 = email2{
		isSubscribed: true,
		body:         "I will send it to you tomorrow",
	}

	getEmail2(e2, e2)

	// SMS 1

	sm2 := sms2{
		isSubscribed:  true,
		toPhoneNumber: "070000000",
		body:          "Where are you?",
	}

	getSms2(sm2, sm2)

	// SMS 2

	sm2 = sms2{
		isSubscribed:  true,
		toPhoneNumber: "070000000",
		body:          "I am at the Club",
	}

	getSms2(sm2, sm2)

	// SMS 3

	sm2 = sms2{
		isSubscribed:  false,
		toPhoneNumber: "070000000",
		body:          "Okay please come how now?",
	}

	getSms2(sm2, sm2)

	// EMAIL REPORT 1

	getEmailOrSms2(email2{
		isSubscribed: true,
		body:         "Hello baby!",
		toAddress:    "janedoe@gmail.com",
	})

	// EMAIL REPORT 2

	getEmailOrSms2(email2{
		isSubscribed: false,
		body:         "Hello there, my friend!",
		toAddress:    "johndoe@gmail.com",
	})

	// SMS REPORT 1

	getEmailOrSms2(sms2{
		isSubscribed:  true,
		body:          "What would you want to have for dinner?",
		toPhoneNumber: "0790909990",
	})

	// SMS REPORT 2

	getEmailOrSms2(sms2{
		isSubscribed:  false,
		body:          "I would love to have some chicken soup and cassava?",
		toPhoneNumber: "0790909990",
	})

	// INVALID EXPENSE

	getEmailOrSms2(invalid2{})
}
