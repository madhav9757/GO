package main

import "fmt"

func runIfElseDemo() {
	fmt.Println("\n--- IF-ELSE DEMO ---")

	age := 18

	if age < 18 {
		fmt.Println("You are a minor")
	} else if age == 18 {
		fmt.Println("You just became an adult!")
	} else {
		fmt.Println("You are an adult")
	}
}
