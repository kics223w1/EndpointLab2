package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// Create a new HTTP server
	server := &http.Server{
		Addr: ":8080", // Define the address and port
	}

	// Configure the server to use HTTP/2
	http2.ConfigureServer(server, &http2.Server{})

	// Define a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, HTTP/2!"))
	})

	// Start the server
	log.Println("Starting HTTP/2 server on :8080")
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

