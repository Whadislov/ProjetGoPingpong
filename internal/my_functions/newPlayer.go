package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

func NewPlayer(playerName string, club *mt.Club, db *mt.Database) (*mt.Player, error) {
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
	db.AddPlayer(p)

	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}
