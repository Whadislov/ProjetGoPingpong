package myfunctions_test

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestAddPlayerToClub(t *testing.T) {
	c1 := mt.Club{
		ID:        0,
		Name:      "c1",
		PlayerIDs: map[int]string{},
	}

	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		ClubIDs:   map[int]string{},
	}

	p2 := mt.Player{
		ID:        1,
		Firstname: "p2",
		ClubIDs:   map[int]string{0: "c1"},
	}

	c2 := mt.Club{
		ID:        1,
		Name:      "c2",
		PlayerIDs: map[int]string{2: "p3"},
	}

	p3 := mt.Player{
		ID:        2,
		Firstname: "p3",
		ClubIDs:   map[int]string{},
	}

	expectedLen1 := 1
	expectedError2 := "error when adding player p2 to club c1: player p2 is already in club c1"
	expectedError3 := "error when adding player p3 to club c2: player p3 is already in club c2"

	t.Run("Add Player to club", func(t *testing.T) {
		err := mf.AddPlayerToClub(&p1, &c1)
		err2 := mf.AddPlayerToClub(&p2, &c1)
		if err != nil || len(c1.PlayerIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}

		err3 := mf.AddPlayerToClub(&p3, &c2)
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
	})
}
