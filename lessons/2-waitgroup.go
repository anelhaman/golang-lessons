package main

import (
	"fmt"
	"sync"
)

// Define a struct to encapsulate the accumulator and synchronization logic
type Accumulator struct {
	mu          sync.Mutex // Mutex to synchronize access to the counter
	accumulated int        // The accumulated value
}

// Method to increment the accumulated value with a WaitGroup using goroutines
func (a *Accumulator) incrementWithWaitGroup(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		a.mu.Lock()     // Lock to ensure safe access to the accumulated value
		a.accumulated++ // Increment the accumulated value
		a.mu.Unlock()   // Unlock after incrementing
	}
	wg.Done() // Signal that this goroutine is done
}

// Method to increment the accumulated value without using WaitGroup
func (a *Accumulator) incrementWithoutWaitGroup() {
	for i := 0; i < 1000; i++ {
		a.mu.Lock()     // Lock to ensure safe access to the accumulated value
		a.accumulated++ // Increment the accumulated value
		a.mu.Unlock()   // Unlock after incrementing
	}
}

func main() {
	// Part 1: Without WaitGroup (no synchronization for goroutines)
	accumulator := &Accumulator{} // Create an instance of Accumulator
	for i := 0; i < 5; i++ {
		go accumulator.incrementWithoutWaitGroup() // Start goroutines without WaitGroup
	}

	// Sleep for a while to allow goroutines to finish (not the best practice)
	// This is just to simulate waiting for goroutines without using WaitGroup
	// In real code, this is not recommended as it introduces race conditions.
	// time.Sleep(1 * time.Second)
	fmt.Println("Final accumulated value (without WaitGroup):", accumulator.accumulated)

	// Part 2: With WaitGroup (using sync.WaitGroup to wait for goroutines)
	accumulator = &Accumulator{} // Create a new instance of Accumulator
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go accumulator.incrementWithWaitGroup(&wg) // Start goroutines with WaitGroup
	}
	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Final accumulated value (with WaitGroup):", accumulator.accumulated)
}
