package playermethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestAddTeam(t *testing.T) {
	t1 := mt.Team{
		ID:   uuid.New(),
		Name: "t1",
	}

	t2 := mt.Team{
		ID:   uuid.New(),
		Name: "t2",
	}

	p1 := mt.Player{
		ID:        uuid.New(),
		Firstname: "p1",
		TeamIDs:   map[uuid.UUID]string{t2.ID: "t1"},
	}

	expectedLen := 2
	expectedError := "player p1 is already in team t2"

	t.Run("Add team to player", func(t *testing.T) {
		err := p1.AddTeam(&t1)
		err2 := p1.AddTeam(&t2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(p1.TeamIDs) != expectedLen {
			t.Errorf("Expected len of TeamIDs %v, got %v", expectedLen, len(p1.TeamIDs))
		}
	})
}
