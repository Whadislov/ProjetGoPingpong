package my_types

type Club struct {
	Name       string    `json:"name"`
	PlayerList []*Player `json:"player_list"`
	TeamList   []*Team   `json:"team_list"`
}

type Player struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Ranking  int      `json:"ranking"`
	Material []string `json:"material"`
	TeamList []*Team  `json:"team_list"`
}

type Team struct {
	Name       string    `json:"name"`
	PlayerList []*Player `json:"player_list"`
	Club       *Club     `json:"club"`
}

type Database struct {
	ClubList   []*Club   `json:"club_list"`
	TeamList   []*Team   `json:"team_list"`
	PlayerList []*Player `json:"player_list"`
}

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
