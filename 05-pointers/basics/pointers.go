package main

import "fmt"

func main() {
	// 1. Basic Variable
	x := 42
	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Address of x: %v\n", &x) // Using & to get the memory address

	// 2. Declaring a Pointer
	// A pointer holds the memory address of a value.
	// *int is the type "pointer to an integer"
	var p *int

	fmt.Printf("Initial value of p: %v\n", p) // Pointers default to nil

	// 3. Initializing a Pointer
	p = &x // p now stores the address of x
	fmt.Printf("Value of p (address): %v\n", p)
	fmt.Printf("Value pointed to by p: %v\n", *p) // Using * for dereferencing

	// 4. Modifying value through pointer
	*p = 100 // Changes the value at the address p stores
	fmt.Printf("New value of x: %v\n", x)

	// 5. Using the new() function
	// new(T) allocates memory for a type T and returns a pointer to it (*T)
	ptr := new(int)
	fmt.Printf("Value of ptr: %v\n", ptr)
	fmt.Printf("Value pointed to by ptr: %v\n", *ptr) // Defaults to zero value (0)
	*ptr = 500
	fmt.Printf("Updated value of *ptr: %v\n", *ptr)
}
