package my_types

import "fmt"

// Player

func (d *Database) AddPlayer(player *Player) {
	player.ID = len(d.Players)
	d.Players[player.ID] = player
}

func (d *Database) DeletePlayer(playerID int) error {
	if _, ok := d.Players[playerID]; !ok {
		return fmt.Errorf("playerID %v does not exist", playerID)
	}
	delete(d.Players, playerID)
	return nil
}

func (d *Database) GetPlayer(playerID int) (*Player, error) {
	if _, ok := d.Players[playerID]; !ok {
		return nil, fmt.Errorf("playerID %v does not exist", playerID)
	}
	return d.Players[playerID], nil
}

// Team

func (d *Database) AddTeam(team *Team) {
	team.ID = len(d.Teams)
	d.Teams[team.ID] = team
}

func (d *Database) DeleteTeam(teamID int) error {
	if _, ok := d.Players[teamID]; !ok {
		return fmt.Errorf("teamID %v does not exist", teamID)
	}
	delete(d.Players, teamID)
	return nil
}

func (d *Database) GetTeam(teamID int) (*Team, error) {
	if _, ok := d.Players[teamID]; !ok {
		return nil, fmt.Errorf("teamID %v does not exist", teamID)
	}
	return d.Teams[teamID], nil
}

// Club
func (d *Database) AddClub(club *Club) {
	club.ID = len(d.Clubs)
	d.Clubs[club.ID] = club
}

func (d *Database) DeleteClub(clubID int) error {
	if _, ok := d.Players[clubID]; !ok {
		return fmt.Errorf("clubID %v does not exist", clubID)
	}
	delete(d.Players, clubID)
	return nil
}

func (d *Database) GetClub(clubID int) (*Club, error) {
	if _, ok := d.Players[clubID]; !ok {
		return nil, fmt.Errorf("clubID %v does not exist", clubID)
	}
	return d.Clubs[clubID], nil
}
