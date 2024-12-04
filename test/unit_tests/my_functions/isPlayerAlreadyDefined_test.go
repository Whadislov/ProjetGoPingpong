package my_functions_test

import (
	"fmt"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestIsPlayerAlreadyDefined(t *testing.T) {

	player1 := mt.Player{
		Name: "p1",
	}

	player2 := mt.Player{
		Name: "p1",
	}

	club1 := mt.Club{
		Name:       "club",
		PlayerList: []*mt.Player{},
	}

	club2 := mt.Club{}

	var club3 *mt.Club = nil

	err1 := mf.IsPlayerAlreadyDefined(player1.Name, &club1)
	club1.PlayerList = append(club1.PlayerList, &player1)
	err2 := mf.IsPlayerAlreadyDefined(player2.Name, &club1)
	err3 := mf.IsPlayerAlreadyDefined(player1.Name, &club2)
	err4 := mf.IsPlayerAlreadyDefined(player1.Name, club3)

	t.Run("Test if player is already defined", func(t *testing.T) {
		if err1 != nil {
			t.Errorf("Err1 issue: got %v, expected %v", err1, nil)
		}
		if err2 == nil {
			t.Errorf("Err2 issue: got %v, expected %v", err2, fmt.Errorf("player %v is already defined", player1.Name))
		}
		if err3 == nil {
			t.Errorf("Err3 issue: got %v, expected %v", err3, fmt.Errorf("club is not defined"))
		}
		if err4 == nil {
			t.Errorf("Err3 issue: got %v, expected %v", err4, fmt.Errorf("club is nil"))
		}
	})
}
