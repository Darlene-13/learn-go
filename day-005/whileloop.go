package main

import (
	"fmt"
)

//A while loop in go is just a for loop with a condition
// platHeight := 1

//for platHeight < 5{
//	fmt.Printf("Still growing! Current height is:\n", platHeight)
//	platHeight++
//}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	// Code should exit the loop if the cost of the next messages would go over the max cost
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}

	return maxMessagesToSend
}

func test1(costMultiplier float64, maxCostInPennies int) {
	fmt.Println("==========================================")
	fmt.Printf("Sending %v messages...\n", getMaxMessagesToSend(costMultiplier, 100))
}

func main() {
	fmt.Println("While loop detected")
	test1(5, 100)

	fizzbuzz()
	continuekeyword()
	breakkeyword()

}

//Modulo operator calculates remainders
// When using modulo we return the remainder from a division
//LOGICAL AND OPERATOR
//true && false //false
//true && true // true

//LOGICAL OR OPERATOR
//true || false // true
//false || false //false

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

//CONTINUE KEYWORD
// Continue keyword stops the current iteration of a loop and continues to the next iteration.
// It is a powerful way to use guard clause

//BREAK KEYWORD
// The break keyword

func continuekeyword() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func breakkeyword() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			break
		}
		fmt.Println(i)
	}
}

func printPrimes(max int) {
	for n := 2; n <= max+1; n++ {
		if n == 2 {
			fmt.Printf("%d ", n)
			continue
		}

		if n%2 == 0 {
			// Skip
			continue
		}

		isPrime := true

		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				break
				isPrime = false
			}
		}

		if !isPrime {
			continue
		}

		fmt.Println(n)
	}

}
