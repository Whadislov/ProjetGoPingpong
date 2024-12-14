package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func DeletePlayer(p *mt.Player, db *mt.Database) error {
	// Remove club depedences
	var clubIDs []int
	if len(p.ClubIDs) > 0 {
		for clubID := range p.ClubIDs {
			clubIDs = append(clubIDs, clubID)
			err := db.Clubs[clubID].RemovePlayer(p)
			if err != nil {
				return fmt.Errorf("error when deleting player %s: %w", p.Name, err)
			}
		}
		for clubID := range clubIDs {
			err := p.RemoveClub(db.Clubs[clubID])
			if err != nil {
				return fmt.Errorf("error when deleting player %s: %w", p.Name, err)
			}
		}
	}

	// Remove team depedences
	var teamIDs []int
	if len(p.TeamIDs) > 0 {
		for teamID := range p.TeamIDs {
			teamIDs = append(teamIDs, teamID)
			err := db.Teams[teamID].RemovePlayer(p)
			if err != nil {
				return fmt.Errorf("error when deleting player %s: %w", p.Name, err)
			}
		}
		for teamID := range teamIDs {
			err := p.RemoveTeam(db.Teams[teamID])
			if err != nil {
				return fmt.Errorf("error when deleting player %s: %w", p.Name, err)
			}
		}
	}

	// Delete player
	err := db.DeletePlayer(p.ID)
	if err != nil {
		return fmt.Errorf("error when deleting player %s: %w", p.Name, err)
	}

	return nil
}
