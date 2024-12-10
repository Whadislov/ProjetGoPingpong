package my_types

type Club struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"players"`
	TeamIDs   map[int]string `json:"teams"`
}

type Player struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Ranking  int            `json:"ranking"`
	Material []string       `json:"material"`
	TeamIDs  map[int]string `json:"teams"`
	ClubIDs  map[int]string `json:"clubs"`
}

type Team struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	PlayerIDs map[int]string `json:"players"`
	ClubID    map[int]string `json:"clubs"`
}

type Database struct {
	Clubs   map[int]*Club   `json:"clubs"`
	Teams   map[int]*Team   `json:"teams"`
	Players map[int]*Player `json:"players"`
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
