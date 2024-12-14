package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddTeamToClub(t *testing.T) {
	c1 := mt.Club{
		ID:      0,
		Name:    "c1",
		TeamIDs: map[int]string{},
	}

	t1 := mt.Team{
		ID:     0,
		Name:   "t1",
		ClubID: map[int]string{},
	}

	t2 := mt.Team{
		ID:     1,
		Name:   "t2",
		ClubID: map[int]string{0: "c1"},
	}

	c2 := mt.Club{
		ID:      1,
		Name:    "c2",
		TeamIDs: map[int]string{2: "t3"},
	}

	t3 := mt.Team{
		ID:     2,
		Name:   "t3",
		ClubID: map[int]string{},
	}

	expectedLen1 := 1
	expectedError2 := "error when adding Team t2 to club c1: Team t2 is already in club c1"
	expectedError3 := "error when adding Team t3 to club c2: Team t3 is already in club c2"

	t.Run("Add Team to club", func(t *testing.T) {
		err := mf.AddTeamToClub(&t1, &c1)
		err2 := mf.AddTeamToClub(&t2, &c1)
		if err != nil || len(c1.TeamIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}

		err3 := mf.AddTeamToClub(&t3, &c2)
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
	})
}
