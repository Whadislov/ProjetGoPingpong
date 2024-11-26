package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)


func NewPlayer(name string)(p mt.Player) {
	p.Name = name
	p.Age = 0
	p.Ranking = 0
	p.Material = []string{"Unknown", "Unknown", "Unknown"}
	p.Teams = nil
	return p
}

func NewTeam(name string)(t mt.Team) {
	t.Name = name
	t.PlayerList = nil
	t.TeamComposition = nil
	return t
}

func isPlayerInTeam(p mt.Player, t mt.Team) (bool) {
	if p.Teams == nil {
		return false
	}
	fmt.Println("Team: ", t.Name)
	fmt.Println("PlayList: ", t.PlayerList)
	for _, playerFromList := range t.PlayerList {
		if playerFromList.Name == p.Name {
			fmt.Println("Ok, player in PlayList")
			return true
		}
	}
	return false
}

func AddPlayerToTeam(p *mt.Player, t *mt.Team) (err error) {
	if isPlayerInTeam(*p, *t) {
		fmt.Printf("Error ! %v is already in %v.\n", p.Name, t.Name)
		return fmt.Errorf("%v is already in %v", p.Name, t.Name)
	}
	p.Teams = append(p.Teams, *t)
	t.PlayerList = append(t.PlayerList, *p)
	return nil
}

func RemovePlayerFromTeam(p *mt.Player, t *mt.Team) (err error) {
	fmt.Println("Valeur du bool: " ,isPlayerInTeam(*p, *t))
	if !isPlayerInTeam(*p, *t) {
		fmt.Printf("Error ! %v is not in team %v.\n", p.Name, t.Name)
		return fmt.Errorf("%v is not in team %v", p.Name, t.Name)
	} 
	for i := range t.PlayerList {
		if t.PlayerList[i].Name == p.Name {
			t.PlayerList[i] = t.PlayerList[len(t.PlayerList)-1]
			t.PlayerList = t.PlayerList[:len(t.PlayerList)-1]
		}
	}
	for i := range p.Teams {
		if p.Teams[i].Name == t.Name {
			p.Teams[i] = p.Teams[len(p.Teams)-1]
			p.Teams = p.Teams[:len(p.Teams)-1]
		}
	}
	return nil
}

func DeletePlayer(p *mt.Player) (err error) {
	if p.IsEmpty() {
		fmt.Println("Error ! Player does not exist")
		return fmt.Errorf("player does not exist")
	}
	
	if p.Teams != nil {
		for _, team := range p.Teams {
			fmt.Printf("%v removed from %v\n", p.Name, team.Name)
			RemovePlayerFromTeam(p, &team)
		}
	}
	*p = mt.Player{}
	fmt.Println("Player has been successfully deleted")
	return nil 
}

func DeleteTeam(t *mt.Team) (err error) {
	if t.IsEmpty() {
		fmt.Println("Error ! Team does not exist")
		return fmt.Errorf("team does not exist")
	}

	if t.PlayerList != nil {
		for _, player := range t.PlayerList {
			RemovePlayerFromTeam(&player, t)
		}
	}
	*t = mt.Team{}
	fmt.Println("Team has been successfully deleted")
	return nil 
}

