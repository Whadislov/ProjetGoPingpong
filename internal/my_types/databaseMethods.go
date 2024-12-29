package my_types

import "fmt"

// AddPlayer adds a new player to the database.
func (d *Database) AddPlayer(player *Player) {
	d.Players[player.ID] = player
}

// DeletePlayer removes a player from the database by their ID.
// Returns an error if the player does not exist.
func (d *Database) DeletePlayer(playerID int) error {
	if _, ok := d.Players[playerID]; !ok {
		return fmt.Errorf("playerID %v does not exist", playerID)
	}
	delete(d.Players, playerID)
	return nil
}

// GetPlayer retrieves a player from the database by their ID.
// Returns an error if the player does not exist.
func (d *Database) GetPlayer(playerID int) (*Player, error) {
	if _, ok := d.Players[playerID]; !ok {
		return nil, fmt.Errorf("playerID %v does not exist", playerID)
	}
	return d.Players[playerID], nil
}

// AddTeam adds a new team to the database.
func (d *Database) AddTeam(team *Team) {
	d.Teams[team.ID] = team
}

// DeleteTeam removes a team from the database by their ID.
// Returns an error if the team does not exist.
func (d *Database) DeleteTeam(teamID int) error {
	if _, ok := d.Teams[teamID]; !ok {
		return fmt.Errorf("teamID %v does not exist", teamID)
	}
	delete(d.Teams, teamID)
	return nil
}

// GetTeam retrieves a team from the database by their ID.
// Returns an error if the team does not exist.
func (d *Database) GetTeam(teamID int) (*Team, error) {
	if _, ok := d.Teams[teamID]; !ok {
		return nil, fmt.Errorf("teamID %v does not exist", teamID)
	}
	return d.Teams[teamID], nil
}

// AddClub adds a new club to the database.
func (d *Database) AddClub(club *Club) {
	d.Clubs[club.ID] = club
}

// DeleteClub removes a club from the database by their ID.
// Returns an error if the club does not exist.
func (d *Database) DeleteClub(clubID int) error {
	if _, ok := d.Clubs[clubID]; !ok {
		return fmt.Errorf("clubID %v does not exist", clubID)
	}
	delete(d.Clubs, clubID)
	return nil
}

// GetClub retrieves a club from the database by their ID.
// Returns an error if the club does not exist.
func (d *Database) GetClub(clubID int) (*Club, error) {
	if _, ok := d.Clubs[clubID]; !ok {
		return nil, fmt.Errorf("clubID %v does not exist", clubID)
	}
	return d.Clubs[clubID], nil
}
