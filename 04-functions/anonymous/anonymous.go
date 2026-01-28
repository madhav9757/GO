package main

import "fmt"

func main() {

	// Anonymous function assigned to a variable
	plus := func(a, b int) int {
		return a + b
	}

	res := plus(3, 5)
	fmt.Println("3 + 5 =", res)

	// Immediately invoked anonymous function (IIFE)
	func(name string) {
		fmt.Println("Hello,", name)
	}("Go Developer")
}
