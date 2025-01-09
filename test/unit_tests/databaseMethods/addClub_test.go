package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddClub(t *testing.T) {
	d := mt.Database{
		Clubs:   map[int]*mt.Club{},
		Teams:   map[int]*mt.Team{},
		Players: map[int]*mt.Player{},
	}

	c1 := mt.Club{
		ID:   0,
		Name: "c1",
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	expectedLen := 2

	t.Run("Add club to database", func(t *testing.T) {
		d.AddClub(&c1)
		d.AddClub(&c2)
		lenToVerify := len(d.Clubs)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Clubs %v, got %v", expectedLen, lenToVerify)
		}
	})
}
