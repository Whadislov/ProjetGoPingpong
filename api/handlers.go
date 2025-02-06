package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Handler for loading the database
func loadUserDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to load user DB")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	/*
		var userID int
		err := json.NewDecoder(r.Body).Decode(&userID)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			log.Println("Error decoding JSON:", err)
			return
		}
	*/

	userID := r.Header.Get("User-ID")
	if userID == "" {
		http.Error(w, "User is unidentified", http.StatusUnauthorized)
		return
	}

	// Convert the str ID into a int ID
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	mdb.SetUserIDOfSession(id)
	db, err := mdb.LoadDB()
	if err != nil {
		http.Error(w, "Failed to connect to database.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

// Handler for saving the local changes to the database
func saveDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var db *mt.Database
	err := json.NewDecoder(r.Body).Decode(&db)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	err = mdb.SaveDB(db)
	if err != nil {
		http.Error(w, "Failed to save database", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database saved successfully"))
}

// Handler to check if the API is ready to take requests (not yet used)
func isApiReady(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API is ready !"))
}

// loginHandler process the request to analyse the credentials, returns a token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	userID, err := checkUserCredentials(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(userID)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}

// signUpHandler process the request to create a new user, returns a token
func signUpHandler(w http.ResponseWriter, r *http.Request) {
	var signUpData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&signUpData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	/*
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Could not hash password", http.StatusInternalServerError)
			return
		}
	*/

	// Load the whole database to register the new user. Could be optimised to request directly postgres
	db, err := mdb.LoadDB()
	if err != nil {
		http.Error(w, "Could not load database to create the new user", http.StatusInternalServerError)
		return
	}

	// Verify if the user already exists
	exists, _ := checkUserExists(signUpData.Username, signUpData.Email, db)
	if exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Save the user in the database, need to enter x2 password (UI design for the local app)
	newUser, err := mf.NewUser(signUpData.Username, signUpData.Email, signUpData.Password, signUpData.Password, db)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	// User Id is the number of current registered users (there is possibility yet to delete a user, so this should work for now)
	token, err := generateJWT(newUser.ID)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}
