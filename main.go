package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World v1b")
	 
	

}

func main() {
	port := ":8070"

	// Print a message indicating the server is starting
	fmt.Printf("Starting server on port%s\n", port)
	http.HandleFunc("/", handler)

	http.ListenAndServe(port, nil)
	fmt.Printf("Appication stop\n")
}
