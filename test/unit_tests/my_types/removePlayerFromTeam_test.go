package my_types_test

import (
	"testing"
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestRemovePlayerFromTeam(t *testing.T) {

	// Helper function
	setupClub := func() my_types.Club {
        return my_types.Club{Name: "c1"}
    }

	type Testcase struct {
		player					*my_types.Player
		team					*my_types.Team
		club					my_types.Club
		expectedPlayerTeamList	[]*my_types.Team
		expectedTeamPlayerList	[]*my_types.Player
		actualError				error
		expectedError			error
	}

	/* Test 1
	Setup
	Player1 has Team1 in teamList
	Team1 has Player1 in playerList

	Action
	Remove Player1 from Team1

	Expected
	Player1 has 0 team in teamList
	Team1 has 0 player in playerList
	*/

	/* Test 2 : 
	Setup
	Player2 has Team2 in teamList
	Team2 has 0 player in playerList

	Action
	Remove Player2 from Team2

	Expected
	Player2 has 0 team in teamList
	Team2 has 0 player in playerList
	err := fmt.Errorf("player Player2 does not belong to team Team2")
	
	*/

	/* Test 3 : 
	Setup
	Player3 has 0 team in teamList
	Team3 has Player3 in playerList

	Action
	Remove Player3 from Team3

	Expected
	Player3 has 0 team in teamList
	Team3 has 0 player in playerList
	err := fmt.Errorf("team Team3 is not in Player2's team list")
	*/

	
	/* Test 4 : 
	Setup
	Player4 has 0 team in teamList
	Team4 has 0 team in playerList

	Action
	Add Player4 to Team4

	Expected
	Player4 has 0 team in teamList
	Team4 has 0 team in playerList
	err := fmt.Errorf("team Team4 and player Player4 are not in each other's respective list")
	*/

	club1 := setupClub()
	var player1 my_types.Player
	player1.Name =  "p1"
	player1.TeamList = []*my_types.Team{}

	var team1 my_types.Team
	team1.Name = "t1"
	team1.PlayerList = []*my_types.Player{&player1}
	player1.TeamList = append(player1.TeamList, &team1)
	fmt.Println(team1.PlayerList[0])
	
	club1.PlayerList = append(club1.PlayerList, &player1)
	club1.TeamList = append(club1.TeamList, &team1)

	//Action
	actualError1 := club1.RemovePlayerFromTeam(&player1, team1.Name)

	//Expected
	expectedPlayerTeamList1	:= []*my_types.Team{}
	expectedTeamPlayerList1	:= []*my_types.Player{}
	var expectedError1 error = nil

	club2 := setupClub()
	var team2 my_types.Team
	team2.Name = "t2"
	team2.PlayerList = []*my_types.Player{}

	var player2 my_types.Player
	player2.Name = "p2"
	player2.TeamList = []*my_types.Team{&team2}
	
	club2.PlayerList = append(club2.PlayerList, &player2)
	club2.TeamList = append(club2.TeamList, &team2)
	//Action
	actualError2 := club2.RemovePlayerFromTeam(&player2, team2.Name)

	//Expected
	expectedPlayerTeamList2	:= []*my_types.Team{}
	expectedTeamPlayerList2	:= []*my_types.Player{}
	expectedError2 := fmt.Errorf("player %s does not belong to team %s", player2.Name, team2.Name)

	club3 := setupClub()
	var player3 my_types.Player
	player3.Name =  "p3"
	player3.TeamList = []*my_types.Team{}
	
	var team3 my_types.Team
	team3.Name = "t3"
	team3.PlayerList = []*my_types.Player{&player3}

	club3.PlayerList = append(club3.PlayerList, &player3)
	club3.TeamList = append(club3.TeamList, &team3)
	//Action
	actualError3 := club3.RemovePlayerFromTeam(&player3, team3.Name)

	//Expected
	expectedPlayerTeamList3	:= []*my_types.Team{}
	expectedTeamPlayerList3	:= []*my_types.Player{}
	expectedError3 := fmt.Errorf("team %s is not in player %s's team list", team3.Name, player3.Name)

	club4 := setupClub()
	var player4 my_types.Player
	player4.Name =  "p4"
	player4.TeamList = []*my_types.Team{}
	
	var team4 my_types.Team
	team4.Name = "t4"
	team4.PlayerList = []*my_types.Player{}

	club4.PlayerList = append(club4.PlayerList, &player4)
	club4.TeamList = append(club4.TeamList, &team4)
	//Action
	actualError4 := club4.RemovePlayerFromTeam(&player4, team4.Name)

	//Expected
	expectedPlayerTeamList4	:= []*my_types.Team{}
	expectedTeamPlayerList4	:= []*my_types.Player{}
	expectedError4 := fmt.Errorf("player %s does not belong to team %s", team4.Name, player4.Name)

	tests := []Testcase{
		{
			player: &player1,
			team: &team1,
			club: club1,
			expectedPlayerTeamList: expectedPlayerTeamList1,
			expectedTeamPlayerList: expectedTeamPlayerList1,
			actualError: actualError1,
			expectedError: expectedError1,
		},
		{
			player: &player2,
			team: &team2,
			club: club2,
			expectedPlayerTeamList: expectedPlayerTeamList2,
			expectedTeamPlayerList: expectedTeamPlayerList2,
			actualError: actualError2,
			expectedError: expectedError2,
		},
		{
			player: &player3,
			team: &team3,
			club: club3,
			expectedPlayerTeamList: expectedPlayerTeamList3,
			expectedTeamPlayerList: expectedTeamPlayerList3,
			actualError: actualError3,
			expectedError: expectedError3,
		},
		{
			player: &player4,
			team: &team4,
			club: club4,
			expectedPlayerTeamList: expectedPlayerTeamList4,
			expectedTeamPlayerList: expectedTeamPlayerList4,
			actualError: actualError4,
			expectedError: expectedError4,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Remove player %s from team %s", test.player.Name, test.team.Name), func(t *testing.T) {
			if (test.expectedPlayerTeamList == nil && test.player.TeamList == nil) ||
			(test.expectedTeamPlayerList == nil && test.team.PlayerList == nil) ||
			(test.expectedError != nil && test.actualError == nil) || (test.expectedError == nil && test.actualError != nil) {
				t.Errorf(`-----------------------------------
				Testcase:		%v
				Expecting:		(%v, %v, %v)
				Actual:			(%v, %v, %v)
				Fail`, i+1, 
				test.expectedPlayerTeamList, test.expectedTeamPlayerList, test.expectedError,
				test.player.TeamList, test.team.PlayerList, test.actualError)
		} else {
			fmt.Printf(`-----------------------------------
			Testcase:		%v
			Expecting:		(%v, %v, %v)
			Actual:			(%v, %v, %v)
			Pass
			`, i+1, 
			test.expectedPlayerTeamList, test.expectedTeamPlayerList, test.expectedError,
			test.player.TeamList, test.team.PlayerList, test.actualError)
		}
	})
}
}

