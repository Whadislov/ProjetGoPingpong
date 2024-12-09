package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestNewClub(t *testing.T) {
	d := mt.Database{
		Clubs: map[int]*mt.Club{},
	}

	expectedLen := 1
	expectedError1 := "Club name cannot be empty"

	t.Run("Delete Club", func(t *testing.T) {
		_, err1 := mf.NewClub("", &d)
		_, err2 := mf.NewClub("p2", &d)

		if err1 == nil {
			t.Errorf("Expected error %v, got %v", expectedError1, err1)
		}
		if err2 != nil || len(d.Clubs) != expectedLen {
			t.Errorf("Expected error %v, got %v", nil, err2)
		}
	})
}
