package myfunctions_test

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestAddPlayerToTeam(t *testing.T) {
	t1 := mt.Team{
		ID:        0,
		Name:      "t1",
		PlayerIDs: map[int]string{},
	}

	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		TeamIDs:   map[int]string{},
	}

	p2 := mt.Player{
		ID:        1,
		Firstname: "p2",
		TeamIDs:   map[int]string{0: "t1"},
	}

	y2 := mt.Team{
		ID:        1,
		Name:      "y2",
		PlayerIDs: map[int]string{2: "p3"},
	}

	p3 := mt.Player{
		ID:        2,
		Firstname: "p3",
		TeamIDs:   map[int]string{},
	}

	expectedLen1 := 1
	expectedError2 := "error when adding Player p2 to Team t1: Player p2 is already in Team t1"
	expectedError3 := "error when adding Player p3 to Team y2: Player p3 is already in Team y2"

	t.Run("Add Player to Team", func(t *testing.T) {
		err := mf.AddPlayerToTeam(&p1, &t1)
		err2 := mf.AddPlayerToTeam(&p2, &t1)
		if err != nil || len(t1.PlayerIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}

		err3 := mf.AddPlayerToTeam(&p3, &y2)
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
	})
}
