package main

import (
	"log"
	"net/http"

	"github.com/kics223w1/EndpointLab2/api"
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
	http.HandleFunc("/stream", api.StreamHandler)

	// Define an echo handler
	http.HandleFunc("/anything", api.EchoHandler)

	// Start the server with TLS
	log.Println("Starting HTTP/2 server with TLS on :8443")
	if err := server.ListenAndServeTLS("./server.crt", "./server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

