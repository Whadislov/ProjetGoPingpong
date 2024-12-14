package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestRemovePlayerFromTeam(t *testing.T) {
	t1 := mt.Team{
		ID:        0,
		Name:      "t1",
		PlayerIDs: map[int]string{0: "p1", 2: "p3"},
	}

	p1 := mt.Player{
		ID:      0,
		Name:    "p1",
		TeamIDs: map[int]string{0: "t1"},
	}

	p2 := mt.Player{
		ID:   1,
		Name: "p2",
	}

	p3 := mt.Player{
		ID:   2,
		Name: "p3",
	}

	t2 := mt.Team{
		ID:   1,
		Name: "t1",
	}

	p4 := mt.Player{
		ID:   3,
		Name: "p4",
	}

	expectedLen1 := 1
	expectedError2 := "p2 has not been successfully removed from t1. Reason : p2 is not in Team t1"
	expectedError3 := "p3 has not been successfully removed from t1. Reason : p3 is not in Team t1"
	expectedError4 := "p4 has not been successfully removed from t2. Reason : p4 is not in Team t2"

	t.Run("Remove Player from Team", func(t *testing.T) {
		err := mf.RemovePlayerFromTeam(&p1, &t1)
		err2 := mf.RemovePlayerFromTeam(&p2, &t1)
		if err != nil || len(t1.PlayerIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}

		err3 := mf.RemovePlayerFromTeam(&p3, &t1)
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
		err4 := mf.RemovePlayerFromTeam(&p4, &t2)
		if err4 == nil {
			t.Errorf("Expected error %v, got %v", expectedError4, err4)
		}
	})
}
