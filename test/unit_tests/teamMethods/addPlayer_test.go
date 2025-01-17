package teammethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	t1 := mt.Team{
		ID:        0,
		Name:      "t1",
		PlayerIDs: map[int]string{0: "p1"},
	}

	p1 := mt.Player{
		ID:   0,
		Name: "p1",
	}

	p2 := mt.Player{
		ID:   1,
		Name: "p2",
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
