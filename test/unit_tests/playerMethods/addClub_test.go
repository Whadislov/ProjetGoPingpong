package playermethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddClub(t *testing.T) {
	p1 := mt.Player{
		ID:        0,
		Firstname: "p1",
		ClubIDs:   map[int]string{0: "c1"},
	}

	c1 := mt.Club{
		ID:   0,
		Name: "c1",
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	expectedLen := 2
	expectedError := "player p1 is already in club c2"

	t.Run("Add club to player", func(t *testing.T) {
		err := p1.AddClub(&c1)
		err2 := p1.AddClub(&c2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(p1.ClubIDs) != expectedLen {
			t.Errorf("Expected len of ClubIDs %v, got %v", expectedLen, len(p1.ClubIDs))
		}
	})
}
