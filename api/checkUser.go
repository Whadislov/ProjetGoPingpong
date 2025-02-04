package api

import (
	"fmt"
	"log"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// checkUserCredentials verifies if the credentials of the user are correct, returns the userID
func checkUserCredentials(username string, password string) (int, error) {
	db, err := mdb.LoadDB()
	if err != nil {
		return -1, fmt.Errorf("error during connexion to database to check user credentials")
	}

	for _, user := range db.Users {
		if username == user.Name {
			if password == user.PasswordHash {
				log.Println("User credentials are good")
				mdb.SetUserIDOfSession(user.ID)
				return user.ID, nil
			}
		}
	}
	// Wrong credentials or user does not exist
	return -1, fmt.Errorf("username and password missmatch")
}

// checkUserCredentials verifies that the newly created user does not use an already registered username
func checkUserExists(username string, email string, db *mt.Database) (bool, error) {

	for _, user := range db.Users {
		if user.Name == username {
			return true, fmt.Errorf("username already exists")
		}
		if user.Email == email {
			return true, fmt.Errorf("email already exists")
		}
	}
	return false, nil
}
