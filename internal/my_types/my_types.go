package my_types

import (
	"fmt"
)

type Player struct {
	Name 		string
	Age 		int
	Ranking		int
	Material 	[]string
	Teams		[]Team  //Max 2
}

type Team struct {
	Name 			string
	PlayerList 		[]Player
	TeamComposition map[int]string
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

func (t *Team) SetTeamPlayerList(playerList []Player) {
	t.PlayerList = playerList
}

// ajouter erreur si la value de la clé (joueur) ne fait pas partie de la liste des joueurs
func (t *Team) SetTeamComposition(teamComposition map[int]string) {
	t.TeamComposition = teamComposition
}

type IsEmpty interface {
	IsEmpty()
}

func (p *Player) IsEmpty()(bool) {
	return p.Name == "" &&
	p.Age == 0 &&
	p.Ranking == 0 && 
	p.Material == nil &&
	p.Teams == nil
}

func (t *Team) IsEmpty()(bool) {
	return t.Name == "" &&
	t.PlayerList == nil &&
	t.TeamComposition == nil
}

type Show interface {
	Show()
}

func (p *Player) Show()(err error) {
	if p.IsEmpty() {
		fmt.Println("Error !Player does not exist")
		return fmt.Errorf("Player does not exist")
	}

	fmt.Printf("%v is %v years old. Ranking : %v. %v plays with %v on the forehand, %v on the backhand and %v.\n", p.Name, p.Age, p.Ranking, p.Name, p.Material[0], p.Material[1], p.Material[2])
	switch len(p.Teams){
	case 0: {
		fmt.Printf("%v does not have a team.\n", p.Name)
		return nil
	}
	case 1: {
		fmt.Printf("%v plays in %v.\n", p.Name, p.Teams[0].Name)
		return nil
	}
	case 2: {
		fmt.Printf("%v plays in %v and in %v.\n", p.Name, p.Teams[0].Name, p.Teams[1].Name)
		return nil
	}
	default: {
		fmt.Printf("%v plays in more than 2 teams.\n", p.Name)
		return nil
	}
	}
}

func (t *Team) Show()(err error) {
	if t.IsEmpty() {
		fmt.Println("Error ! Team does not exist")
		return fmt.Errorf("Team does not exist")
	}
	fmt.Printf("%v has %v players. The players are :\n", t.Name, len(t.PlayerList))
	for i := 0; i < len(t.PlayerList); i++ {
		fmt.Println(t.PlayerList[i].Name)
	}
	fmt.Printf("The team composition for the next match is :\n")
	for j := 0; j < len(t.TeamComposition); j++ {
		fmt.Printf("Position %v : %v\n", j+1, t.TeamComposition[j])
	}
	return nil
}
