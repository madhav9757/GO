package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// Lock so only one goroutine at a time can access the map c.v.
	c.mu.Lock()
	// Unlock is guaranteed to be called using defer
	defer c.mu.Unlock()

	// Use a slight delay to prove concurrency issues without lock (if we removed it)
	// val := c.v[key]
	// time.Sleep(1 * time.Millisecond)
	// c.v[key] = val + 1
	c.v[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	fmt.Println("--- Mutex (Mutual Exclusion) ---")

	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup

	// 1000 goroutines trying to increment the same key at once
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("somekey")
		}()
	}

	wg.Wait()
	fmt.Println("Expected value: 1000")
	fmt.Println("Actual value:  ", c.Value("somekey"))
}
