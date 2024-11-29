package my_types_test

import (
	"testing"
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	//"github.com/stretchr/testify/mock"
)




func TestAddPlayerToTeamTest(t *testing.T) {
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
	Player1 has 0 team in teamList
	Team1 has 0 player in playerList

	Action
	Add Player1 to Team1

	Expected
	Player2 has Team1 in teamList
	Team1 has Player1 in playerList
	*/

	club1 := setupClub()
	var player1 my_types.Player
	player1.Name =  "p1"
	player1.TeamList = []*my_types.Team{}

	var team1 my_types.Team
	team1.Name = "t1"
	team1.PlayerList = []*my_types.Player{}

	club1.PlayerList = append(club1.PlayerList, &player1)
	club1.TeamList = append(club1.TeamList, &team1)
	//Action
	actualError1 := club1.AddPlayerToTeam(&player1, team1.Name)

	//Expected
	expectedPlayerTeamList1	:= []*my_types.Team{&team1}
	expectedTeamPlayerList1	:= []*my_types.Player{&player1}
	var expectedError1 error 
	expectedError1 = nil

	/* Test 2 : 
	Setup
	Player2 has Team2 in teamList
	Team2 has 0 player in playerList

	Action
	Add Player2 to Team2

	Expected
	Player2 has Team2 in teamList
	Team2 has Player2 in playerList
	err := fmt.Errorf("team Team2 is already in team list of Player2, the link issue is now solved")
	
	*/

	club2 := setupClub()
	var player2 my_types.Player
	player2.Name = "p2"
	player2.TeamList = []*my_types.Team{}
	
	var team2 my_types.Team
	team2.Name = "t2"
	team2.PlayerList = []*my_types.Player{&player2}

	club2.PlayerList = append(club2.PlayerList, &player2)
	club2.TeamList = append(club2.TeamList, &team2)
	//Action
	actualError2 := club2.AddPlayerToTeam(&player2, team2.Name)

	//Expected
	expectedPlayerTeamList2	:= []*my_types.Team{&team2}
	expectedTeamPlayerList2	:= []*my_types.Player{&player2}
	expectedError2 := fmt.Errorf("player %s is already in team %s", player2.Name, team2.Name)

	/* Test 3 : 
	Setup
	Player3 has Team3 in teamList
	Team3 has 0 player in playerList

	Action
	Add Player3 to Team3

	Expected
	Player3 has Team3 in teamList
	Team3 has Player3 in playerList
	err := fmt.Errorf("team Team3 and player Player3 are already in each other's respective list")
	*/

	club3 := setupClub()
	var player3 my_types.Player
	player3.Name =  "p3"
	
	var team3 my_types.Team
	team3.Name = "t3"
	team3.PlayerList = []*my_types.Player{}
	player3.TeamList = []*my_types.Team{&team3}

	club3.PlayerList = append(club3.PlayerList, &player3)
	club3.TeamList = append(club3.TeamList, &team3)
	//Action
	actualError3 := club3.AddPlayerToTeam(&player3, team3.Name)

	//Expected
	expectedPlayerTeamList3	:= []*my_types.Team{&team3}
	expectedTeamPlayerList3	:= []*my_types.Player{&player3}
	expectedError3 := fmt.Errorf("team %s and player %s are already in each other's respective list", team3.Name, player3.Name)


	/* Test 4 : 
	Setup
	Player4 has Team4 in teamList
	Team4 has Player4 in playerList

	Action
	Add Player4 to Team4

	Expected
	Player4 has Team4 in teamList
	Team4 has Player4 in playerList
	err := fmt.Errorf("team Team4 and player Player4 are already in each other's respective list")
	*/

	club4 := setupClub()
	var player4 my_types.Player
	player4.Name =  "p4"
	
	var team4 my_types.Team
	team4.Name = "t4"
	team4.PlayerList = []*my_types.Player{&player4}
	player4.TeamList = []*my_types.Team{&team4}

	club4.PlayerList = append(club4.PlayerList, &player4)
	club4.TeamList = append(club4.TeamList, &team4)
	//Action
	actualError4 := club4.AddPlayerToTeam(&player4, team4.Name)

	//Expected
	expectedPlayerTeamList4	:= []*my_types.Team{&team4}
	expectedTeamPlayerList4	:= []*my_types.Player{&player4}
	expectedError4 := fmt.Errorf("team %s and player %s are already in each other's respective list", team4.Name, player4.Name)

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
		t.Run(fmt.Sprintf("Add a player %s to team %s", test.player.Name, test.team.Name), func(t *testing.T) {
			if test.expectedPlayerTeamList[0] !=  test.player.TeamList[0] ||
			test.expectedTeamPlayerList[0] !=  test.team.PlayerList[0] ||
			(test.expectedError != nil && test.actualError == nil) || (test.expectedError == nil && test.actualError != nil) {
				t.Errorf(`-----------------------------------
				Testcase:		%v
				Expecting:		(%v, %v, %v)
				Actual:			(%v, %v, %v)
				Fail`, i+1, 
				test.expectedPlayerTeamList[0], test.expectedTeamPlayerList[0], test.expectedError,
				test.player.TeamList[0], test.team.PlayerList[0], test.actualError)
		} else {
			fmt.Printf(`-----------------------------------
			Testcase:		%v
			Expecting:		(%v, %v, %v)
			Actual:			(%v, %v, %v)
			Pass
			`, i+1, 
			test.expectedPlayerTeamList[0], test.expectedTeamPlayerList[0], test.expectedError,
			test.player.TeamList[0], test.team.PlayerList[0], test.actualError)
		}
	})
}
}