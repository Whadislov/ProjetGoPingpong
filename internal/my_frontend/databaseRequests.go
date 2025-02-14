package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
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

// SaveDB saves the database.
func SaveDB(authToken string, golangDB *mt.Database) error {

	dataToSave, err := json.Marshal(golangDB)
	if err != nil {
		return fmt.Errorf("failed to marshal database: %w", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8001/api/save-database", bytes.NewBuffer(dataToSave))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add the token in the Authorization header
	log.Println("authToken=", authToken)
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-OK status: %d", resp.StatusCode)
	}

	log.Println("Database saved successfully")
	return nil
}
