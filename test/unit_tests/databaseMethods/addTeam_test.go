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
		ID:   0,
		Name: "t1",
	}

	t2 := mt.Team{
		ID:   1,
		Name: "t2",
	}

	expectedLen := 2

	t.Run("Add team to database", func(t *testing.T) {
		d.AddTeam(&t1)
		d.AddTeam(&t2)
		lenToVerify := len(d.Teams)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Teams %v, got %v", expectedLen, lenToVerify)
		}
	})
}
