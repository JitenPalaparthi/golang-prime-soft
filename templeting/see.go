package main

import (
	"fmt"
	"net/http"
	"time"
)

// Event structure to define an SSE event
type Event struct {
	Data string `json:"data"`
}

// SSE Handler function to stream events to the client
func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Set the event stream to flush immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Send a message to the client every second
	for {
		event := Event{
			Data: fmt.Sprintf("Current Time: %s", time.Now().Format(time.RFC3339)),
		}

		// Write the event to the response writer in SSE format
		fmt.Fprintf(w, "data: %s\n\n", event.Data)

		// Flush the response to the client
		flusher.Flush()

		// Wait for 1 second before sending the next event
		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	// Start the HTTP server
	fmt.Println("Server started at http://localhost:8087")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
