package mytypes_test

import (
	"testing"
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)


func TestDeleteTeam(t *testing.T) {

	var player1 my_types.Player
	player1.Name = "p1"
	expectedLenOfTeamListOfPLayer := 0

	var team1 my_types.Team
	team1.Name = "t1"
	team1.PlayerList = []*my_types.Player{&player1}
	player1.TeamList = []*my_types.Team{&team1}

	var club1 my_types.Club
	club1.Name = "TSG Heilbronn"
	club1.PlayerList = append(club1.PlayerList, &player1)
	club1.TeamList = append(club1.TeamList, &team1)
	expectedLenOfTeamListOfClub  := 0
	var expectedError error = nil


	t.Run(fmt.Sprintf("Remove team from the club teamlist %s", club1.Name), func(t *testing.T) {
		actualError := club1.DeleteTeam(&team1)
	for i := range club1.PlayerList {
		if team1.Name != "" || 
		team1.PlayerList != nil ||
		len(player1.TeamList) != expectedLenOfTeamListOfPLayer ||
		len(club1.TeamList) != expectedLenOfTeamListOfClub ||
		actualError != nil {
			t.Errorf(`-----------------------------------
			Testcase:		%v
			Expecting:		(%v, %v, %v)
			Actual:			(%v, %v, %v)
			Fail`, i+1, 
			expectedLenOfTeamListOfPLayer, expectedLenOfTeamListOfClub, expectedError,
			len(player1.TeamList), len(club1.TeamList), actualError)
		} else {
			fmt.Printf(`-----------------------------------
			Testcase:		%v
			Expecting:		(%v, %v, %v)
			Actual:			(%v, %v, %v)
			Pass
			`, i+1, 
			expectedLenOfTeamListOfPLayer, expectedLenOfTeamListOfClub, expectedError,
			len(player1.TeamList), len(club1.TeamList), actualError)
		}
	}
	})
}