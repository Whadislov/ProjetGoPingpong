package playermethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestRemoveClub(t *testing.T) {
	p1 := mt.Player{
		ID:      0,
		Name:    "p1",
		ClubIDs: map[int]string{0: "c1"},
	}

	c1 := mt.Club{
		ID:   0,
		Name: "c1",
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	expectedLen1 := 0
	expectedError2 := "player p1 is not in club c2"

	t.Run("Remove club from player", func(t *testing.T) {
		err := p1.RemoveClub(&c1)
		err2 := p1.RemoveClub(&c2)
		if err != nil || len(p1.ClubIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
