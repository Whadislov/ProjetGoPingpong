package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeleteTeam(t *testing.T) {
	d := mt.Database{
		Clubs: map[int]*mt.Club{},
		Teams: map[int]*mt.Team{0: {
			ID:   0,
			Name: "c",
		}},
		Players: map[int]*mt.Player{},
	}

	expectedLen := 0
	expectedError := "teamID 1 does not exist"

	t.Run("Delete team from database", func(t *testing.T) {
		err := d.DeleteTeam(0)
		err2 := d.DeleteTeam(1)
		lenToVerify := len(d.Teams)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Players %v, got %v", expectedLen, lenToVerify)
		}
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err2)
		}
	})
}
