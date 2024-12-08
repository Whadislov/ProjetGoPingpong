package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	d := mt.Database{
		Clubs:   map[int]*mt.Club{},
		Teams:   map[int]*mt.Team{},
		Players: map[int]*mt.Player{},
	}

	p1 := mt.Player{
		Name: "p1",
	}

	p2 := mt.Player{
		Name: "p2",
	}

	expectedLen := 2
	expectedID1 := 0
	expectedID2 := 1

	t.Run("Add player to database", func(t *testing.T) {
		d.AddPlayer(&p1)
		d.AddPlayer(&p2)
		lenToVerify := len(d.Players)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Players %v, got %v", expectedLen, lenToVerify)
		}
		if p1.ID != expectedID1 {
			t.Errorf("Expected ID %v, got %v", expectedID1, p1.ID)
		}
		if p2.ID != expectedID2 {
			t.Errorf("Expected ID %v, got %v", expectedID2, p2.ID)
		}
	})
}
