package my_functions

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// ChangeUsername changes the current username with a new one.
func ChangeUsername(currentUsername string, newUsername string, db *mt.Database) error {
	b, err := IsValidUsername(newUsername)
	if !b {
		return err
	}
	log.Println("User modification : New username is valid.")

	var u *mt.User
	for _, user := range db.Users {
		if user.Name == newUsername {
			return fmt.Errorf("username is already taken")
		} else if user.Name == currentUsername {
			u = user
		}
	}

	u.Name = newUsername
	log.Printf("User modification : Username (%v -> %v) sucessfully modified.", currentUsername, newUsername)
	return nil
}

// ChangePassword changes the current password with a new one.
func ChangePassword(newPassword string, user *mt.User) error {
	b, err := IsValidPassword(newPassword)
	if !b {
		return err
	}
	log.Println("User modification : New password is valid.")

	user.PasswordHash = newPassword
	log.Printf("User modification : User %v's password sucessfully modified.", user.Name)
	return nil
}

// ChangeEmail changes the current email with a new one.
func ChangeEmail(newEmail string, user *mt.User) error {
	b, err := IsValidEmail(newEmail)
	if !b {
		return err
	}
	log.Println("User modification : New email is valid.")

	user.Email = newEmail
	log.Printf("User modification : User %v's email sucessfully modified.", user.Name)
	return nil
}

// DeleteUser deletes the user from the database.
func DeleteUser(user *mt.User, db *mt.Database) error {
	return db.DeleteUser(user.ID)
}
