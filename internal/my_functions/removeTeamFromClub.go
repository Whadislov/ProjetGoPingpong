package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func RemoveTeamFromClub(t *mt.Team, c *mt.Club) error {

	err := c.RemoveTeam(t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully removed from %s. Reason : %w", t.Name, c.Name, err)
	}

	err = t.RemoveClub(c)
	if err != nil {
		return fmt.Errorf("%s has not been successfully removed from %s. Reason : %w", t.Name, c.Name, err)
	}
	return nil
}
