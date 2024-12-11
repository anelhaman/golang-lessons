package main

import (
	"sync"

	. "github.com/onsi/ginkgo/v2" // Ginkgo for testing
	. "github.com/onsi/gomega"    // Gomega for assertions
)

var _ = Describe("FileProcessor", func() {
	var fp *FileProcessor
	var wg sync.WaitGroup
	var jobs chan string
	var results chan string

	BeforeEach(func() {
		fp = &FileProcessor{}
		jobs = make(chan string, 5)
		results = make(chan string, 5)
	})

	Describe("File Validation", func() {
		It("should validate the file correctly", func() {
			ch := make(chan string)
			go fp.validateFile("file1.txt", ch)
			result := <-ch
			Expect(result).To(Equal("file1.txt"))
		})
	})

	Describe("File Transformation", func() {
		It("should transform the file correctly", func() {
			ch := make(chan string)
			go fp.transformFile("file2.txt", ch)
			result := <-ch
			Expect(result).To(Equal("file2.txt"))
		})
	})

	Describe("Storing Results", func() {
		It("should increment the counter when a result is stored", func() {
			wg.Add(1)
			go fp.storeResult("file3.txt", &wg)
			wg.Wait()
			Expect(fp.counter).To(Equal(1))
		})
	})

	Describe("Complete File Processing", func() {
		It("should process the files correctly through all steps", func() {
			files := []string{"file4.txt", "file5.txt"}
			for _, file := range files {
				wg.Add(1)
				jobs <- file
			}

			go fp.worker(1, jobs, results, &wg)
			close(jobs)

			wg.Wait()
			close(results)

			for result := range results {
				Expect(result).To(BeElementOf(files))
			}

			Expect(fp.counter).To(Equal(2))
		})
	})

})
