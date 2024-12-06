package my_types

type Club struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"player_id_list"`
	TeamIDs   map[int]string `json:"team_id_list"`
	// map[player.ID] = player.Name, nil per default
	// map[team.ID] = team.Name, nil per default
}

type Player struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Ranking  int            `json:"ranking"`
	Material []string       `json:"material"`
	TeamIDs  map[int]string `json:"team_id_list"`
	ClubIDs  map[int]string `json:"club_id"`
	// map[team.ID] = team.Name, nil per default
	// map[club.ID] = club.Name, nil per default
}

type Team struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"player_id_list"`
	ClubID    map[int]string `json:"club_id"`
	// map[player.ID] = player.Name, nil per default
	// map[club.ID] = club.Name, nil per default
}

type Database struct {
	Clubs   map[int]*Club   `json:"club_list"`
	Teams   map[int]*Team   `json:"team_list"`
	Players map[int]*Player `json:"player_list"`
	//  map[club.ID] = *Club
	//  map[team.ID] = *Team
	//  map[player.ID] = *Player
}

/*
type PlayerMatch struct {
	League             string
	Player             string
	Adversary          string
	PlayerMatchOutcome Outcome
}

type Outcome struct {
	Victory string
	Defeat  string
	Draw    string
}

type TeamMatch struct {
	League          string
	TeamComposition map[int]string
}

type Match struct {
	HomeTeam         Team
	GuestTeam        Team
	TeamMatchOutcome Outcome
}

// ajouter erreur si la value de la clé (joueur) ne fait pas partie de la liste des joueurs
func (t *TeamMatch) SetTeamComposition(teamComposition map[int]string) {
	t.TeamComposition = teamComposition
}
*/
