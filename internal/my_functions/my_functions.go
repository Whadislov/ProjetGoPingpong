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
	_, err := c.FindPlayer(playerName)
	return err == nil
}

func isTeamAlreadyDefined(teamName string, c mt.Club)(bool) {
	_, err := c.FindTeam(teamName)
	return err == nil
}

func NewPlayer(playerName string, club *mt.Club)(*mt.Player, error) {
	if isPlayerAlreadyDefined(playerName, *club) {
		return nil, fmt.Errorf("player %v already exists", playerName)
	}
	p := &mt.Player{
		Name:		playerName,
		Age:		0,
		Ranking: 	0,
		Material: 	[]string{"Unknown", "Unknown", "Unknown"},
		TeamList: 	nil,
	}
	// Add player on player list
	club.AddPlayer(p)

	fmt.Printf("Player %v sucessfully created.\n", playerName)
	return p, nil
}

func NewTeam(teamName string, club *mt.Club)(*mt.Team, error) {
	if isTeamAlreadyDefined(teamName, *club) {
		return nil, fmt.Errorf("team %v already exists", teamName)
	}
	t := &mt.Team{
		Name:			teamName,
		PlayerList:		nil,
	}
	// Add team on team list
	club.AddTeam(t)

	fmt.Printf("Team %v sucessfully created.\n", teamName)
	return t, nil
}

func AddPlayerToTeam(p *mt.Player, teamName string, c *mt.Club) (error) {
	if p == nil {
		return fmt.Errorf("player is nil")
	}

	_, err := c.FindTeam(teamName)
	if err != nil {
		return fmt.Errorf("team %s not found: %w", teamName, err)
	}

	err = c.AddPlayerToTeam(p, teamName)
	if  err == nil {
		fmt.Printf("%s has been successfully added in %s\n", p.Name, teamName)
	} else {
		return fmt.Errorf("%s has not been successfully added in %s : %w", p.Name, teamName, err)
	}
	return nil
}

func RemovePlayerFromTeam(p *mt.Player, teamName string, c *mt.Club) (e error) {
	if c.RemovePlayerFromTeam(p, teamName) == nil {
		fmt.Println(p.Name, "removed from", teamName)
	} else {
		fmt.Println("has not been successfully removed from", teamName)
	}
	return nil
}

func DeletePlayer(p *mt.Player, c *mt.Club) (error) {
	return c.RemovePlayer(p)
}

func DeleteTeam(teamName string, c *mt.Club) (error) {
	return c.RemoveTeam(teamName)
}

func GetName(x interface{})(string) {
	switch v:= x.(type) {
	case mt.Player:{
		return v.Name
	}
	case *mt.Player:{
		return v.Name
	}
	case mt.Team:{
		return v.Name
	}
	case *mt.Team:{
		return v.Name
	}
	case mt.Club:{
		return v.Name
	}
	case *mt.Club:{
		return v.Name
	}
	default: {
		return ""
	}
	}
}