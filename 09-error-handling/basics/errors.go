package main

import (
	"errors"
	"fmt"
	"math"
)

// In Go, errors are values and are handled explicitly.
// By convention, errors are the last return value of a function.

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New is used to create a simple error
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// fmt.Errorf allows us to format error messages
		return 0, fmt.Errorf("math error: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println("--- Basic Error Handling ---")

	// Case 1: Successful division
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %g\n", res)
	}

	// Case 2: Division by zero
	res, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error caught:", err)
	} else {
		fmt.Printf("10 / 0 = %g\n", res)
	}

	// Case 3: Formatted error
	res, err = sqrt(-1)
	if err != nil {
		fmt.Println("Error caught:", err)
	} else {
		fmt.Printf("sqrt(-1) = %g\n", res)
	}
}
