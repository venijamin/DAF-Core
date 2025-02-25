package api

import (
	"DAF-Core/app/repository"
	"encoding/json"
	"log"
	"net/http"
)

var boardRepository repository.BoardRepository

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	// Get boards from repository
	data, err := boardRepository.GetAll()
	if err != nil {
		log.Printf("Failed to retrieve boards: %v", err)
		http.Error(w, `{"error": "Failed to retrieve boards"}`, http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Encode response
	json.NewEncoder(w).Encode(data)
}
