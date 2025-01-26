package my_functions

import (
	"fmt"
	"log"
	"time"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// NewUser creates a new user with the given name and adds it to the database.
// Returns the created user and an error if the user name is empty or if there is an issue with the operation.
func NewUser(username string, email string, passwordHash string, db *mt.Database) (*mt.User, error) {
	b, err := isValidUsername(username)
	if !b {
		return nil, err
	}

	b, err = isValidEmail(email)
	if !b {
		return nil, err
	}

	for _, user := range db.Users {
		if user.Name == username {
			return nil, fmt.Errorf("username is already taken")
		} else if user.Email == email {
			return nil, fmt.Errorf("email is already used")
		}
	}

	// ISO 8601 timestamp
	timestamp := time.Now().Format(time.RFC3339)

	u := &mt.User{
		ID:           len(db.Users),
		Name:         username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    timestamp,
	}

	db.AddUser(u)
	log.Printf("User %v sucessfully created.", username)
	return u, nil
}
