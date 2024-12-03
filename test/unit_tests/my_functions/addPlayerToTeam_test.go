package my_functions_test

import (
	"testing"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestAddPlayerToTeam(t *testing.T) {
	p1 := mt.Player{
		Name: "p1",
		TeamList: []*mt.Team{},
	}

	t1 := mt.Team{
		Name: "t1",
		PlayerList: []*mt.Player{},
	}

	c1 := mt.Club{
		Name: "c1",
		PlayerList: []*mt.Player{&p1},
		TeamList: []*mt.Team{&t1},
	}

	err := mf.AddPlayerToTeam(&p1, t1.Name, &c1)
	err2 := mf.AddPlayerToTeam(&p1, "NonExistantTeam", &c1)


	t.Run("Add a player to a team", func(t *testing.T){
		if p1.TeamList[0] != &t1 {
			t.Errorf("Player's team list issue: got %v, expected %v", p1.TeamList[0], &t1)
		}
		if t1.PlayerList[0] != &p1 {
			t.Errorf("Team's player list issue: got %v, expected %v", t1.PlayerList[0], &p1)
		}
		if err != nil {
			t.Errorf("Error issue: got %v, expected %v", err, nil)
		}
		if err2 == nil {
			t.Errorf("Error issue: got %v, expected error", err2)
		}
	})










}

/*
func AddPlayerToTeam(p *mt.Player, teamName string, c *mt.Club) (error) {
	if p == nil {
		return fmt.Errorf("player is nil")
	}

	_, err := c.FindTeam(teamName)
	if err != nil {
		return fmt.Errorf("team %s not found: %w", teamName, err)
	}

	err = c.AddPlayerToTeam(p, teamName)
	if  err == nil {
		log.Printf("%s has been successfully added in %s", p.Name, teamName)
	} else {
		return fmt.Errorf("%s has not been successfully added in %s : %w", p.Name, teamName, err)
	}
	return nil
}

*/