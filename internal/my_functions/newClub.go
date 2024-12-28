package my_functions

import (
	"fmt"
	"log"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// NewClub creates a new club with the given name and adds it to the database.
// Returns the created club and an error if the club name is empty or if there is an issue with the operation.
func NewClub(clubName string, db *mt.Database) (*mt.Club, error) {
	if clubName == "" {
		return nil, fmt.Errorf("club name cannot be empty")
	}

	c := &mt.Club{
		ID:        len(db.Clubs),
		Name:      clubName,
		PlayerIDs: map[int]string{},
		TeamIDs:   map[int]string{},
	}

	db.AddClub(c)
	log.Printf("Club %v sucessfully created.", clubName)
	return c, nil
}
