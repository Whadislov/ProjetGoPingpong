package playermethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestRemoveTeam(t *testing.T) {
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

	expectedLen1 := 0
	expectedError2 := "player p1 is not in team t2"

	t.Run("Remove team from player", func(t *testing.T) {
		err := p1.RemoveTeam(&t1)
		err2 := p1.RemoveTeam(&t2)
		if err != nil || len(p1.TeamIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
