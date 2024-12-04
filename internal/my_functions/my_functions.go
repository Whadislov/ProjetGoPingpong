package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

const DefaultMaterial = "Unknown"

func DefaultPlayerMaterial() []string {
	return []string{DefaultMaterial, DefaultMaterial, DefaultMaterial}
}

func NewClub(name string) (c mt.Club, err error) {
	if name == "" {
		err = fmt.Errorf("club name cannot be empty")
		return mt.Club{}, err
	}

	c.Name = name
	c.PlayerList = []*mt.Player{}
	c.TeamList = []*mt.Team{}
	return c, nil
}

func IsPlayerAlreadyDefined(playerId int, c *mt.Club) error {
	if c == nil {
		return fmt.Errorf("club is nil")
	}

	if c.Name == "" {
		return fmt.Errorf("club is not defined")
	}

	for _, player := range c.PlayerList {
		if player.Id == playerId {
			return fmt.Errorf("player %v is already defined", playerId)
		}
	}
	return nil
}

func IsTeamAlreadyDefined(teamName string, c *mt.Club) error {
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

func NewPlayer(playerName string, club *mt.Club) (*mt.Player, error) {
	if playerName == "" {
		return nil, fmt.Errorf("player name cannot be empty")
	}

	p := &mt.Player{
		Id:       len(club.PlayerList),
		Name:     playerName,
		Age:      0,
		Ranking:  0,
		Material: DefaultPlayerMaterial(),
		TeamList: []*mt.Team{},
	}
	// Add player on player list
	club.AddPlayer(p)

	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}

func NewTeam(teamName string, club *mt.Club) (*mt.Team, error) {
	if teamName == "" {
		return nil, fmt.Errorf("team name cannot be empty")
	}

	if err := IsTeamAlreadyDefined(teamName, club); err != nil {
		return nil, fmt.Errorf("team %v already exists", teamName)
	}
	t := &mt.Team{
		Name:       teamName,
		PlayerList: []*mt.Player{},
	}
	// Add team on team list
	club.AddTeam(t)

	log.Printf("Team %v sucessfully created.", teamName)
	return t, nil
}

func AddPlayerToTeam(p *mt.Player, t *mt.Team, c *mt.Club) error {
	if p == nil {
		return fmt.Errorf("player has not been successfully added in %s. Reason : player is nil", t.Name)
	}

	err := c.FindTeam(t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully added in %s. Reason : %w", p.Name, t.Name, err)
	}

	err = c.AddPlayerToTeam(p, t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully added in %s. Reason : %w", p.Name, t.Name, err)
	}
	return nil
}

func RemovePlayerFromTeam(p *mt.Player, t *mt.Team, c *mt.Club) error {
	err := c.RemovePlayerFromTeam(p, t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully removed from %s. Reason : %w", p.Name, t.Name, err)
	}
	return nil
}

func DeletePlayer(p *mt.Player, c *mt.Club) error {
	if err := c.DeletePlayer(p); err != nil {
		return fmt.Errorf("error when deleting player %s. Reason : %w", p.Name, err)
	}
	return nil
}

func DeleteTeam(t *mt.Team, c *mt.Club) error {
	if err := c.DeleteTeam(t); err != nil {
		return fmt.Errorf("error when deleting team %s. Reason : %w", t.Name, err)
	}
	return nil
}

func GetName(x interface{}) string {
	switch v := x.(type) {
	case mt.Player:
		{
			return v.Name
		}
	case *mt.Player:
		{
			return v.Name
		}
	case mt.Team:
		{
			return v.Name
		}
	case *mt.Team:
		{
			return v.Name
		}
	case mt.Club:
		{
			return v.Name
		}
	case *mt.Club:
		{
			return v.Name
		}
	default:
		{
			return ""
		}
	}
}
