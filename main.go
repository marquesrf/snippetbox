package main

import (
	"log"
	"net/http"
)

// Handler function, works like a controller.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// View snippet handler.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Create snippet handler.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Create and initialize a web server.
	mux := http.NewServeMux()

	// Register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Start the web server on a given port.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	// If the web server returns an error, we log it.
	log.Fatal(err)
}
