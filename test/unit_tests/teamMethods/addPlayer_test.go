package teammethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	p1 := mt.Player{
		ID:        uuid.New(),
		Firstname: "p1",
	}

	t1 := mt.Team{
		ID:        uuid.New(),
		Name:      "t1",
		PlayerIDs: map[uuid.UUID]string{p1.ID: p1.Firstname},
	}

	p2 := mt.Player{
		ID:        uuid.New(),
		Firstname: "p2",
	}

	expectedLen := 2
	expectedError := "player p1 is already in team t1"

	t.Run("Add player to team", func(t *testing.T) {
		err := t1.AddPlayer(&p1)
		err2 := t1.AddPlayer(&p2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(t1.PlayerIDs) != expectedLen {
			t.Errorf("Expected len of PlayerIDs %v, got %v", expectedLen, len(t1.PlayerIDs))
		}
	})
}
