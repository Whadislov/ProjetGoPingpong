package my_functions

import (
	"log"
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

const DefaultMaterial = "Unknown"

func DefaultPlayerMaterial() []string {
    return []string{DefaultMaterial, DefaultMaterial, DefaultMaterial}
}

func NewClub(name string)(c mt.Club) {
	c.Name = name
	c.PlayerList = nil
	c.TeamList = nil
	return c
}

func isPlayerAlreadyDefined(playerName string, c *mt.Club)(error) {
	if c == nil {
		return fmt.Errorf("club is not defined")
	}
	_, err := c.FindPlayer(playerName)
	return err
}

func isTeamAlreadyDefined(teamName string, c *mt.Club)(error) {
	if c == nil {
		return fmt.Errorf("club is not defined")
	}
	_, err := c.FindTeam(teamName)
	return err
}

func NewPlayer(playerName string, club *mt.Club)(mt.Player, error) {
	
	if err := isPlayerAlreadyDefined(playerName, club); err == nil {
		return mt.Player{}, fmt.Errorf("player %v already exists", playerName)
	}
	p := mt.Player{
		Name:		playerName,
		Age:		0,
		Ranking: 	0,
		Material: 	DefaultPlayerMaterial(),
		TeamList: 	[]*mt.Team{},
	}
	// Add player on player list
	club.AddPlayer(&p)

	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}

func NewTeam(teamName string, club *mt.Club)(mt.Team, error) {
	if err := isTeamAlreadyDefined(teamName, club); err == nil {
		return mt.Team{}, fmt.Errorf("team %v already exists", teamName)
	}
	t := mt.Team{
		Name:			teamName,
		PlayerList:		[]*mt.Player{},
	}
	// Add team on team list
	club.AddTeam(&t)

	log.Printf("Team %v sucessfully created.", teamName)
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
		log.Println(p.Name, "removed from", teamName)
	} else {
		log.Println("has not been successfully removed from", teamName)
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