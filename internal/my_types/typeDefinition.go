package my_types

type Club struct {
	Name       string
	PlayerList []*Player
	TeamList   []*Team //Teams map[string]*Team
}

type Player struct {
	Id       int
	Name     string
	Age      int
	Ranking  int
	Material []string
	TeamList []*Team
}

type Team struct {
	Name       string
	PlayerList []*Player
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
