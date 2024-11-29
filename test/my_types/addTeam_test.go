package my_types_test

import (
	"testing"
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestAddTeam(t *testing.T) {

	var club my_types.Club
	club.Name = "TSG Heilbronn"
	club.TeamList = []*my_types.Team{}

	var team my_types.Team
	team.Name = "Mannschaft 1"
	expectedTeamList  := []*my_types.Team{&team}


	t.Run(fmt.Sprintf("Add player to player list of club %s", club.Name), func(t *testing.T) {
		club.AddTeam(&team)
	for i := range club.TeamList {
		if club.TeamList[i] != &team {
			t.Errorf("Team list of %s is currently %v and is expected to be %v", club.Name, club.TeamList[i], expectedTeamList[i])
		} else {
			fmt.Printf("Team list of %s is currently %v and is expected to be %v", club.Name, club.TeamList[i], expectedTeamList[i])
		}
	}
	})
}