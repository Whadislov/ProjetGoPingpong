package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Login requests a credentials check to the API, if everything is fine, the database of the user is returned
func Login(username string, password string) (*mt.Database, string, error) {
	var token string
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	creds.Username = username
	creds.Password = password

	credentialsToCheck, err := json.Marshal(creds)
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal credentials: %w", err)
	}
	log.Println("Before http post")

	resp, err := http.Post("http://localhost:8001/api/login", "application/json", bytes.NewBuffer(credentialsToCheck))
	if err != nil {
		return nil, "", fmt.Errorf("failed to post credentials: %w", err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding JSON: %w", err)
	} else {
		log.Println("Succeed to log user %w in", username)
		db, err := LoadDB(token)
		if err != nil {
			return nil, "", fmt.Errorf("failed to load database: %w", err)
		}

		return db, token, nil
	}
}
