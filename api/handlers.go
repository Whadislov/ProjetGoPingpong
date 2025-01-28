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

// Handler for loading the users
func loadUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to load users")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db, err := mdb.LoadUsersOnly()
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

// Handler for authenticating the user
func authenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to authenticate")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user *mt.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		log.Println("Error decoding JSON:", err)
		return
	}

	db, err := mdb.LoadUsersOnly()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Println("Error connecting to database:", err)
		return
	}

	for i, u := range db.Users {
		if u.Name != user.Name && i == len(db.Users) {
			log.Println("Wrong username or password", err)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Authentification is unsuccessfull"))
		} else if u.Name == user.Name {
			if u.PasswordHash != user.PasswordHash {
				log.Println("Wrong username or password", err)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Authentification is unsuccessfull"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Authentification is successfull"))
			}
		}
	}
}
