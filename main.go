package main

import (
	"fmt"
	"github.com/benjamint08/go-server-template/handlers"
	"net/http"
)

func main() {
	// Define API routes here
	http.HandleFunc("/api/hello", handlers.HelloHandler)
	// End API routes

	fmt.Println("api server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
