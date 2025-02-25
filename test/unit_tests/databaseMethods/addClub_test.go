package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestAddClub(t *testing.T) {
	d := mt.Database{
		Clubs:   map[uuid.UUID]*mt.Club{},
		Teams:   map[uuid.UUID]*mt.Team{},
		Players: map[uuid.UUID]*mt.Player{},
	}

	c1 := mt.Club{
		ID:   uuid.New(),
		Name: "c1",
	}

	c2 := mt.Club{
		ID:   uuid.New(),
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
