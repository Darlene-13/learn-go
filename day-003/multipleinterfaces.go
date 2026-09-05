package main

import "fmt"

// STRUCTS
type email struct {
	isSubscribed bool
	body         string
	toAddress    string
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

// HELPER FUNCTION
func getEmail(e expense, p printer) {
	fmt.Println("=============================================")
	fmt.Printf("Printing with cost %2f ........\n", e.cost())
	p.print()
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

}
