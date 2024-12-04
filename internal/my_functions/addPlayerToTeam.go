package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func AddPlayerToTeam(p *mt.Player, t *mt.Team, c *mt.Club) error {
	if p == nil {
		return fmt.Errorf("player has not been successfully added in %s. Reason : player is nil", t.Name)
	}

	err := c.FindTeam(t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully added in %s. Reason : %w", p.Name, t.Name, err)
	}

	err = c.AddPlayerToTeam(p, t)
	if err != nil {
		return fmt.Errorf("%s has not been successfully added in %s. Reason : %w", p.Name, t.Name, err)
	}
	return nil
}
