package api

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var boardService service.BoardService

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	data, err := boardService.GetAll()
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve boards"}`, http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("Content-Type", "text/html")
		renderBoardList(w, *data)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func renderBoardList(w http.ResponseWriter, boards []model.Board) {
	w.Write([]byte("<ul>"))
	for _, board := range boards {
		w.Write([]byte(fmt.Sprintf(
			`<li id="board-%s">
                %s
                <button hx-delete="/boards/%s" hx-target="#board-%s" hx-swap="outerHTML">Delete</button>
            </li>`,
			board.BoardUUID, board.Name, board.BoardUUID, board.BoardUUID)))
	}
	w.Write([]byte("</ul>"))
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBoard function called")
	vars := mux.Vars(r)
	boardUUID, ok := vars["board_uuid"]
	if !ok || boardUUID == "" {
		log.Println("Missing board UUID")
		http.Error(w, `{"error": "Missing board UUID"}`, http.StatusBadRequest)
		return
	}

	log.Printf("Attempting to delete board with UUID: %s", boardUUID)
	err := boardService.Delete(boardUUID)
	if err != nil {
		log.Printf("Failed to delete board: %v", err)
		http.Error(w, `{"error": "Failed to delete board"}`, http.StatusInternalServerError)
		return
	}

	log.Println("Board deleted successfully")
	if r.Header.Get("HX-Request") == "true" {
		w.WriteHeader(http.StatusOK)
		// Return empty response to remove the element
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Board deleted successfully"}`))
	}
}
func CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boardDTO dto.CreateBoard
	err := json.NewDecoder(r.Body).Decode(&boardDTO)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	boardUUID, err := boardService.Create(boardDTO)
	if err != nil {
		http.Error(w, `{"error": "Creation failed"}`, http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprintf(
			`<li id="board-%s">
                %s
                <button hx-delete="/boards/%s" hx-target="#board-%s" hx-swap="outerHTML">Delete</button>
            </li>`,
			boardUUID, boardUUID, boardUUID, boardUUID, boardUUID)))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(boardUUID)
	}
	w.WriteHeader(http.StatusCreated)
}
