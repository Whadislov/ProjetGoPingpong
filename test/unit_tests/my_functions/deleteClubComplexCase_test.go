package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeleteClubComplexCase(t *testing.T) {

	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		TeamIDs:   map[int]string{0: "t1"},
		ClubIDs:   map[int]string{0: "c1"},
	}

	t1 := mt.Team{
		ID:        0,
		Name:      "t1",
		PlayerIDs: map[int]string{0: "p1"},
		ClubID:    map[int]string{0: "c1"},
	}

	c1 := mt.Club{
		ID:        0,
		Name:      "c1",
		PlayerIDs: map[int]string{p1.ID: p1.Firstname},
		TeamIDs:   map[int]string{t1.ID: t1.Name},
	}

	d := mt.Database{
		Clubs:   map[int]*mt.Club{0: &c1},
		Teams:   map[int]*mt.Team{0: &t1},
		Players: map[int]*mt.Player{0: &p1},
	}

	expectedLenDClubs := 0
	expectedLenPlayerClubIDs := 0
	expectedLenTeamClubID := 0

	t.Run("Delete Club", func(t *testing.T) {
		err := mf.DeleteClub(&c1, &d)
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if len(d.Clubs) != expectedLenDClubs {
			t.Errorf("Expected LenDClubs %v, got %v", expectedLenDClubs, len(d.Clubs))
		}
		if len(p1.ClubIDs) != expectedLenPlayerClubIDs {
			t.Errorf("Expected LenPlayerClubIDs %v, got %v", expectedLenPlayerClubIDs, len(p1.ClubIDs))
		}
		if len(t1.ClubID) != expectedLenTeamClubID {
			t.Errorf("Expected LenPlayerClubIDs %v, got %v", expectedLenTeamClubID, len(t1.ClubID))
		}
	})
}
