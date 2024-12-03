package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	// Create a new HTTP server
	server := &http.Server{
		Addr: ":8443", // Define the address and port for HTTPS
	}

	// Configure the server to use HTTP/2
	http2Server := &http2.Server{}
	http2.ConfigureServer(server, http2Server)

	// Define a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTP/2 with TLS!"))
	})

	// Define a streaming handler
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		// Ensure the writer supports flushing
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// Set the headers for streaming
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Stream data
		for i := 0; i < 10; i++ {
			// Write data to the client
			fmt.Fprintf(w, "data: Message %d\n\n", i)
			// Flush the data immediately
			flusher.Flush()
			// Simulate a delay
			time.Sleep(1 * time.Second)
		}
	})

	// Start the server with TLS
	log.Println("Starting HTTP/2 server with TLS on :8443")
	if err := server.ListenAndServeTLS("./server.crt", "./server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

