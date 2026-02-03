package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate expensive task
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("--- Worker Pool Pattern ---")

	const numJobs = 10
	const numWorkers = 3

	// Buffered channels
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start 3 workers
	// They are now blocked waiting for jobs
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send 10 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to signal that's all the work

	// Collect 10 results
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("Result: %d\n", res)
	}
	// We don't necessarily close results here because main finishes anyway
}
