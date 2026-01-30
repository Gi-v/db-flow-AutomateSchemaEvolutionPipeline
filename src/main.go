package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/api"
)

func main() {
	fmt.Println("Mini Relational DBMS - Starting...")
	
	// Initialize API server
	router := api.NewRouter()
	
	port := ":8080"
	fmt.Printf("Server listening on %s\n", port)
	
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
