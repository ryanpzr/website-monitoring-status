package main

import (
	"log"
	"net/http"
	"website-monitoring/internal/router"
	"website-monitoring/internal/service"
)

func main() {
	r := router.SetupRouter()
	go service.VerifyWebStatus()
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
