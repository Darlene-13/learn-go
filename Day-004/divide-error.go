package main

import (
	"errors"
	"fmt"
)

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("Cannot divide %v by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("Cannot divide %v by zero", divisor)
	}

	if divisor > dividend {
	}

	return dividend / divisor, nil
}
func main() {

}

// Alternatively we could use the error.New to directly create the  errors without defining a struct
func divide1(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by zero")
	}

	return x / y, nil
}
