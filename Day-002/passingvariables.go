package main

import "fmt"

func main() {

	//First function
	x := 5
	increment(x)

	//Second function
	sendSoFar := 430
	const sendsToAdd = 25
	sendSoFar = incrementSends(sendSoFar, sendsToAdd)

	//Third function
	firstName := "Darlene"
	secondName := "Wendy"
	myFirstName, _ := getNames(firstName, secondName)
	_, mySecondName := getNames(firstName, secondName)

	//Fourth function
	cordX := 98
	cordY := 99
	cordX, cordY = getCoord(cordX, cordY)

	//Fifth function
	carRental, adult, drinking := yearsUntilEvents(19)

	// 6th function
	checkAge(15)

	fmt.Println("===========================================")
	fmt.Println(x)
	fmt.Println("===========================================")
	fmt.Println("You have sent", sendSoFar, "messages so Far")
	fmt.Println("===========================================")
	fmt.Println("My first name is", myFirstName)
	fmt.Println("My second name is", mySecondName)
	fmt.Println("===========================================")

	fmt.Println("The co-ordinates for X and Y are", cordX, "and", cordY)
	fmt.Println("===========================================")

	fmt.Println("Years until car rental:", carRental)
	fmt.Println("===========================================")
	fmt.Println("Years until adulthood:", adult)
	fmt.Println("===========================================")
	fmt.Println("Years until drinking:", drinking)
	fmt.Println("===========================================")

}

func increment(x int) {
	x++
}

func incrementSends(sendSoFar, sendsToAdd int) int {
	sendSoFar = sendSoFar + sendsToAdd
	return sendSoFar
}

// When we have multiple return values we wrap the return types in brackets as well
func getNames(firstName, secondName string) (string, string) {
	return firstName, secondName
}

// Go does not allow one to have unused variables, we can ignore the value using an underscore _
func getCoord(cordX, cordY int) (int, int) {
	return cordX, cordX
}

func yearsUntilEvents(age int) (int, int, int) {
	yearsUntilAdult := 18 - age

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking := 21 - age

	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental := 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking
}

// Guard clause example
func checkAge(age int) {
	if age < 0 {
		fmt.Println("Invalid age")
		return
	}

	if age < 18 {
		fmt.Println("You are a minor")
		return
	}

	fmt.Println("You are an adult")
}

// Use named returned for small functions
//Use named values when there are many values of the same type to be returned
// Guard clauses provide a linear approach to logic trees, and it embodies a clean code
