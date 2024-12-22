package my_functions

import (
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// AddPlayerToTeam adds a player to a team by updating both the player's and the team's records.
// Returns an error if the player is already in the team or if there is an issue with the operation.
func AddPlayerToTeam(p *mt.Player, t *mt.Team) error {

	err := p.AddTeam(t)
	if err != nil {
		return fmt.Errorf("error when adding team %v to player %v: %w", t.Name, p.Name, err)
	}
	err = t.AddPlayer(p)
	if err != nil {
		return fmt.Errorf("error when adding player %v to team %v: %w", p.Name, t.Name, err)
	}

	return nil
}
