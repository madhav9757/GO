package main

import "fmt"

func main() {
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// While-like loop (Go has only 'for')
	n := 1
	for n <= 5 {
		fmt.Println("While-like loop:", n)
		n++
	}

	// Infinite loop with break
	count := 1
	for {
		if count > 3 {
			break
		}
		fmt.Println("Infinite loop with break:", count)
		count++
	}
}
