package main

import (
	"github.com/mongo-Api/router"
	"log"
	"net/http"
)

func main() {

	r := router.CreateRouter()

	// Start the server on port 8080
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
