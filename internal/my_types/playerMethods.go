package my_types

import "fmt"

/*
type Player struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Ranking  int            `json:"ranking"`
	Material []string       `json:"material"`
	TeamIDs  map[int]string `json:"team_id_list"`
	ClubID   int            `json:"club_id"`
	// map[team.ID] = team.Name, nil per default
	// ClubID == -1 per default
}
*/

func (p *Player) SetPlayerID(id int) {
	p.ID = id
}

func (p *Player) SetPlayerName(name string) {
	p.Name = name
}

func (p *Player) SetPlayerAge(age int) {
	p.Age = age
}

func (p *Player) SetPlayerRanking(ranking int) {
	p.Ranking = ranking
}

func (p *Player) SetPlayerMaterial(forehand string, backhand string, blade string) {
	material := []string{forehand, backhand, blade}
	p.Material = material
}

func (p *Player) AddTeam(team *Team) error {
	if _, ok := p.TeamIDs[team.ID]; ok {
		return fmt.Errorf("player %v is already in team %v.", p.Name, team.Name)
	}
	p.TeamIDs[team.ID] = team.Name
	return nil
}

func (p *Player) AddClub(club *Club) error {
	if _, ok := p.ClubIDs[club.ID]; ok {
		return fmt.Errorf("player %v is already in club %v.", p.Name, club.Name)
	}
	p.ClubIDs[club.ID] = club.Name
	return nil
}

func (p *Player) RemoveTeam(team *Team) error {
	if _, ok := p.TeamIDs[team.ID]; !ok {
		return fmt.Errorf("player %v is not in team %v.", p.Name, team.Name)
	}
	delete(p.TeamIDs, team.ID)
	return nil
}

func (p *Player) RemoveClub(club *Club) error {
	if _, ok := p.ClubIDs[club.ID]; !ok {
		return fmt.Errorf("player %v is not in club %v.", p.Name, club.Name)
	}
	delete(p.ClubIDs, club.ID)
	return nil
}

func (p *Player) HasTeam() bool {
	if len(p.TeamIDs) > 0 {
		return false
	}
	return true
}

func (p *Player) HasClub() bool {
	if len(p.ClubIDs) > 0 {
		return false
	}
	return true
}
