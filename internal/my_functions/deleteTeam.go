package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func DeleteTeam(t *mt.Team, db *mt.Database) error {
	// Remove club depedences
	if len(t.ClubID) > 0 {
		// need clubID for t.RemoveClub
		clubID := -1
		for ID := range t.ClubID {
			clubID = ID
			err := db.Clubs[clubID].RemoveTeam(t)
			if err != nil {
				return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
			}
		}
		err := t.RemoveClub(db.Clubs[clubID])
		if err != nil {
			return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
		}
	}

	// Remove player depedences
	var playerIDs []int
	if len(t.PlayerIDs) > 0 {
		for playerID := range t.PlayerIDs {
			playerIDs = append(playerIDs, playerID)
			err := db.Players[playerID].RemoveTeam(t)
			if err != nil {

				return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
			}
		}
		for _, playerID := range playerIDs {
			err := t.RemovePlayer(db.Players[playerID])
			if err != nil {
				return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
			}
		}
	}

	// Delete team

	err := db.DeleteTeam(t.ID)
	if err != nil {
		return fmt.Errorf("error when deleting team %s: %w", t.Name, err)
	}

	return nil
}
