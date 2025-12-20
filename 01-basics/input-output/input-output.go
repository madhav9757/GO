package main

import "fmt"

func input_output() {
	var name string
	var age int

	// Input
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// Output
	fmt.Println("\n--- User Details ---")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
