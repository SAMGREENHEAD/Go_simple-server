package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler is a handler function for the /form endpoint. It handles
// POST requests and parses the form data from the request. It then
// writes a success message and the parsed data back to the
// response writer.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request.
	if err := r.ParseForm(); err != nil {
		// If there was an error parsing the form data, write an
		// error message to the response writer.
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Write a success message to the response writer.
	fmt.Fprintf(w, "POST request successful")

	// Get the parsed form values from the request.
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Write the parsed form values back to the response writer.
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)
}

// helloHandler is a handler function for the /hello endpoint. It
// checks the URL and the HTTP method of the request to ensure that
// it is a valid request. If the request is valid, it writes a
// success message to the response writer.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check the URL of the request.
	if r.URL.Path != "/hello" {
		// If the URL is not /hello, write a 404 error to the response writer.
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check the HTTP method of the request.
	if r.Method != "GET" {
		// If the HTTP method is not GET, write a 404 error to the response writer.
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// Write a success message to the response writer.
	fmt.Fprintf(w, "hello!")
}

// main is the entry point of the program. It sets up the
// http server and its routes, and then starts the server.
func main() {
	// Set up the http server to look at the static folder.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Set up the route for the /form endpoint.
	http.HandleFunc("/form", formHandler)

	// Set up the route for the /hello endpoint.
	http.HandleFunc("/hello", helloHandler)

	// Start the server.
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Log any errors that occurred while starting the server.
		log.Fatal(err)
	}
}
