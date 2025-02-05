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

	req, err := http.NewRequest("GET", "http://localhost:8001/api/load-database", nil)
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
