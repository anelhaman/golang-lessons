package main

import (
	"fmt"
	"sync"
)

// Define a JobCounter struct with a counter field and mutex for synchronization
type JobCounter struct {
	mu      sync.Mutex // Mutex to synchronize access to the counter
	counter int        // The counter value
}

// Method to increment the counter with Mutex synchronization
func (jc *JobCounter) incrementWithMutex(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when goroutine finishes
	for i := 0; i < 10; i++ {
		jc.mu.Lock() // Lock the mutex to prevent race condition
		jc.counter++ // Increment the counter
		fmt.Println("Counter value (with Mutex):", jc.counter)
		jc.mu.Unlock() // Unlock the mutex after the operation
	}
}

// Method to increment the counter without Mutex (race condition possible)
func (jc *JobCounter) incrementWithoutMutex(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when goroutine finishes
	for i := 0; i < 10; i++ {
		jc.counter++ // Increment the counter without synchronization (race condition may occur)
		fmt.Println("Counter value (without Mutex):", jc.counter)
	}
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	// Part 1: Without Mutex (race condition may occur)
	jc1 := &JobCounter{} // Create an instance of JobCounter for part 1
	for i := 0; i < 2; i++ {
		wg.Add(1)                         // Add a goroutine to WaitGroup
		go jc1.incrementWithoutMutex(&wg) // Start goroutines without mutex
	}

	// Wait for all goroutines to finish before printing the result
	wg.Wait()
	fmt.Println("Final counter value (without Mutex):", jc1.counter)
	fmt.Println("========================")

	// Part 2: With Mutex (synchronized increment to prevent race condition)
	jc2 := &JobCounter{} // Create a new instance of JobCounter for part 2
	for i := 0; i < 2; i++ {
		wg.Add(1)                      // Add a goroutine to WaitGroup
		go jc2.incrementWithMutex(&wg) // Start goroutines with mutex for synchronization
	}

	// Wait for all goroutines to finish before printing the result
	wg.Wait()
	fmt.Println("Final counter value (with Mutex):", jc2.counter)

}
