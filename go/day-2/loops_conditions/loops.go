package main

import "fmt"

func runLoopsDemo() {
	fmt.Println("\n--- FOR LOOP DEMO ---")

	// Classic for loop
	fmt.Println("Numbers from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While-style loop
	fmt.Println("Counting down from 5:")
	count := 5
	for count > 0 {
		fmt.Println(count)
		count--
	}

	// For-range loop
	fruits := []string{"Apple", "Banana", "Mango"}
	fmt.Println("Fruits list:")
	for index, fruit := range fruits {
		fmt.Println(index+1, ":", fruit)
	}
}
