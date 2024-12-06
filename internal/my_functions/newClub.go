package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"log"
)

func NewClub(clubName string, db *mt.Database) (*mt.Club, error) {
	if clubName == "" {
		return nil, fmt.Errorf("club name cannot be empty")
	}

	c := &mt.Club{
		ID:   len(db.Clubs),
		Name: clubName,
	}

	db.AddClub(c)
	log.Printf("Club %v sucessfully created.", clubName)
	return c, nil
}
