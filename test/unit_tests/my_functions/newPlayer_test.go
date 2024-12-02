package my_functions_test

import (
	"fmt"
	"testing"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestNewPlayer(t *testing.T) {
	expectedPlayer := mt.Player{
		Name: "p1",
		Age: 0,
		Ranking: 0,
		Material: []string{"","",""},
		TeamList: []*mt.Team{},
	}

	club := mt.Club{
		Name: "club",
	}

	p1, err1 := mf.NewPlayer("p1", &club)
	_, err2 := mf.NewPlayer("p1", &club)

	expectedLen := 0

	t.Run(fmt.Sprintf("Add new player to club %s", club.Name), func(t *testing.T) {
		if p1.Name != expectedPlayer.Name {
			t.Errorf("Name issue: got %v, expected %v", p1.Name, expectedPlayer.Name)
		}
		if p1.Age != expectedPlayer.Age {
			t.Errorf("Age issue: got %v, expected %v", p1.Age, expectedPlayer.Age)
		}
		if p1.Ranking != expectedPlayer.Ranking {
			t.Errorf("Ranking issue: got %v, expected %v", p1.Ranking, expectedPlayer.Ranking)
		}
		if p1.Material[0] != "Unknown" {
			t.Errorf("Material issue: got %v, expected %v", p1.Material[0], "Unknown")
		}
		if p1.Material[1] != "Unknown" {
			t.Errorf("Material issue: got %v, expected %v", p1.Material[1], "Unknown")
		}
		if p1.Material[2] != "Unknown" {
			t.Errorf("Material issue: got %v, expected %v", p1.Material[2], "Unknown")
		}
		if len(p1.TeamList) != expectedLen {
			t.Errorf("TeamList issue: got %v, expected %v", len(p1.TeamList), expectedLen)
		}
		if err1 != nil {
			t.Errorf("Error 1 issue: got %v, expected %v", err1, nil)
		}
		if err2 == nil {
			t.Errorf("Error 2 issue: got %v, expected an error", err2)
		}
	})
}

/*

	if p1.Name != expectedPlayer.Name ||
	p1.Age != expectedPlayer.Age ||
	p1.Ranking != expectedPlayer.Ranking ||
	p1.Material[0] != "" ||
	p1.Material[1] != "" ||
	p1.Material[2] != "" ||
	len(p1.TeamList) != expectedLen ||
	err1 != nil ||
	err2 == nil {
		t.Errorf(`-----------------------------------
		Expecting:		(%v, %v, %v, "", "", "", %v, %v, %v)
		Actual:			(%v, %v, %v, %v, %v, %v, %v, %v, %v)
		Fail`, 
		expectedPlayer.Name, expectedPlayer.Age, expectedPlayer.Ranking, expectedLen, nil, nil,
		p1.Name, p1.Age, p1.Ranking, p1.Material[0], p1.Material[1], p1.Material[2], len(p1.TeamList), err1, err2)
	} else {
		fmt.Printf(`-----------------------------------
		Expecting:		(%v, %v, %v, "", "", "", %v, %v, %v)
		Actual:			(%v, %v, %v, %v, %v, %v, %v, %v, %v)
		Pass
		`, 
		expectedPlayer.Name, expectedPlayer.Age, expectedPlayer.Ranking, expectedLen, nil, nil,
		p1.Name, p1.Age, p1.Ranking, p1.Material[0], p1.Material[1], p1.Material[2], len(p1.TeamList), err1, err2)
	}
*/