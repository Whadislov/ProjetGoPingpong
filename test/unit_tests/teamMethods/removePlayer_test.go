package teammethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestRemovePlayer(t *testing.T) {
	t1 := mt.Team{
		ID:        0,
		Name:      "t1",
		PlayerIDs: map[int]string{0: "t1"},
	}

	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
	}

	p2 := mt.Player{
		ID:        1,
		Firstname: "p2",
	}

	expectedLen1 := 0
	expectedError2 := "player p1 is not in team t2"

	t.Run("Remove team from player", func(t *testing.T) {
		err := t1.RemovePlayer(&p1)
		err2 := t1.RemovePlayer(&p2)
		if err != nil || len(t1.PlayerIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
