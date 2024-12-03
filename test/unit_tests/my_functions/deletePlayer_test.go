package my_functions_test

import (
	"testing"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestDeletePlayer(t *testing.T) {
	p1 := mt.Player{
		Name: "p1",
		TeamList: []*mt.Team{},
	}

	c1 := mt.Club{
		Name: "c1",
		PlayerList: []*mt.Player{&p1},
	}

	err := mf.DeletePlayer(&p1, &c1)


	t.Run("Delete player", func(t *testing.T){
		if len(c1.PlayerList) != 0 {
			t.Errorf("Club's player list len issue: got %v, expected %v", len(c1.PlayerList), 0)
		}
		if err != nil {
			t.Errorf("Error issue: got %v, expected %v", err, nil)
		}
	})
}
