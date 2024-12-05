package my_types

func (d *Database) AddClub(club *Club) {
	d.ClubList = append(d.ClubList, club)
}

func (d *Database) AddTeam(team *Team) {
	d.TeamList = append(d.TeamList, team)
}

func (d *Database) AddPlayer(player *Player) {
	d.PlayerList = append(d.PlayerList, player)
}
