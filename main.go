package main

import (
	"log"
	"net/http"
)

// Handler function, works like a controller.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Create and initialize a web server.
	mux := http.NewServeMux()

	// Register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	// Start the web server on a given port.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	// If the web server returns an error, we log it.
	log.Fatal(err)
}
