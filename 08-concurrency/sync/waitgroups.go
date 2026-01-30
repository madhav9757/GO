package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Schedule the call to Done to tell the WaitGroup that this worker is done.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("--- WaitGroups ---")

	var wg sync.WaitGroup

	// Launch 5 workers
	for i := 1; i <= 5; i++ {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they're done.
	fmt.Println("Main: Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("Main: All workers done")
}
