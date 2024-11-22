package my_types

type Player struct {
	Name 		string
	Age 		int
	Ranking		int
	Material 	[]string
	Teams		[]string
}

type Team struct {
	Name 		string
	PlayerList 		[]string
	TeamComposition map[string]int
}

type PlayerMatch struct{
	League						string
	ReceiverTeam 				string
	GuestTeam 					string
	ReceiverPlayerMatchOutcome	Outcome			
}

type Outcome struct {
	Victory	string
	Defeat	string
	Draw	string
}


type TeamMatch struct{
	League						string
	ReceiverTeam 				string
	GuestTeam 					string
	ReceiverTeamMatchOutcome	Outcome
}


