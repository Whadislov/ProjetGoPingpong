package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

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
