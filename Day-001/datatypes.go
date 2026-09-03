package main

import "fmt"

// Basic data types in go.
//1. Float: float 16, float 32 etc.
//2. Integers: Unsigned which do not take negatives and Signed integers which actually take negatives.
// Additionally, when it comes to integers we represent them based on the desired memory or space allocation: int8, int16, int32 they all take it different bits at the end of the day
//3. Boolean
//4. String
//5. Byte
// 6. rune -- this is an alias for int32
//7. complex

//DECLARING A VARIABLE
//1. Var {name of the variable} data type, and it could be used for global variables
// := we could use that, and it actually infers to the data type automatically.
// := {name of the variable}, it is mostly used for local variables

func main() {

	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var username string

	//Call the function variables
	variables()

	permissiveness()

	multivariables()

	accountage()

	constants()

	computedcontants()

	stringformatting()

	fmt.Printf("%v %f %v %q\n",
		smsSendingLimit,
		costPerSms,
		hasPermission,
		username,
	)

}

// DECLARING VARIABLES THE SIMPLE WAY

func variables() {

	congrats := "Happy Birthday Darlene!"

	fmt.Println(congrats)
}

func permissiveness() {
	permissionPerText := 2

	fmt.Println("The type of permission per text ", permissionPerText)
}

// We can declare multiple variables on the same line
func multivariables() {

	averageOpenRate, displayMessage := .23, "is the Average Open Rate"

	fmt.Println(averageOpenRate, displayMessage)
}

// We could convert data types of variables using type casting

func accountage() {

	accountAge := 2.6

	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}

// Nibbles is actually 4 bits

// CONSTANTS IN GO
func constants() {
	const premiumPlanName string = "Premium plan"
	const basicPlanName string = "Basic plan"

	fmt.Println("Available Plans are:", premiumPlanName, "and ", basicPlanName)
}

// In go all values that are constants must be computed at compile time
func computedcontants() {

	const firstName string = "Darlene"
	const secondName string = "Wendy"
	const fullName string = firstName + " " + secondName
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("My name is: " + fullName)
	fmt.Println(secondsInMinute, minutesInHour, secondsInHour)
}

//STRING FORMATTING
//We can use fmt.Printf or fmt.Sprintf
//%v %s replaced with the actual values

func stringformatting() {

	const name = "Darlene Wendy"
	const openRate = 23.5

	fmt.Printf("My name is %s and I come from Bungoma!\n", name)
	fmt.Printf("The open rate is %v !\n", openRate)
}
