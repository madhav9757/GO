package main

import (
	"context"
	"fmt"
	"time"
)

func operation(ctx context.Context) {
	// Simulate work using a select statement
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Operation completed successfully")
	case <-ctx.Done(): // Context cancelled or timed out
		err := ctx.Err()
		fmt.Println("Operation stopped:", err)
	}
}

func main() {
	fmt.Println("--- Context Pattern ---")

	// 1. WithTimeout
	// Creating a context that expires after 200ms
	// This is too short for our 500ms operation, so it should timeout.
	fmt.Println("1. Testing Timeout:")
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // Always call cancel for cleanup

	operation(ctx)

	// 2. WithCancel
	// Manually cancelling
	fmt.Println("\n2. Testing Manual Cancel:")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel2() // Cancel after 100ms
	}()

	operation(ctx2)

	// 3. Successful run
	fmt.Println("\n3. Testing Success:")
	ctx3, cancel3 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel3()

	operation(ctx3)
}
