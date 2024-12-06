package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func AddTeamToClub(t *mt.Team, c *mt.Club) error {

	err := t.AddClub(c)
	if err != nil {
		return fmt.Errorf("error when adding team %v to club %v: %w", t.Name, c.Name, err)
	}
	err = c.AddTeam(t)
	if err != nil {
		return fmt.Errorf("error when adding team %v to club %v: %w", t.Name, c.Name, err)
	}

	return nil
}
