package clubmethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestRemovePlayer(t *testing.T) {
	c1 := mt.Club{
		ID:        0,
		Name:      "c1",
		PlayerIDs: map[int]string{0: "p1"},
	}

	p1 := mt.Player{
		ID:   0,
		Name: "p1",
	}

	p2 := mt.Player{
		ID:   1,
		Name: "p2",
	}

	expectedLen1 := 0
	expectedError2 := "player p1 is not in club c2"

	t.Run("Remove team from player", func(t *testing.T) {
		err := c1.RemovePlayer(&p1)
		err2 := c1.RemovePlayer(&p2)
		if err != nil || len(c1.PlayerIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
