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

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/status",statusHandler )
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
