package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// LoadDB loads the database.
func LoadDB(authToken string) (*mt.Database, error) {
	var golangDB *mt.Database

	req, err := http.NewRequest("POST", "http://localhost:8001/api/load-database", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch database: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&golangDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	log.Println("Database loaded successfully")
	return golangDB, nil
}

// LoadDB loads the users.
func LoadUsersOnly() (*mt.Database, error) {
	var golangDB *mt.Database

	resp, err := http.Get("http://localhost:8001/api/load-users")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch database: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&golangDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	log.Println("Database loaded successfully")
	return golangDB, nil
}

// SavedDB saves the database.
func SaveDB(golangDB *mt.Database) error {

	dataToSave, err := json.Marshal(golangDB)
	if err != nil {
		return fmt.Errorf("failed to marshal database: %w", err)
	}

	resp, err := http.Post("http://localhost:8001/api/save-database", "application/json", bytes.NewBuffer(dataToSave))
	if err != nil {
		return fmt.Errorf("failed to sent request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server error: %s", resp.Status)
	}
	log.Println("Database saved successfully")
	return nil
}

// IsApiReady posts a message if the API is started
func IsApiReady() bool {
	resp, err := http.Get("http://localhost:8001/")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return true
}

// Login requests a credentials check to the API, if everything is fine, the database of the user is returned
func Login(username string, password string) (*mt.Database, error) {
	var token string
	var db *mt.Database
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	creds.Username = username
	creds.Password = password

	credentialsToCheck, err := json.Marshal(creds)
	if err != nil {
		return db, fmt.Errorf("failed to marshal credentials: %w", err)
	}

	resp, err := http.Post("http://localhost:8001/api/login", "application/json", bytes.NewBuffer(credentialsToCheck))
	if err != nil {
		return db, fmt.Errorf("failed to post credentials: %w", err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return db, fmt.Errorf("error decoding JSON: %w", err)
	} else {
		log.Println("Succeed to log user %w in", username)
		db, err := LoadDB(token)
		if err != nil {
			return db, fmt.Errorf("failed to load database: %w", err)
		}

		return db, nil
	}
}

// Login requests a new user creation, if everything is fine, the database of the new user is returned
func SignUp(username string, password string, email string) (*mt.Database, error) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	var token string
	var db *mt.Database
	data.Username = username
	data.Password = password
	data.Email = email

	jsonData, err := json.Marshal(data)
	if err != nil {
		return db, err
	}

	resp, err := http.Post("http://localhost:8001/api/signup", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return db, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return db, err
	} else {
		log.Println("Succeed to sign user %w in", username)
		db, err := LoadDB(token)
		if err != nil {
			return db, fmt.Errorf("failed to load database: %w", err)
		}
		return db, nil
	}
}
