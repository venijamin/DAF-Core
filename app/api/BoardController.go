package api

import (
	"DAF-Core/app/model"
	"DAF-Core/app/model/dto"
	"DAF-Core/app/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var boardService service.BoardService

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	var boardListTemplate *template.Template
	boardListTemplate = template.Must(template.ParseFiles("app/src/template/board-list.html"))

	data, err := boardService.GetAll()
	if err != nil {
		http.Error(w, `{"error": "Failed to retrieve boards"}`, http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("Content-Type", "text/html")
		err = boardListTemplate.ExecuteTemplate(w, "boardList", data)
		if err != nil {
			http.Error(w, "Failed to renderer template", http.StatusInternalServerError)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func renderBoardList(w http.ResponseWriter, boards []model.Board) {
	w.Write([]byte("<ul>"))
	for _, board := range boards {
		w.Write([]byte(fmt.Sprintf(
			`<a href=/api/boards/%s>
				<li id="board-%s">
                %s
                <button hx-delete="/api/boards/%s" hx-target="#board-%s" hx-swap="outerHTML">Delete</button>
            	</li>
				</a>`,
			board.BoardUUID, board.BoardUUID, board.Name, board.BoardUUID, board.BoardUUID)))
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if r.Header.Get("HX-Request") == "true" {
		// Return an empty JSON object for HTMX requests
		w.Write([]byte(`{}`))
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "Board deleted successfully"})
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(boardUUID)
}
