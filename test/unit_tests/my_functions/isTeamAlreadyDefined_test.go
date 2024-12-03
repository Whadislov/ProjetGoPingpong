package my_functions_test

import (
	"fmt"
	"testing"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestIsTeamAlreadyDefined(t *testing.T) {

	team1 := mt.Team{
		Name: "t1",
	}

	team2 := mt.Team{
		Name: "t1",
	}

	club1 := mt.Club{
		Name: "club",
		TeamList: []*mt.Team{},
	}

	club2 := mt.Club{}

	var club3 *mt.Club = nil

	err1 := mf.IsTeamAlreadyDefined(team1.Name, &club1)
	club1.TeamList = append(club1.TeamList, &team1)
	err2 := mf.IsTeamAlreadyDefined(team2.Name, &club1)
	err3 := mf.IsTeamAlreadyDefined(team1.Name, &club2)
	err4 := mf.IsTeamAlreadyDefined(team1.Name, club3)


	t.Run("Test if team is already defined", func(t *testing.T) {
		if err1 != nil {
			t.Errorf("Err1 issue: got %v, expected %v", err1, nil)
		}
		if err2 == nil {
			t.Errorf("Err2 issue: got %v, expected %v", err2, fmt.Errorf("player %v is already defined", team1.Name))
		}
		if err3 == nil {
			t.Errorf("Err3 issue: got %v, expected %v", err3, fmt.Errorf("club is not defined"))
		}
		if err4 == nil {
			t.Errorf("Err3 issue: got %v, expected %v", err4, fmt.Errorf("club is nil"))
		}
	})
}
