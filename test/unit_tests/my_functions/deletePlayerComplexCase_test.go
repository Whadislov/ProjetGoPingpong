package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeletePlayerComplexCase(t *testing.T) {

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

	expectedLenDPlayers := 0
	expectedLenClubPlayerIDs := 0
	expectedLenTeamPlayerIDs := 0

	t.Run("Delete Club", func(t *testing.T) {
		err := mf.DeletePlayer(&p1, &d)
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if len(d.Players) != expectedLenDPlayers {
			t.Errorf("Expected LenDClubs %v, got %v", expectedLenDPlayers, len(d.Players))
		}
		if len(c1.PlayerIDs) != expectedLenClubPlayerIDs {
			t.Errorf("Expected LenPlayerClubIDs %v, got %v", expectedLenClubPlayerIDs, len(c1.PlayerIDs))
		}
		if len(t1.PlayerIDs) != expectedLenTeamPlayerIDs {
			t.Errorf("Expected LenPlayerClubIDs %v, got %v", expectedLenTeamPlayerIDs, len(t1.PlayerIDs))
		}
	})
}
