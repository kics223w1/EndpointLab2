package api

import (
	"io"
	"net/http"
)

// EchoHandler returns anything that is passed in the request body
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to indicate HTTP/2
	w.Header().Set("Content-Type", "application/octet-stream")

	// Copy the request body to the response writer
	_, err := io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
}

