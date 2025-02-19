package my_functions

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
)

// NewTeam creates a new team with the given name and adds it to the database.
// Returns the created team and an error if the team name is empty or if there is an issue with the operation.
func NewTeam(teamName string, db *mt.Database) (*mt.Team, error) {
	if teamName == "" {
		return nil, fmt.Errorf("team name cannot be empty")
	}

	t := &mt.Team{
		ID:        uuid.New(),
		Name:      teamName,
		PlayerIDs: make(map[uuid.UUID]string),
		ClubID:    make(map[uuid.UUID]string, 1), // Capacity 1
		IsNew:     true,
	}

	db.AddTeam(t)
	log.Printf("Team %v sucessfully created.", teamName)
	return t, nil
}
