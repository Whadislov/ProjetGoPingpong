package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

func NewPlayer(playerName string, db *mt.Database) (*mt.Player, error) {
	if playerName == "" {
		return nil, fmt.Errorf("player name cannot be empty")
	}

	p := &mt.Player{
		ID:       len(db.Players),
		Name:     playerName,
		Age:      -1,
		Ranking:  -1,
		Material: DefaultPlayerMaterial(),
		TeamIDs:  map[int]string{},
		ClubIDs:  map[int]string{},
	}

	db.AddPlayer(p)
	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}
