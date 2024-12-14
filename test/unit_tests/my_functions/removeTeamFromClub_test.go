package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestRemoveTeamFromClub(t *testing.T) {
	c1 := mt.Club{
		ID:      0,
		Name:    "c1",
		TeamIDs: map[int]string{0: "t1", 2: "t3"},
	}

	t1 := mt.Team{
		ID:     0,
		Name:   "t1",
		ClubID: map[int]string{0: "c1"},
	}

	t2 := mt.Team{
		ID:   1,
		Name: "t2",
	}

	t3 := mt.Team{
		ID:   2,
		Name: "t3",
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	t4 := mt.Team{
		ID:     3,
		Name:   "t4",
		ClubID: map[int]string{1: "c2"},
	}

	expectedLen1 := 1
	expectedError2 := "t2 has not been successfully removed from c1. Reason : t2 is not in club c1"
	expectedError3 := "t3 has not been successfully removed from c1. Reason : t3 is not in club c1"
	expectedError4 := "t4 has not been successfully removed from c2. Reason : t4 is not in club c2"

	t.Run("Remove team from club", func(t *testing.T) {
		err := mf.RemoveTeamFromClub(&t1, &c1)
		err2 := mf.RemoveTeamFromClub(&t2, &c1)
		if err != nil || len(c1.TeamIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}

		err3 := mf.RemoveTeamFromClub(&t3, &c1)
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
		err4 := mf.RemoveTeamFromClub(&t4, &c2)
		if err4 == nil {
			t.Errorf("Expected error %v, got %v", expectedError4, err4)
		}
	})
}
