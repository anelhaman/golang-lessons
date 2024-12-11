package main

import (
	"fmt"
	"sync"
)

// Define a struct to encapsulate the counter
type Counter struct {
	mu      sync.Mutex // Mutex to synchronize access to the counter
	counter int        // The counter value
}

// Method to increment the counter safely using a Mutex
func (c *Counter) Increment() {
	c.mu.Lock()   // Lock to ensure safe access to the counter
	c.counter++   // Increment the counter
	c.mu.Unlock() // Unlock after incrementing
}

// Method to get the current counter value safely using a Mutex
func (c *Counter) GetValue() int {
	c.mu.Lock()         // Lock to ensure safe access to the counter
	defer c.mu.Unlock() // Unlock after getting the value
	return c.counter
}

// Function to increment the counter without using goroutines
func (c *Counter) incrementWithoutGoroutines() {
	for i := 0; i < 1000; i++ {
		c.Increment() // Increment the counter using method
	}
}

// Function to increment the counter using goroutines
func (c *Counter) incrementWithGoroutines(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		c.Increment() // Increment the counter using method
	}
}

func main() {
	// Part 1: Without Goroutines (Single-threaded)
	counter := &Counter{}                // Create a Counter instance
	counter.incrementWithoutGoroutines() // Increment without goroutines
	fmt.Println("Final counter value (without goroutines):", counter.GetValue())

	// Part 2: With Goroutines (Concurrent processing)
	counter = &Counter{} // Create a Counter instance
	var wg sync.WaitGroup

	// Start 5 goroutines to increment the counter concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter.incrementWithGoroutines(&wg) // Increment with goroutines
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter value (with goroutines):", counter.GetValue())
}
