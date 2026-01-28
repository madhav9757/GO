package main

import "fmt"

// standard function
func add(a int, b int) int {
	return a + b
}

// function with omitted type for consecutive arguments of same type
func subtract(a, b int) int {
	return a - b
}

// named return values
func calculate(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return // naked return
}

func main() {
	res := add(5, 3)
	fmt.Println("5 + 3 =", res)

	res = subtract(10, 4)
	fmt.Println("10 - 4 =", res)

	s, d := calculate(10, 5)
	fmt.Println("10 + 5 =", s)
	fmt.Println("10 - 5 =", d)
}
