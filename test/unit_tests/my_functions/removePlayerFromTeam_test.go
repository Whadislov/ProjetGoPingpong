package my_functions_test

import (
	"testing"

	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestRemovePlayerFromTeam(t *testing.T) {
	player := mt.Player{
		Name: "player",
	}
	
	team := mt.Team{
		Name: "team",
		PlayerList: []*mt.Player{&player},
	}
	player.TeamList = []*mt.Team{&team}
	
	club := mt.Club{
		Name: "club",
		PlayerList: []*mt.Player{&player},
		TeamList: []*mt.Team{&team},
	}

	err := mf.RemovePlayerFromTeam(&player, &team, &club)

	var expectedLen = 0



	t.Run("Remove player from team", func(t *testing.T) {
		if len(team.PlayerList) != 0 {
			t.Errorf("Playerlist of team issue: got %v, expected %v", len(team.PlayerList), expectedLen)
		}
		if len(player.TeamList) != 0 {
			t.Errorf("Teamlist of player issue: got %v, expected %v", len(player.TeamList), expectedLen)
		}
		if err != nil {
			t.Errorf("Error issue: got %v, expected %v", err, nil)
		}
	})
}
