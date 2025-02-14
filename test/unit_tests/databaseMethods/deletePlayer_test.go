package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestDeletePlayer(t *testing.T) {
	d := mt.Database{
		Clubs: map[int]*mt.Club{},
		Teams: map[int]*mt.Team{},
		Players: map[int]*mt.Player{0: {
			ID:        0,
			Firstname: "c",
		}},
	}

	expectedLen := 0
	expectedError := "playerID 1 does not exist"

	t.Run("Delete player from database", func(t *testing.T) {
		err := d.DeletePlayer(0)
		err2 := d.DeletePlayer(1)
		lenToVerify := len(d.Players)
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
