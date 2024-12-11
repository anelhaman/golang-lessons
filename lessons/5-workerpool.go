package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Define a JobProcessor struct with a counter field and a Mutex for synchronization
type JobProcessor struct {
	mu      sync.Mutex // Mutex for synchronization
	counter int        // Counter for processed jobs
}

// Method to process a job and increment the counter
func (jp *JobProcessor) processJob(job string) {
	time.Sleep(100 * time.Millisecond)
	log.Println("Processed:", job)
	jp.mu.Lock()   // Lock to prevent race condition when incrementing counter
	jp.counter++   // Increment the job counter
	jp.mu.Unlock() // Unlock the mutex after the operation
}

// Method to simulate worker pool processing
func (jp *JobProcessor) workerPool(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		jp.processJob(job) // Process the job using the JobProcessor method
	}
}

// Method to process jobs without a worker pool (direct processing)
func (jp *JobProcessor) processJobsWithoutPool(jobs []string) {
	for _, job := range jobs {
		jp.processJob(job) // Directly process each job
	}
}

func main() {

	// List of jobs (can be any tasks you want to process)
	jobs := []string{"job1", "job2", "job3", "job4", "job5", "job6", "job7", "job8", "job9", "job10"}

	// --- Without Worker Pool ---
	jp := &JobProcessor{} // Create an instance of JobProcessor for direct processing
	jp.counter = 0        // Reset counter
	startTime := time.Now()
	jp.processJobsWithoutPool(jobs) // Process jobs directly
	log.Printf("Processed %d jobs without Worker Pool in %s\n", jp.counter, time.Since(startTime))
	fmt.Println("========================")

	// --- With Worker Pool ---
	jp.counter = 0 // Reset counter for worker pool
	jobsChan := make(chan string, len(jobs))
	var wg sync.WaitGroup
	numWorkers := 10 // Number of workers in the pool

	// Add jobs to the channel
	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	// Create worker pool with goroutines
	startTime = time.Now()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go jp.workerPool(jobsChan, &wg) // Start worker pool
	}

	wg.Wait() // Wait for all workers to finish
	log.Printf("Processed %d jobs with Worker Pool in %s\n", jp.counter, time.Since(startTime))

}
