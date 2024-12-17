package main

import (
	"github.com/benjamint08/go-server-template/handlers"
	"net/http"
)

func main() {
	// Define API routes here
	http.HandleFunc("/api/hello", handlers.HelloHandler)
	// End API routes

	http.ListenAndServe(":8080", nil)
}
