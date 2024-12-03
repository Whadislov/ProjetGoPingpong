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

func NewClub(name string)(c mt.Club, err error) {
	if name == "" {
		err = fmt.Errorf("club name cannot be empty")
		return mt.Club{}, err
	}

	c.Name = name
	c.PlayerList = []*mt.Player{}
	c.TeamList = []*mt.Team{}
	return c, nil
}

func IsPlayerAlreadyDefined(playerName string, c *mt.Club)(error) {
	if c == nil {
		return fmt.Errorf("club is nil")
	}

	if c.Name == "" {
		return fmt.Errorf("club is not defined")
	}

	for _, player := range c.PlayerList {
		if player.Name == playerName {
			return fmt.Errorf("player %v is already defined", playerName)
		}
	}
	return nil
}

func IsTeamAlreadyDefined(teamName string, c *mt.Club)(error) {
	if c == nil {
		return fmt.Errorf("club is nil")
	}

	if c.Name == "" {
		return fmt.Errorf("club is not defined")
	}

	for _, team := range c.TeamList {
		if team.Name == teamName {
			return fmt.Errorf("team %v is already defined", teamName)
		}
	}
	return nil
}

func NewPlayer(playerName string, club *mt.Club)(*mt.Player, error) {
	if playerName == "" {
		return nil, fmt.Errorf("player name cannot be empty")
	}
	
	if err := IsPlayerAlreadyDefined(playerName, club); err != nil {
		return nil, fmt.Errorf("player %v already exists", playerName)
	}
	p := &mt.Player{
		Name:		playerName,
		Age:		0,
		Ranking: 	0,
		Material: 	DefaultPlayerMaterial(),
		TeamList: 	[]*mt.Team{},
	}
	// Add player on player list
	club.AddPlayer(p)

	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}

func NewTeam(teamName string, club *mt.Club)(*mt.Team, error) {
	if teamName == "" {
		return nil, fmt.Errorf("team name cannot be empty")
	}

	if err := IsTeamAlreadyDefined(teamName, club); err != nil {
		return nil, fmt.Errorf("team %v already exists", teamName)
	}
	t := &mt.Team{
		Name:			teamName,
		PlayerList:		[]*mt.Player{},
	}
	// Add team on team list
	club.AddTeam(t)

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
		log.Printf("%s has been successfully added in %s", p.Name, teamName)
	} else {
		return fmt.Errorf("%s has not been successfully added in %s : %w", p.Name, teamName, err)
	}
	return nil
}

func RemovePlayerFromTeam(p *mt.Player, teamName string, c *mt.Club) (error) {
	if err := c.RemovePlayerFromTeam(p, teamName); err != nil {
		log.Printf("%v has not been successfully removed from %v. Reason : %v", p.Name, teamName, err)
		return err
	}
	log.Printf("%v has been successfully removed from %v", p.Name, teamName)
	return nil
}

func DeletePlayer(p *mt.Player, c *mt.Club) (error) {
	if err := c.DeletePlayer(p); err != nil {
		log.Printf("%v has not been successfully deleted from %v. Reason : %v", p.Name, c.Name, err)
		return err
	}
	log.Printf("%v has been successfully deleted from %v", p.Name, c.Name)
	return nil
}

func DeleteTeam(teamName string, c *mt.Club) (error) {
	if err := c.DeleteTeam(teamName); err != nil {
		log.Printf("%v has not been successfully deleted from %v. Reason : %v", teamName, c.Name, err)
		return err
	}
	log.Printf("%v has been successfully deleted from %v", teamName, c.Name)
	return nil
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