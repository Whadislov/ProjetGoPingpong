package my_functions_test

import (
	"fmt"
	"testing"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)


func TestNewClub(t *testing.T) {
	expectedClub:= mt.Club{
		Name: "club",
		PlayerList: []*mt.Player{},
		TeamList: []*mt.Team{},
	}

	expectedLen := 0

	club := mt.Club{
		Name: "club",
	}

	c, _ := mf.NewClub("club")
	_, err := mf.NewClub("")
	

	t.Run(fmt.Sprintf("Create club %s", club.Name), func(t *testing.T) {
		if c.Name != expectedClub.Name {
			t.Errorf("Name issue: got %v, expected %v", c.Name, expectedClub.Name)
		}
		if len(c.PlayerList) != expectedLen {
			t.Errorf("Playerlist len issue: got %v, expected %v", len(c.PlayerList), expectedLen)
		}

		if len(c.TeamList) != expectedLen {
			t.Errorf("TeamList len issue: got %v, expected %v", len(c.TeamList), expectedLen)
		}
		if err == nil {
			t.Errorf("Club name issue: got %v, expected %v", err, fmt.Errorf("team name cannot be empty"))
		}
	})
}