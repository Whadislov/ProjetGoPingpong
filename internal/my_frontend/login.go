package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// Login requests a credentials check to the API, if everything is fine, the database of the user is returned
func Login(username string, password string) (*mt.Database, string, error) {
	var response struct {
		Token string `json:"token"`
	}

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

	resp, err := http.Post("http://localhost:8001/api/login", "application/json", bytes.NewBuffer(credentialsToCheck))
	if err != nil {
		return nil, "", fmt.Errorf("failed to post credentials: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse struct {
			Error string `json:"error"`
			Code  string `json:"code"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, "", fmt.Errorf("error decoding server response: %w", err)
		}

		if errorResponse.Code == "INVALID_USERNAME_OR_PASSWORD" {
			return nil, "", fmt.Errorf("username or password is invalid")
		} else if errorResponse.Code == "INVALID_REQUEST" {
			return nil, "", fmt.Errorf("invalid request")
		} else {
			return nil, "", fmt.Errorf("internal error: %s", errorResponse.Error)
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	defer resp.Body.Close()
	if err != nil {
		return nil, "", fmt.Errorf("error decoding JSON: %w", err)
	} else {
		log.Println("Succeed to log user %w in", username)
		db, err := LoadDB(response.Token)
		if err != nil {
			return nil, "", fmt.Errorf("failed to load database: %w", err)
		}

		return db, response.Token, nil
	}
}
