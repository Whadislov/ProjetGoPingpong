package playermethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestHasClub(t *testing.T) {
	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		ClubIDs:   map[int]string{0: "c1"},
	}

	p2 := mt.Player{
		ID:        1,
		Firstname: "p2",
	}

	expectedBool1 := true
	expectedBool2 := false

	t.Run("Has player a club", func(t *testing.T) {
		bool1 := p1.HasClub()
		bool2 := p2.HasClub()
		if bool1 != expectedBool1 {
			t.Errorf("Expected bool %v, got %v", expectedBool1, bool1)
		}
		if bool2 != expectedBool2 {
			t.Errorf("Expected bool %v, got %v", expectedBool2, bool2)
		}
	})
}
