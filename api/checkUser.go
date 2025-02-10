package api

import (
	"fmt"
	"log"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"golang.org/x/crypto/bcrypt"
)

// checkUserCredentials verifies if the credentials of the user are correct, returns the userID
func checkUserCredentials(username string, password string) (int, error) {
	log.Println("Loading DB to check user credentials")
	db, err := mdb.LoadUsersOnly()
	if err != nil {
		return -1, fmt.Errorf("error during connexion to database to check user credentials")
	}

	for _, user := range db.Users {
		if username == user.Name {

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return -1, fmt.Errorf("failed to hash password: %v", err)
			}

			if string(hashedPassword) == user.PasswordHash {
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
func checkUserExists(username string, email string, db *mt.Database) (int, error) {
	log.Println("Loading DB to check user existence before creation")

	// Check Email first as email is the first field in the UI
	for _, user := range db.Users {
		if user.Email == email {
			return 1, fmt.Errorf("email already exists")
		}
		if user.Name == username {
			return 2, fmt.Errorf("username already exists")
		}

	}
	return 0, nil
}
