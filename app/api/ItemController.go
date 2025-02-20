package api

import (
	"DAF-Core/app/model/dto"
	"DAF-Core/app/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var itemRepository repository.ItemRepository

func GetAllItemsByBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardUUID, _ := vars["board_uuid"]
	_ = json.NewDecoder(r.Body).Decode(&boardUUID)
	data := itemRepository.GetAllByBoard(boardUUID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	itemCreationDTO := dto.CreateItem{
		ParentUUID:  "",
		BoardUUID:   "",
		Name:        "",
		Description: "",
		Quantity:    0,
		Tags:        nil,
		Picture:     "",
		Barcode:     "",
		Fields:      nil,
	}
}
