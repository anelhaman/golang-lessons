package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct to encapsulate the state (counter)
type FileProcessor struct {
	counter int
	mu      sync.Mutex
}

// Method to validate a file
func (fp *FileProcessor) validateFile(file string, ch chan<- string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Validated:", file)
	ch <- file
}

// Method to transform a file
func (fp *FileProcessor) transformFile(file string, ch chan<- string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Transformed:", file)
	ch <- file
}

// Method to store the result (with Mutex)
func (fp *FileProcessor) storeResult(file string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	fp.mu.Lock()
	fp.counter++
	fmt.Println("Stored result for:", file)
	fp.mu.Unlock()
}

// Worker function to handle job processing
func (fp *FileProcessor) worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	for file := range jobs {
		fmt.Printf("Worker %d processing file: %s\n", id, file)
		// Simulating the three steps of file processing
		ch := make(chan string)
		go fp.validateFile(file, ch)
		validatedFile := <-ch
		go fp.transformFile(validatedFile, ch)
		transformedFile := <-ch
		fp.storeResult(transformedFile, wg)
		results <- transformedFile
	}
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
	jobs := make(chan string, len(files))
	results := make(chan string, len(files))
	var wg sync.WaitGroup

	// Instantiate FileProcessor
	fp := &FileProcessor{}

	// Creating worker pool with 3 workers
	for w := 1; w <= 3; w++ {
		go fp.worker(w, jobs, results, &wg)
	}

	// Sending jobs to workers
	for _, file := range files {
		wg.Add(1) // Add to WaitGroup before sending jobs
		jobs <- file
	}

	close(jobs) // Close jobs channel after all jobs have been sent

	// Wait for all workers to finish
	wg.Wait()

	// Close results channel after all jobs are processed
	close(results)

	// Output results
	for result := range results {
		fmt.Println("Processed:", result)
	}

	// Final count of processed files
	fmt.Println("Total processed files:", fp.counter)
}
