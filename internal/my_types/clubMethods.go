package my_types

import (
	"fmt"
)

/*
type Club struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"player_id_list"`
	TeamIDs   map[int]string `json:"team_id_list"`
	// map[player.ID] = player.Name, nil per default
	// map[team.ID] = team.Name, nil per default
}
*/

func (c *Club) SetClubID(id int) {
	c.ID = id
}

func (c *Club) SetClubName(name string) {
	c.Name = name
}

func (c *Club) AddPlayer(player *Player) error {
	if _, ok := c.PlayerIDs[player.ID]; ok {
		return fmt.Errorf("player %v is already in club %v", player.Name, c.Name)
	}
	c.PlayerIDs[player.ID] = player.Name
	return nil
}

func (c *Club) AddTeam(team *Team) error {
	if _, ok := c.TeamIDs[team.ID]; ok {
		return fmt.Errorf("team %v is already in club %v", team.Name, c.Name)
	}
	c.TeamIDs[team.ID] = team.Name
	return nil
}

func (c *Club) RemovePlayer(player *Player) error {
	if _, ok := c.PlayerIDs[player.ID]; !ok {
		return fmt.Errorf("player %v is not in club %v", player.Name, c.Name)
	}
	delete(c.PlayerIDs, player.ID)
	return nil
}

func (c *Club) RemoveTeam(team *Team) error {
	if _, ok := c.TeamIDs[team.ID]; !ok {
		return fmt.Errorf("team %v is not in club %v", team.Name, c.Name)
	}
	delete(c.TeamIDs, team.ID)
	return nil
}

func (c *Club) HasPlayer() bool {
	return len(c.PlayerIDs) > 0
}

func (c *Club) HasTeam() bool {
	return len(c.TeamIDs) > 0
}
