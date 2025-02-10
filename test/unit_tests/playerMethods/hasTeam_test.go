package playermethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestHasTeam(t *testing.T) {
	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		TeamIDs:   map[int]string{0: "t1"},
	}

	p2 := mt.Player{
		ID:        1,
		Firstname: "p2",
	}

	expectedBool1 := true
	expectedBool2 := false

	t.Run("Has player a team", func(t *testing.T) {
		bool1 := p1.HasTeam()
		bool2 := p2.HasTeam()
		if bool1 != expectedBool1 {
			t.Errorf("Expected bool %v, got %v", expectedBool1, bool1)
		}
		if bool2 != expectedBool2 {
			t.Errorf("Expected bool %v, got %v", expectedBool2, bool2)
		}
	})
}
