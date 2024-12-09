package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestRemovePlayerFromClub(t *testing.T) {
	c1 := mt.Club{
		ID:        0,
		Name:      "c1",
		PlayerIDs: map[int]string{0: "p1", 2: "p3"},
	}

	p1 := mt.Player{
		ID:      0,
		Name:    "p1",
		ClubIDs: map[int]string{0: "c1"},
	}

	p2 := mt.Player{
		ID:   1,
		Name: "p2",
	}

	p3 := mt.Player{
		ID:   2,
		Name: "p3",
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	p4 := mt.Player{
		ID:   3,
		Name: "p4",
	}

	expectedLen1 := 1
	expectedError2 := "p2 has not been successfully removed from c1. Reason : p2 is not in club c1"
	expectedError3 := "p3 has not been successfully removed from c1. Reason : p3 is not in club c1"
	expectedError4 := "p4 has not been successfully removed from c2. Reason : p4 is not in club c2"

	t.Run("Remove Player from club", func(t *testing.T) {
		err := mf.RemovePlayerFromClub(&p1, &c1)
		err2 := mf.RemovePlayerFromClub(&p2, &c1)
		if err != nil || len(c1.PlayerIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}

		err3 := mf.RemovePlayerFromClub(&p3, &c1)
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
		err4 := mf.RemovePlayerFromClub(&p4, &c2)
		if err4 == nil {
			t.Errorf("Expected error %v, got %v", expectedError4, err4)
		}
	})
}
