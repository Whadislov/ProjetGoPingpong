package my_types

import "fmt"

/*
type Team struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"player_id_list"`
	ClubID    map[int]string `json:"club_id"`
	// map[player.ID] = player.Name, nil per default
	// map[club.ID] = club.Name, nil per default
}
*/

func (t *Team) SetTeamID(id int) {
	t.ID = id
}

func (t *Team) SetTeamName(name string) {
	t.Name = name
}

func (t *Team) AddPlayer(player *Player) error {
	if _, ok := t.PlayerIDs[player.ID]; ok {
		return fmt.Errorf("player %v is already in team %v", player.Name, t.Name)
	}
	t.PlayerIDs[player.ID] = player.Name
	return nil
}

func (t *Team) AddClub(club *Club) error {
	if len(t.ClubID) > 0 {
		return fmt.Errorf("team %v is already in a club", t.Name)
	}
	t.ClubID[club.ID] = club.Name
	return nil
}

func (t *Team) RemovePlayer(player *Player) error {
	if _, ok := t.PlayerIDs[player.ID]; !ok {
		return fmt.Errorf("player %v is not in team %v", player.Name, t.Name)
	}
	delete(t.PlayerIDs, player.ID)
	return nil
}

func (t *Team) RemoveClub(club *Club) error {
	if _, ok := t.ClubID[club.ID]; !ok {
		return fmt.Errorf("team %v is not in club %v", t.Name, club.Name)
	}
	delete(t.ClubID, club.ID)
	return nil
}

func (t *Team) HasPlayer() bool {
	return len(t.PlayerIDs) > 0
}

func (t *Team) HasClub() bool {
	return len(t.ClubID) > 0
}
