package main

import (
	"fmt"
	"sync"
)

// Define a TaskCounter struct to encapsulate the counter and related methods
type TaskCounter struct {
	mu      sync.Mutex // Mutex to synchronize access to the counter
	counter int        // The counter value
}

// Method to increment the counter using channels
func (tc *TaskCounter) incrementWithChannel(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- 1 // Send increment value to the channel
		// Debug: Show goroutine sending value to channel
		fmt.Println("Sent value to channel: ", i)
	}
}

// Method to increment the counter without using channels
func (tc *TaskCounter) incrementWithoutChannel(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		tc.mu.Lock()   // Lock to ensure safe access to the counter
		tc.counter++   // Directly increment the counter
		tc.mu.Unlock() // Unlock after incrementing
		// Debug: Show direct counter increment
		fmt.Println("Direct increment to counter: ", i)
	}
}

func main() {
	// Part 1: Without Channel (no synchronization or communication)
	taskCounter := &TaskCounter{} // Create an instance of TaskCounter
	var wg sync.WaitGroup

	// Start 2 goroutines to increment the counter directly
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go taskCounter.incrementWithoutChannel(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final result without using channels
	fmt.Println("Final counter value (without channel):", taskCounter.counter)
	fmt.Println("========================")

	// Part 2: With Channel (using channel for communication between goroutines)
	taskCounter = &TaskCounter{} // Reset the counter
	ch := make(chan int, 5)      // Buffered channel
	wg = sync.WaitGroup{}

	// Start 2 goroutines to increment using the channel
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go taskCounter.incrementWithChannel(ch, &wg)
	}

	// Goroutine to receive from the channel and increment the counter
	go func() {
		wg.Wait() // Wait for all goroutines to finish sending values
		close(ch) // Close the channel once done
	}()

	// Receiving values from the channel and incrementing the counter
	for val := range ch {
		taskCounter.mu.Lock()      // Lock to safely access the counter
		taskCounter.counter += val // Increment the counter
		taskCounter.mu.Unlock()    // Unlock after incrementing
		// Debug: Show counter increment when receiving from channel
		fmt.Println("Received value", val, "from channel and incremented counter:", taskCounter.counter)
	}

	// Print the final result with using channels
	fmt.Println("Final counter value (with channel):", taskCounter.counter)
}
