package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	fmt.Println("Sender: sending 'Hello'...")
	ch <- "Hello" // This blocks until receiver is ready
	fmt.Println("Sender: sent 'Hello'")

	fmt.Println("Sender: sending 'World'...")
	ch <- "World"
	fmt.Println("Sender: sent 'World'")

	close(ch) // Close channel to signal no more data
}

func main() {
	fmt.Println("--- Unbuffered Channels ---")

	// Create an unbuffered channel
	// Unbuffered means capacity is 0. Blocks until both sender and receiver are ready.
	messageChannel := make(chan string)

	go sender(messageChannel)

	time.Sleep(500 * time.Millisecond) // Simulate some work in main
	fmt.Println("Main: Ready to receive")

	// Receive data
	// msg := <-messageChannel // Blocking receive

	// Range over channel receives until closed
	for msg := range messageChannel {
		fmt.Printf("Received: %s\n", msg)
	}

	fmt.Println("Channel closed, done.")
}
