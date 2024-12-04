package mytypes_test

import (
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeletePlayer(t *testing.T) {

	var player1 my_types.Player
	player1.Name = "p1"
	expectedTeamListOfPLayer := []*my_types.Team{}

	var team1 my_types.Team
	team1.Name = "t1"
	team1.PlayerList = []*my_types.Player{&player1}
	player1.TeamList = []*my_types.Team{&team1}
	expectedLenOfPlayerListOfTeam := 0

	var club1 my_types.Club
	club1.Name = "TSG Heilbronn"
	club1.PlayerList = append(club1.PlayerList, &player1)
	club1.TeamList = append(club1.TeamList, &team1)
	expectedLenOfPlayerListOfClub := 0
	var expectedError error = nil

	t.Run(fmt.Sprintf("Remove a player from the club playerlist %s", club1.Name), func(t *testing.T) {
		actualError := club1.DeletePlayer(&player1)
		for i := range club1.PlayerList {
			if player1.Name != "" ||
				player1.Age != 0 ||
				player1.Ranking != 0 ||
				player1.Material != nil ||
				player1.TeamList != nil ||
				len(team1.PlayerList) != expectedLenOfPlayerListOfTeam ||
				len(club1.PlayerList) != expectedLenOfPlayerListOfClub ||
				actualError != nil {
				t.Errorf(`-----------------------------------
			Testcase:		%v
			Expecting:		(%v, %v, %v, %v)
			Actual:			(%v, %v, %v, %v)
			Fail`, i+1,
					expectedTeamListOfPLayer, expectedLenOfPlayerListOfTeam, expectedLenOfPlayerListOfClub, expectedError,
					player1.TeamList, len(team1.PlayerList), len(club1.PlayerList), actualError)
			} else {
				fmt.Printf(`-----------------------------------
			Testcase:		%v
			Expecting:		(%v, %v, %v, %v)
			Actual:			(%v, %v, %v, %v)
			Pass
			`, i+1,
					expectedTeamListOfPLayer, expectedLenOfPlayerListOfTeam, expectedLenOfPlayerListOfClub, expectedError,
					player1.TeamList, len(team1.PlayerList), len(club1.PlayerList), actualError)
			}
		}
	})
}
