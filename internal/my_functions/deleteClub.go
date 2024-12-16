package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func DeleteClub(c *mt.Club, db *mt.Database) error {
	// Remove player depedences
	var playerIDs []int
	if len(c.PlayerIDs) > 0 {
		for playerID := range c.PlayerIDs {
			playerIDs = append(playerIDs, playerID)
			err := db.Players[playerID].RemoveClub(c)
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
		for _, playerID := range playerIDs {
			err := c.RemovePlayer(db.Players[playerID])
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
	}

	// Remove team depedences
	var teamIDs []int
	if len(c.TeamIDs) > 0 {
		for teamID := range c.TeamIDs {
			teamIDs = append(teamIDs, teamID)
			err := db.Teams[teamID].RemoveClub(c)
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
		for _, teamID := range teamIDs {
			err := c.RemoveTeam(db.Teams[teamID])
			if err != nil {
				return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
			}
		}
	}

	// Delete club
	err := db.DeleteClub(c.ID)
	if err != nil {
		return fmt.Errorf("error when deleting club %s: %w", c.Name, err)
	}

	return nil
}
