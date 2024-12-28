package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

// NewTeam creates a new team with the given name and adds it to the database.
// Returns the created team and an error if the team name is empty or if there is an issue with the operation.
func NewTeam(teamName string, db *mt.Database) (*mt.Team, error) {
	if teamName == "" {
		return nil, fmt.Errorf("team name cannot be empty")
	}

	t := &mt.Team{
		ID:        len(db.Teams),
		Name:      teamName,
		PlayerIDs: map[int]string{},
		ClubID:    map[int]string{},
	}

	db.AddTeam(t)
	log.Printf("Team %v sucessfully created.", teamName)
	return t, nil
}
