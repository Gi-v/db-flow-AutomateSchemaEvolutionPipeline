package api

import (
	"encoding/json"
	"net/http"
)

// NewRouter creates and configures the HTTP router
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	
	// Health check endpoint
	mux.HandleFunc("/health", healthHandler)
	
	// API endpoints (to be implemented)
	mux.HandleFunc("/api/query", queryHandler)
	
	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"service": "mini-dbms",
	})
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Query endpoint - to be implemented",
	})
}
