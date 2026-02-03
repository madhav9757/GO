package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- Buffered Channels ---")

	// Create a buffered channel with capacity 2
	// Sends will NOT block unless the buffer is full (2 items)
	messages := make(chan string, 2)

	messages <- "Buffered"
	messages <- "Channel"
	fmt.Println("Successfully sent 2 messages without a receiver")

	// messages <- "Full" // This would BLOCK (deadlock if on main thread) because buffer is full

	fmt.Println("Reading from buffer:")
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// Select statement for non-blocking operations
	fmt.Println("\n--- Select Statement ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// select waits for ONE, randomly picks if multiple are ready
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}
