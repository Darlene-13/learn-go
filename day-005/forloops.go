package main

import "fmt"

//The major difference is that we do not use parenthesis here
// for INITIAL; CONDITION; AFTER{do something}

//Initial: This is run once at the begining of the loop and can create variables within the scope
//Condition: This is checked before each iteration if the condition does not pass the loop breaks
//After: This is run after each iteration
// for i = 0; i < 10; i++{fmt.Println(i)}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.00 + float64(0.01*float64(i))
	}
	return totalCost

}

// for loop with no condition
func bulkSendMessages(threshold float64) float64 {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.00 + float64(0.01*float64(i))

		if totalCost > threshold {
			return float64(i)
		}
	}

}

func test(numMessages int, threshold float64) {
	fmt.Println("===========================================")
	fmt.Printf("Sending %v messages...\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Cost of sending %v messages.\n", numMessages)
	fmt.Printf("Bulk send complete! Cost = %2.f\n", cost)

	fmt.Println("===========================================")
	fmt.Printf("The threshold is: %v \n", threshold)
	messages := bulkSendMessages(threshold)
	fmt.Printf("For %v threshold you can send %v messages.\n", threshold, messages)

}

func main() {
	test(10, 10.5)
	test(12, 25.5)
}
