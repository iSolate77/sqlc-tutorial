package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type controller struct {
	s Service
}

func newController(s Service) *controller {
	return &controller{s: s}
}

func (c *controller) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var auth Author
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.s.CreateAuthor(ctx, auth); err != nil {
		log.Printf("Error creating author: %v", err) // Logging the error
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(auth)
}
