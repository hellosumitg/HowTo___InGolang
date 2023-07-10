package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// ApiServer handles incoming HTTP requests and routes them to the appropriate handlers.
type ApiServer struct {
	svc Service
}

// NewApiServer creates a new instance of ApiServer with the provided Service.
func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

// Start starts the HTTP server and listens for incoming requests on the specified address.
func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

// handleGetCatFact is the HTTP handler function for retrieving a cat fact.
func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	// Call the GetCatFact method of the underlying service to retrieve a cat fact.
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]interface{}{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

// writeJSON writes the provided data as JSON response with the specified status code.
func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
