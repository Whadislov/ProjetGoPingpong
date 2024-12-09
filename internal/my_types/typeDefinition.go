package my_types

type Club struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"player_id_map"`
	TeamIDs   map[int]string `json:"team_id_map"`
}

type Player struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Ranking  int            `json:"ranking"`
	Material []string       `json:"material"`
	TeamIDs  map[int]string `json:"team_id_map"`
	ClubIDs  map[int]string `json:"club_id_map"`
}

type Team struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"player_id_map"`
	ClubID    map[int]string `json:"club_id_map"`
}

type Database struct {
	Clubs   map[int]*Club   `json:"club_map"`
	Teams   map[int]*Team   `json:"team_map"`
	Players map[int]*Player `json:"player_map"`
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

// ajouter erreur si la value de la clé (joueur) ne fait pas partie de la mape des joueurs
func (t *TeamMatch) SetTeamComposition(teamComposition map[int]string) {
	t.TeamComposition = teamComposition
}
*/
