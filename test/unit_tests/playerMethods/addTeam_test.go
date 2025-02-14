package playermethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddTeam(t *testing.T) {
	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		TeamIDs:   map[int]string{0: "t1"},
	}

	t1 := mt.Team{
		ID:   0,
		Name: "t1",
	}

	t2 := mt.Team{
		ID:   1,
		Name: "t2",
	}

	expectedLen := 2
	expectedError := "player p1 is already in team t2"

	t.Run("Add team to player", func(t *testing.T) {
		err := p1.AddTeam(&t1)
		err2 := p1.AddTeam(&t2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(p1.TeamIDs) != expectedLen {
			t.Errorf("Expected len of TeamIDs %v, got %v", expectedLen, len(p1.TeamIDs))
		}
	})
}
