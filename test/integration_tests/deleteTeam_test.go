package integrationtests_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeleteTeam(t *testing.T) {

	p1 := mt.Player{
		ID:      0,
		Name:    "p1",
		TeamIDs: map[int]string{0: "t1"},
		ClubIDs: map[int]string{0: "c1"},
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
		PlayerIDs: map[int]string{p1.ID: p1.Name},
		TeamIDs:   map[int]string{t1.ID: t1.Name},
	}

	d := mt.Database{
		Clubs:   map[int]*mt.Club{0: &c1},
		Teams:   map[int]*mt.Team{0: &t1},
		Players: map[int]*mt.Player{0: &p1},
	}

	expectedLenDTeams := 0
	expectedLenClubTeamIDs := 0
	expectedLenPlayerTeamIDs := 0

	t.Run("Delete Club", func(t *testing.T) {
		err := mf.DeleteTeam(&t1, &d)
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if len(d.Teams) != expectedLenDTeams {
			t.Errorf("Expected LenDTeams %v, got %v", expectedLenDTeams, len(d.Teams))
		}
		if len(c1.TeamIDs) != expectedLenClubTeamIDs {
			t.Errorf("Expected LenClubTeamIDs %v, got %v", expectedLenClubTeamIDs, len(c1.TeamIDs))
		}
		if len(p1.TeamIDs) != expectedLenPlayerTeamIDs {
			t.Errorf("Expected LenPlayerTeamIDs %v, got %v", expectedLenPlayerTeamIDs, len(p1.TeamIDs))
		}
	})
}
