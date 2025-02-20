package api

import (
	"DAF-Core/app/repository"
	"encoding/json"
	"net/http"
)

var boardRepository repository.BoardRepository

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	data := boardRepository.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
