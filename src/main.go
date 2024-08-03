package main

import (
	"fmt"
	"net/http"
	"time"
)

// greet function handles HTTP requests and responds with "Hello World!" and the current time.
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// Set up the HTTP server and listen on port 8080.
	http.HandleFunc("/", greet)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
