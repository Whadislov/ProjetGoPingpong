package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddTeam(t *testing.T) {
	d := mt.Database{
		Clubs:   map[int]*mt.Club{},
		Teams:   map[int]*mt.Team{},
		Players: map[int]*mt.Player{},
	}

	t1 := mt.Team{
		Name: "t1",
	}

	t2 := mt.Team{
		Name: "t2",
	}

	expectedLen := 2
	expectedID1 := 0
	expectedID2 := 1

	t.Run("Add team to database", func(t *testing.T) {
		d.AddTeam(&t1)
		d.AddTeam(&t2)
		lenToVerify := len(d.Teams)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Teams %v, got %v", expectedLen, lenToVerify)
		}
		if t1.ID != expectedID1 {
			t.Errorf("Expected ID %v, got %v", expectedID1, t1.ID)
		}
		if t2.ID != expectedID2 {
			t.Errorf("Expected ID %v, got %v", expectedID2, t2.ID)
		}
	})
}
