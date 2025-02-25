package api

import (
	"DAF-Core/app/model/dto"
	"DAF-Core/app/repository"
	"DAF-Core/app/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var itemService service.ItemService

func GetAllItemsByBoard(w http.ResponseWriter, r *http.Request) {
	// Extract board UUID from URL path parameters
	vars := mux.Vars(r)
	boardUUID, ok := vars["board_uuid"]
	if !ok || boardUUID == "" {
		http.Error(w, `{"error": "Missing board UUID"}`, http.StatusBadRequest)
		return
	}

	// Get items from repository
	data, err := itemService.GetAllByBoard(boardUUID)
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve items"}`, http.StatusInternalServerError)
		return
	}

	// Set response headers and encode response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemUUID, ok := vars["item_uuid"]
	if !ok || itemUUID == "" {
		http.Error(w, `{"error": "Missing item UUID"}`, http.StatusBadRequest)
		return
	}

	data, err := itemService.Get(itemUUID)
	if err != nil {
		http.Error(w, `{"error": "Item not found"}`, http.StatusNotFound)
		return
	}

	if data == nil {
		http.Error(w, `{"error": "Item not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var itemDTO dto.CreateItem

	json.NewDecoder(r.Body).Decode(&itemDTO)

	item, err := itemService.Create(itemDTO)
	if err != nil {
		log.Printf("Failed to create item in repository: %v", err)
		http.Error(w, `{"error": "Creation failed"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
