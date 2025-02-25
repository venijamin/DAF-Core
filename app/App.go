package main

import (
	"DAF-Core/app/api"
	"DAF-Core/app/util"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	util.ConnectToDB()         // Establish database connection
	router := mux.NewRouter()  // Create a router
	SetRoutes(router)          // Set the HTTP routes
	corsRouter := CORS(router) // Wrap the router with CORS middleware

	InitData()
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsRouter))

}

// CORS middleware function
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetRoutes(router *mux.Router) {
	// Create a new router
	router.PathPrefix("/src/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			// Set the correct MIME type for CSS files
			w.Header().Set("Content-Type", "text/css")

			// Serve the CSS file
			http.StripPrefix("/src/", http.FileServer(http.Dir("src"))).ServeHTTP(w, r)
		} else {
			http.NotFound(w, r) // Return 404 for non-CSS files
		}
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("src/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})
	router.HandleFunc("/api/boards", api.GetAllBoards).Methods("GET")
	router.HandleFunc("/api/boards/{board_uuid}", api.GetAllItemsByBoard).Methods("GET")
	router.HandleFunc("/api/items/{item_uuid}", api.GetItem).Methods("GET")
	router.HandleFunc("/api/items/", api.CreateItem).Methods("POST")
}
