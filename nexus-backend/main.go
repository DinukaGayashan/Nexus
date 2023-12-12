package main

import (
	"encoding/json"
	"net/http"
)

// Response struct represents the structure of the response message.
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Handle requests to the "/api" endpoint with the helloHandler function.
	http.HandleFunc("/api", helloHandler)

	// Start the HTTP server on port 8080.
	http.ListenAndServe(":8080", nil)
}

// helloHandler is a simple HTTP handler function.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Create a Response struct with a message.
	response := Response{
		Message: "Hello, this is your API response!",
	}

	// Convert the Response struct to JSON.
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client.
	w.Write(responseJSON)
}
