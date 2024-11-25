package my_types

type Player struct {
	Name 		string
	Age 		int
	Ranking		int
	Material 	[]string
	Teams		[]Team
}

type Team struct {
	Name 			string
	PlayerList 		[]Player
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

func (p *Player) SetPlayerAge(age int) {
	p.Age = age
}

func (p *Player) SetPlayerRanking(ranking int) {
	p.Ranking = ranking
}

func (p *Player) SetPlayerMaterial(forehand string, backhand string, blade string) {
	material := []string{forehand, backhand, blade}
	p.Material = material
}

func (p *Player) SetPlayerTeam(team Team) {
	p.Teams = append(p.Teams, team)
}

func (t *Team) SetTeamPlayerList(playerList []Player) {
	t.PlayerList = playerList
}

func (t *Team) SetTeamComposition(teamComposition map[string]int) {
	t.TeamComposition = teamComposition
}
