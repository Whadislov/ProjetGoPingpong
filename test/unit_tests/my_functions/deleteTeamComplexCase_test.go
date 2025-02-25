package myfunctions_test

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestDeleteTeamComplexCase(t *testing.T) {
	p1ID := uuid.New()
	t1Id := uuid.New()
	c1Id := uuid.New()
	p1Name := "p1"
	t1Name := "t1"
	c1Name := "c1"

	p1 := mt.Player{
		ID:        p1ID,
		Firstname: p1Name,
		TeamIDs:   map[uuid.UUID]string{t1Id: t1Name},
		ClubIDs:   map[uuid.UUID]string{c1Id: c1Name},
	}

	t1 := mt.Team{
		ID:        t1Id,
		Name:      t1Name,
		PlayerIDs: map[uuid.UUID]string{p1ID: p1Name},
		ClubID:    map[uuid.UUID]string{c1Id: c1Name},
	}

	c1 := mt.Club{
		ID:        c1Id,
		Name:      c1Name,
		PlayerIDs: map[uuid.UUID]string{p1.ID: p1.Firstname},
		TeamIDs:   map[uuid.UUID]string{t1.ID: t1.Name},
	}

	d := mt.Database{
		Clubs:           map[uuid.UUID]*mt.Club{c1Id: &c1},
		Teams:           map[uuid.UUID]*mt.Team{t1Id: &t1},
		Players:         map[uuid.UUID]*mt.Player{p1ID: &p1},
		DeletedElements: map[string][]uuid.UUID{},
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
