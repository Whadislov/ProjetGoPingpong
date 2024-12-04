package my_functions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeleteTeam(t *testing.T) {
	t1 := mt.Team{
		Name:       "t1",
		PlayerList: []*mt.Player{},
	}

	c1 := mt.Club{
		Name:     "c1",
		TeamList: []*mt.Team{&t1},
	}

	err := mf.DeleteTeam(&t1, &c1)

	t.Run("Delete team", func(t *testing.T) {
		if len(c1.TeamList) != 0 {
			t.Errorf("Club's team list len issue: got %v, expected %v", len(c1.TeamList), 0)
		}
		if err != nil {
			t.Errorf("Error issue: got %v, expected %v", err, nil)
		}
	})
}
