package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestAddTeam(t *testing.T) {
	d := mt.Database{
		Clubs:   map[uuid.UUID]*mt.Club{},
		Teams:   map[uuid.UUID]*mt.Team{},
		Players: map[uuid.UUID]*mt.Player{},
	}

	t1 := mt.Team{
		ID:   uuid.New(),
		Name: "t1",
	}

	t2 := mt.Team{
		ID:   uuid.New(),
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
