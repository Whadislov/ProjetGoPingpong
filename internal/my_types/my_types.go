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

func (p *Player) setMaterial(forehand string, backhand string, blade string) (){
	material := []string{forehand, backhand, blade}
	p.Material = material
}

func (p *Player) NewPlayer(name string, age int, ranking int, material []string, teams []string){
	p.Name = name
	p.Age = age
	if ranking == 0 {
		p.Ranking = 0
	} else {
		p.Ranking  = ranking
	}
	if material == nil {
		p.setMaterial("","","")
	} else {
		p.Material = material
	}
	if teams == nil {
		p.Teams = nil
	} else {
		for _, team := range teams {
			p.Teams = append(p.Teams, team)
		}
	}
}

func (t *Team) addPlayerToTeam(player *Player){
	t.PlayerList = append(t.PlayerList, player.Name)
	player.Teams = append(player.Teams, t.Name)

}
