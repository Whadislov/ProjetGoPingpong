package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestNewUser(t *testing.T) {
	d := mt.Database{
		Users: map[int]*mt.User{0: {
			ID:           0,
			Name:         "u1",
			Email:        "a1@a1.com",
			PasswordHash: "b",
			CreatedAt:    "c",
		}},
	}

	expectedLen := 2
	expectedError1 := "username cannot be empty"
	expectedError2 := "username is already taken"
	expectedError3 := "email cannot be empty"
	expectedError4 := "email is already used"

	t.Run("New user", func(t *testing.T) {
		passwordHash := "b"
		_, err1 := mf.NewUser("", "", passwordHash, &d)
		_, err2 := mf.NewUser("u1", "", passwordHash, &d)
		_, err3 := mf.NewUser("u2", "", passwordHash, &d)
		_, err4 := mf.NewUser("u2", "a1@a1.com", passwordHash, &d)
		_, err5 := mf.NewUser("u2", "a2@a2.com", passwordHash, &d)

		if err1 == nil {
			t.Errorf("Expected error %v, got %v", expectedError1, err1)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err2)
		}
		if err4 == nil {
			t.Errorf("Expected error %v, got %v", expectedError4, err2)
		}
		if err5 != nil || len(d.Users) != expectedLen {
			t.Errorf("Expected error %v, got %v", nil, err5)
			t.Errorf("Expected length %v, got %v", 2, len(d.Users))
		}
	})
}
