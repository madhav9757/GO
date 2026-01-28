package main

import "fmt"

// vals returns two integers
func vals() (int, int) {
	return 3, 7
}

// swap returns two strings in reverse order
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// assigning both return values
	a, b := vals()
	fmt.Println("vals (a, b):", a, b)

	// using blank identifier to ignore one value
	_, c := vals()
	fmt.Println("vals (_, c):", c)

	// swapping values
	f, s := swap("hello", "world")
	fmt.Println("swap:", f, s)
}
