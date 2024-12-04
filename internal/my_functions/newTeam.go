package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

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
