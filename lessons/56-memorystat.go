package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Force garbage collection to demonstrate memory stats
	runtime.GC()

	// Get memory statistics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Display memory stats
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("HeapAlloc = %v MiB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("HeapSys = %v MiB\n", m.HeapSys/1024/1024)
	fmt.Printf("HeapIdle = %v MiB\n", m.HeapIdle/1024/1024)
	fmt.Printf("HeapInuse = %v MiB\n", m.HeapInuse/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
