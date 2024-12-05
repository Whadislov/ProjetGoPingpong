package my_functions_test

import (
	"fmt"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestNewTeam(t *testing.T) {
	db := mt.Database{}

	club := mt.Club{
		Name: "club",
	}

	expectedTeam := mt.Team{
		Name:       "t1",
		PlayerList: []*mt.Player{},
		Club:       &club,
	}

	t1, err1 := mf.NewTeam("t1", &club, &db)
	_, err2 := mf.NewTeam("t1", &club, &db)
	_, err3 := mf.NewTeam("", &club, &db)

	expectedLen := 0

	t.Run(fmt.Sprintf("Add new player to club %s", club.Name), func(t *testing.T) {
		if t1.Name != expectedTeam.Name {
			t.Errorf("Name issue: got %v, expected %v", t1.Name, expectedTeam.Name)
		}
		if len(t1.PlayerList) != expectedLen {
			t.Errorf("Playerlist issue: got %v, expected %v", len(t1.PlayerList), expectedLen)
		}
		if err1 != nil {
			t.Errorf("Error 1 issue: got %v, expected %v", err1, nil)
		}
		if err2 == nil {
			t.Errorf("Error 2 issue: got %v, expected an error", err2)
		}
		if err3 == nil {
			t.Errorf("Team name issue: got %v, expected %v", err3, fmt.Errorf("team name cannot be empty"))
		}
	})
}
