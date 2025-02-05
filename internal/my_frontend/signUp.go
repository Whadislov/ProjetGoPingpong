package myfrontend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

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
