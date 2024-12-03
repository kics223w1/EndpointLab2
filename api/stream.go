package api

import (
	"fmt"
	"net/http"
	"time"
)

func StreamHandler(w http.ResponseWriter, r *http.Request) {
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
}
