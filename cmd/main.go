package main

import (
	"log"
	"net/http"
	"website-monitoring/internal/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
