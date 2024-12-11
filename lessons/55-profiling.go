package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof
)

func main() {
	// Start the pprof server in a separate goroutine
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil)) // Start pprof server on port 6060
	}()

	// Simulate some work in your app
	fmt.Println("Application is running...")

	// Keeping the main thread alive
	select {} // Block forever
}

// go tool pprof http://localhost:6060/debug/pprof/heap
// (pprof) top
// (pprof) list main.FunctionA
// (pprof) graph
// (pprof) web
// (pprof) heap
