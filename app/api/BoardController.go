package api

import (
	"DAF-Core/app/model/dto"
	"DAF-Core/app/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var boardService service.BoardService

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	// Get boards from repository
	data, err := boardService.GetAll()
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve boards"}`, http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Encode response
	json.NewEncoder(w).Encode(data)
}
func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardUUID, ok := vars["board_uuid"]
	if !ok || boardUUID == "" {
		http.Error(w, `{"error": "Missing board UUID"}`, http.StatusBadRequest)
		return
	}

	err := boardService.Delete(boardUUID)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete board"}`, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func CreateBoard(w http.ResponseWriter, r *http.Request) {

	var boardDTO dto.CreateBoard

	json.NewDecoder(r.Body).Decode(&boardDTO)

	_, err := boardService.Create(boardDTO)
	if err != nil {
		http.Error(w, `{"error": "Creation failed"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
