package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	d := mt.Database{
		Clubs:   map[uuid.UUID]*mt.Club{},
		Teams:   map[uuid.UUID]*mt.Team{},
		Players: map[uuid.UUID]*mt.Player{},
	}

	p1 := mt.Player{
		ID:        uuid.New(),
		Firstname: "p1",
	}

	p2 := mt.Player{
		ID:        uuid.New(),
		Firstname: "p2",
	}

	expectedLen := 2

	t.Run("Add player to database", func(t *testing.T) {
		d.AddPlayer(&p1)
		d.AddPlayer(&p2)
		lenToVerify := len(d.Players)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Players %v, got %v", expectedLen, lenToVerify)
		}
	})
}
