package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CI/CD Go backend WORKS 🚀")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: OK")
}

func newEndpoint(w http.ResponseWriter, r *http.Request) {
	// New functionality can be added here
	fmt.Fprintln(w, "New Endpoint Reached")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/status",statusHandler )
	http.HandleFunc("/new", newEndpoint)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
