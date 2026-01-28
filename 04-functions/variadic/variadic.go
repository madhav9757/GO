package main

import "fmt"

// sum accepts a variable number of integers
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Call with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// Call with a slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
