package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

func NewClub(clubName string, db *mt.Database) (*mt.Club, error) {
	if clubName == "" {
		err := fmt.Errorf("club name cannot be empty")
		return nil, err
	}

	c := &mt.Club{
		Name:       clubName,
		PlayerList: []*mt.Player{},
		TeamList:   []*mt.Team{},
	}

	db.AddClub(c)
	log.Printf("Club %v sucessfully created.", clubName)
	return c, nil
}
