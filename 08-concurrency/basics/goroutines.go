package main

import (
	"fmt"
	"time"
)

// printNumbers prints numbers with a delay
func printNumbers(prefix string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("--- Goroutines Basic ---")

	// 1. Direct function call (Synchronous)
	// This will block until it finishes
	fmt.Println("Synchronous call:")
	printNumbers("Sync")

	// 2. Goroutine call (Asynchronous)
	// This runs concurrently with the main function
	fmt.Println("\nAsynchronous calls:")
	go printNumbers("Goroutine 1")
	go printNumbers("Goroutine 2")

	// 3. Anonymous self-executing goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from Anonymous Goroutine!")

	// NOTE: We sleep here to wait for goroutines to finish.
	// In production, we use WaitGroups (covered in sync section).
	// If main returns, all goroutines are killed immediately.
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
