package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

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
