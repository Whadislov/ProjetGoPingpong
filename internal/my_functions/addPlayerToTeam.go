package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func AddPlayerToTeam(p *mt.Player, t *mt.Team) error {

	err := p.AddTeam(t)
	if err != nil {
		return fmt.Errorf("error when adding player %v to team %v: %w", p.Name, t.Name, err)
	}
	err = t.AddPlayer(p)
	if err != nil {
		return fmt.Errorf("error when adding player %v to team %v: %w", p.Name, t.Name, err)
	}

	return nil
}
