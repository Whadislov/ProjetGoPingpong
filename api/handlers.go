package api

import (
	"encoding/json"
	"log"
	"net/http"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Handler for loading the database
func loadDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to load database")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db, err := mdb.LoadDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Println("Error connecting to database:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}

// Handler for saving the local changes to the database
func saveDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to save database")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var db *mt.Database
	err := json.NewDecoder(r.Body).Decode(&db)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		log.Println("Error decoding JSON:", err)
		return
	}
	err = mdb.SaveDB(db)
	if err != nil {
		http.Error(w, "Failed to save database", http.StatusBadRequest)
		log.Println("Error saving database:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database saved successfully"))
}

// Handler for saving the local changes to the database
func isApiReady(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to check if API is ready")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is ready !"))
}
