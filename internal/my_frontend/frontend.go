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
func LoadDB() (*mt.Database, error) {
	var golangDB *mt.Database

	resp, err := http.Get("http://localhost:8001/api/load-database")
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

// Checks if the API is launched
func AuthenticateUser() (bool, error) {
	var authResult string
	resp, err := http.Get("http://localhost:8001/api/authenticate-user")
	if err != nil {
		return false, fmt.Errorf("failed to fetch user authentification: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&authResult)
	if err != nil {
		return false, fmt.Errorf("error decoding JSON: %w", err)
	}

	if authResult == "Authentification is successfull" {
		return true, nil
	} else if authResult == "Authentification is unsuccessfull" {
		return false, nil
	} else {
		return false, fmt.Errorf("authentification result is unknown: %w", err)
	}
}
