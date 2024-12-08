package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeleteClub(t *testing.T) {
	d := mt.Database{
		Clubs: map[int]*mt.Club{0: {
			ID:   0,
			Name: "c",
		},
		},
		Teams:   map[int]*mt.Team{},
		Players: map[int]*mt.Player{},
	}

	expectedLen := 0
	expectedError := "clubID 1 does not exist"

	t.Run("Delete club from database", func(t *testing.T) {
		err := d.DeleteClub(0)
		err2 := d.DeleteClub(1)
		lenToVerify := len(d.Clubs)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Clubs %v, got %v", expectedLen, lenToVerify)
		}
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err2)
		}
	})
}
