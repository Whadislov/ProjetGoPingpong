package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func NewClub(name string)(c mt.Club) {
	c.Name = name
	c.PlayerList = nil
	c.TeamList = nil
	return c
}

func isPlayerAlreadyDefined(playerName string, c mt.Club)(bool) {
	for _, player := range c.GetPlayerList() {
		if player.Name == playerName {
			return true
		}
	}
	return false
}

func isTeamAlreadyDefined(teamName string, c mt.Club)(bool) {
	for _, team := range c.GetTeamList() {
		if team.Name == teamName {
			return true
		}
	}
	return false
}

func NewPlayer(playerName string, club *mt.Club)(p mt.Player, err error) {
	if isPlayerAlreadyDefined(playerName, *club) {
		fmt.Printf("Error ! Player %v already exists.", playerName)
		return mt.Player{}, fmt.Errorf("player already exists")
	}
	p.Name = playerName
	p.Age = 0
	p.Ranking = 0
	p.Material = []string{"Unknown", "Unknown", "Unknown"}
	p.TeamList = nil
	fmt.Printf("Player %v sucessfully created.\n", playerName)
	club.PlayerList = append(club.PlayerList, p)
	return p, err
}

func NewTeam(teamName string, club *mt.Club)(t mt.Team, err error) {
	if isTeamAlreadyDefined(teamName, *club) {
		fmt.Printf("Error ! Team %v already exists.\n", teamName)
		return mt.Team{}, fmt.Errorf("team already exists")
	}
	t.Name = teamName
	t.PlayerList = nil
	fmt.Printf("Team %v sucessfully created.\n", teamName)
	club.TeamList = append(club.TeamList, t)
	return t, err
}

func isPlayerInTeam(p *mt.Player, t *mt.Team) (bool) {
	fmt.Println("t.PlayerList", t.Name)
	fmt.Println("t.PlayerList", t.PlayerList)
	for _, player := range t.PlayerList {
		fmt.Println("Say ok")
		if player.Name == p.Name {

			return true
		}
	}
	return false
}

//Ajouter vérification de doublon (2x meme équipe, 2x meme joueur)
func AddPlayerToTeam(p *mt.Player, t *mt.Team) (err error) {
	if isPlayerInTeam(p, t) {
		fmt.Printf("Error ! %v is already in %v.\n", p.Name, t.Name)
		return fmt.Errorf("%v is already in %v", p.Name, t.Name)
	}
	p.TeamList = append(p.TeamList, t.Name)
	t.PlayerList = append(t.PlayerList, *p)
	fmt.Printf("Successfully added %v in %v.\n", p.Name, t.Name)
	return nil
}

func RemovePlayerFromTeam(p *mt.Player, t *mt.Team) (err error) {
	if !isPlayerInTeam(p, t) {
		fmt.Printf("Error ! %v is not in team %v.\n", p.Name, t.Name)
		return fmt.Errorf("%v is not in team %v", p.Name, t.Name)
	} 
	// Remove player name from the team player list
	fmt.Printf("%v removed from %v\n", p.Name, t.Name)
	t.RemovePlayer(p)
	// Remove the team from the team list of the player
	fmt.Printf("%v removed from %v\n", t.Name, p.Name)
	p.RemoveTeam(t.Name)
	return nil
}

func DeletePlayer(p *mt.Player, c *mt.Club) (err error) {
	if p.IsEmpty() {
		fmt.Println("Error ! Player does not exist")
		return fmt.Errorf("player does not exist")
	}
	
	if p.TeamList != nil {
		for _, team := range p.TeamList {
			t := p.GetTeam(team, *c)
			RemovePlayerFromTeam(p, t)
		}
	}
	fmt.Printf("%v removed from %v\n", p.Name, c.Name)
	c.RemovePlayer(p.Name)

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

