package main

import "fmt"

// updateValue takes an integer (Pass by Value)
// It receives a copy of the argument.
func updateValue(val int) {
	val = val + 10
	fmt.Printf("Inside updateValue: %d\n", val)
}

// updateWithPointer takes a pointer to an integer (Pass by Reference/Pointer)
// It receives the memory address.
func updateWithPointer(ptr *int) {
	if ptr != nil {
		*ptr = *ptr + 10
		fmt.Printf("Inside updateWithPointer: %d\n", *ptr)
	}
}

func main() {
	number := 5

	fmt.Printf("Original number: %d\n", number)

	// Case 1: Pass by Value
	fmt.Println("--- Calling updateValue ---")
	updateValue(number)
	fmt.Printf("After updateValue, back in main: %d (No change!)\n", number)

	// Case 2: Pass by Pointer
	fmt.Println("--- Calling updateWithPointer ---")
	// We pass the address of 'number'
	updateWithPointer(&number)
	fmt.Printf("After updateWithPointer, back in main: %d (Changed!)\n", number)

	// Complex types like slices and maps are technically passed by value,
	// but they contain pointers to the underlying data, so they behave
	// somewhat like they are passed by reference.
}
