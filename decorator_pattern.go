package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Base handler function (without logging)
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// LoggingDecorator is a higher-order function that wraps an http.HandlerFunc
func LoggingDecorator(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the original handler
		handler(w, r)

		// Log request details
		log.Printf("[LOG] %s %s | Execution Time: %v\n", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	// Wrap the handler with LoggingDecorator
	http.HandleFunc("/", LoggingDecorator(HelloHandler))
	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
