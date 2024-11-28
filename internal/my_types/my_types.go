package my_types

import (
	"fmt"
)



type Player struct {
	Name 		string
	Age 		int
	Ranking		int
	Material 	[]string
	TeamList	[]string  //Max 2
}

type Team struct {
	Name 		string
	PlayerList 	[]Player
}

type Club struct {
	Name 		string
	PlayerList	[]Player
	TeamList	[]Team
}

type PlayerMatch struct{
	League				string
	Player				string
	Adversary 			string
	PlayerMatchOutcome	Outcome			
}

type Outcome struct {
	Victory	string
	Defeat	string
	Draw	string
}

type TeamMatch struct{
	League				string
	TeamComposition 	map[int]string
	
}

type Match struct {
	HomeTeam			Team
	GuestTeam			Team
	TeamMatchOutcome	Outcome
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
func (t *TeamMatch) SetTeamComposition(teamComposition map[int]string) {
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
	p.TeamList == nil
}

func (t *Team) IsEmpty()(bool) {
	return t.Name == "" &&
	t.PlayerList == nil
}

func (c *Club) IsEmpty()(bool) {
	return c.Name == "" &&
	c.PlayerList == nil &&
	c.TeamList == nil
}

func (p *Player) Show()(err error) {
	if p.IsEmpty() {
		fmt.Println("Error ! Player does not exist")
		return fmt.Errorf("Player does not exist")
	}

	fmt.Printf("%v is %v years old. Ranking : %v. %v plays with %v on the forehand, %v on the backhand and %v.\n", p.Name, p.Age, p.Ranking, p.Name, p.Material[0], p.Material[1], p.Material[2])
	switch len(p.TeamList){
	case 0: {
		fmt.Printf("%v does not have a team.\n", p.Name)
		return nil
	}
	case 1: {
		fmt.Printf("%v plays in %v.\n", p.Name, p.TeamList[0])
		return nil
	}
	case 2: {
		fmt.Printf("%v plays in %v and in %v.\n", p.Name, p.TeamList[0], p.TeamList[1])
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
	fmt.Println("Team name:", t.Name)
	switch len(t.PlayerList)  {
	case 0: {
		fmt.Printf("%v has 0 player.\n", t.Name)
	}
	case 1: {
		fmt.Printf("%v has 1 player.\n", t.Name)
		fmt.Println("The player is :", t.PlayerList[0])
	}
	default: {
		fmt.Printf("%v has %v players.\n", t.Name, len(t.PlayerList))
		fmt.Println("The players are :")
		for i := 0; i < len(t.PlayerList); i++ {
			fmt.Println(t.PlayerList[i])
		}
	}
	}
	return nil
}

func (c *Club) Show()(err error) {
	if c.IsEmpty() {
		fmt.Println("Error ! Club does not exist")
		return fmt.Errorf("Club does not exist")
	}

	fmt.Println("Club name :", c.Name)
	switch len(c.TeamList)  {
	case 0: {
		fmt.Printf("%v has 0 team.\n", c.Name)
	}
	case 1: {
		fmt.Printf("%v has 1 team.\n", c.Name)
		fmt.Println("The team is :", c.TeamList[0])
	}
	default: {
		fmt.Printf("%v has %v teams.\n", c.Name, len(c.TeamList))
		fmt.Println("The teams are :")
		for i := 0; i < len(c.TeamList); i++ {
			fmt.Println(c.TeamList[i])
		}
	}
	}
	switch len(c.PlayerList)  {
	case 0: {
		fmt.Printf("%v has 0 player.\n", c.Name)
	}
	case 1: {
		fmt.Printf("%v has 1 player.\n", c.Name)
		fmt.Println("The player is :", c.PlayerList[0])
	}
	default: {
		fmt.Printf("%v has %v players.\n", c.Name, len(c.PlayerList))
		fmt.Println("The players are :")
		for i := 0; i < len(c.PlayerList); i++ {
			fmt.Println(c.PlayerList[i])
		}
	}
	}
	return nil
}

func (t *Team) AddPlayer(player Player) {
	t.PlayerList = append(t.PlayerList, player)
}

func (c *Club) AddPlayer(player Player) {
	c.PlayerList = append(c.PlayerList, player)
}

func (t *Team) RemovePlayer(player *Player) {
	if len(t.PlayerList) == 1 {
		t.PlayerList = nil
	}
	for i, p := range t.PlayerList {
		if p.Name == player.Name {
			fmt.Println("%88888888888888888888888888888888")
			fmt.Println(t.PlayerList[len(t.PlayerList)-1])
			fmt.Println("888888888888888888888888888888888")
			t.PlayerList[i] = t.PlayerList[len(t.PlayerList)-1]
			t.PlayerList = t.PlayerList[:len(t.PlayerList)-1]
			break
		}
	}
}

func (c *Club) RemovePlayer(player string) {
	for i, p := range c.PlayerList {
		if p.Name == player {
			c.PlayerList[i] = c.PlayerList[len(c.PlayerList)-1]
			c.PlayerList = c.PlayerList[:len(c.PlayerList)-1]
			break
		}
	}
}

func (p *Player) AddTeam(team string) {
	p.TeamList = append(p.TeamList, team)
}

func (c *Club) AddTeam(team Team) {
	c.TeamList = append(c.TeamList, team)
}

func (p *Player) RemoveTeam(team string) {
	for i, t := range p.TeamList {
		if t == team {
			p.TeamList[i] = p.TeamList[len(p.TeamList)-1]
			p.TeamList = p.TeamList[:len(p.TeamList)-1]
			break
		}
	}
}

func (c *Club) RemoveTeam(team string) {
	for i, t := range c.TeamList {
		if t.Name == team {
			c.TeamList[i] = c.TeamList[len(c.TeamList)-1]
			c.TeamList = c.TeamList[:len(c.TeamList)-1]
			break
		}
	}
}

func (c *Club) GetPlayerList()([]Player) {
	return c.PlayerList
}

func (p *Player) GetTeam(teamName string, c Club)(*Team) {
	for _, team := range c.TeamList {
		fmt.Println("c.TeamList", c.TeamList)
		if team.Name == teamName {
			return &team
		}
	}
	return nil
}

func (c *Club) GetTeamList()([]Team) {
	return c.TeamList
}


/*

if len(t.TeamComposition) > 0 {
	fmt.Printf("The team composition for the next match is :\n")
	for j := 0; j < len(t.TeamComposition); j++ {
		fmt.Printf("Position %v : %v\n", j+1, t.TeamComposition[j])
	}
}
*/