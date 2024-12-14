package clubmethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestHasPlayer(t *testing.T) {
	c1 := mt.Club{
		ID:        0,
		Name:      "c1",
		PlayerIDs: map[int]string{0: "p1"},
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	expectedBool1 := true
	expectedBool2 := false

	t.Run("Has club a player", func(t *testing.T) {
		bool1 := c1.HasPlayer()
		bool2 := c2.HasPlayer()
		if bool1 != expectedBool1 {
			t.Errorf("Expected bool %v, got %v", expectedBool1, bool1)
		}
		if bool2 != expectedBool2 {
			t.Errorf("Expected bool %v, got %v", expectedBool2, bool2)
		}
	})
}
