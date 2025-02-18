package playermethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestAddClub(t *testing.T) {
	c1 := mt.Club{
		ID:   uuid.New(),
		Name: "c1",
	}

	c2 := mt.Club{
		ID:   uuid.New(),
		Name: "c2",
	}

	p1 := mt.Player{
		ID:        uuid.New(),
		Firstname: "p1",
		ClubIDs:   map[uuid.UUID]string{c2.ID: c2.Name},
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
