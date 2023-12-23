package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/usermp/go-api.git/handlers"
)

func main() {
	// Create a new router
	r := mux.NewRouter()
	// Define routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.GetUserHandler).Methods("GET")

	// Start the server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
