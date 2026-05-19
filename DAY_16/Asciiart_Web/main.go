package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Create a new ServeMux here
	mux := http.NewServeMux()

	// Define routes here
	mux.HandleFunc("/profile", handleProfile)
	mux.HandleFunc("/blog", handleBlogs)
	mux.HandleFunc("/", handleRoot)

	// Below is the server setup
	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root!")
}

func handleBlogs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Blog Page!")
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Profile Page")
}
