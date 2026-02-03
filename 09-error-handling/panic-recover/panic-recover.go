package main

import "fmt"

// Panic is used when something unrecoverable happens.
// Recover can stop a panic from crashing the program.
// Recover is only useful inside a deferred function.

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func riskOperation() {
	defer handlePanic()

	fmt.Println("Executing risky operation...")

	// Something goes wrong
	panic("something went horribly wrong!")

	// This line will never be reached
	fmt.Println("This will never print")
}

func main() {
	fmt.Println("--- Panic and Recover ---")

	fmt.Println("Starting program")

	riskOperation()

	fmt.Println("Program continued execution after recovery")

	// Example of a runtime panic (slice out of bounds)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from runtime panic:", r)
		}
	}()

	fmt.Println("\nAttempting out of bounds access...")
	nums := []int{1, 2, 3}
	_ = nums[10] // This will trigger a panic

	fmt.Println("This line will never print because we deferred recover in main too late for this specific sequence if not handled.")
}
