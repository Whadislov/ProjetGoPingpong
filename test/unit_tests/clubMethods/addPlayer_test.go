package clubmethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	c1 := mt.Club{
		ID:        0,
		Name:      "c1",
		PlayerIDs: map[int]string{0: "p1"},
	}

	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
	}

	p2 := mt.Player{
		ID:        1,
		Firstname: "p2",
	}

	expectedLen := 2
	expectedError := "player p1 is already in club c1"

	t.Run("Add player to club", func(t *testing.T) {
		err := c1.AddPlayer(&p1)
		err2 := c1.AddPlayer(&p2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(c1.PlayerIDs) != expectedLen {
			t.Errorf("Expected len of PlayerIDs %v, got %v", expectedLen, len(c1.PlayerIDs))
		}
	})
}
