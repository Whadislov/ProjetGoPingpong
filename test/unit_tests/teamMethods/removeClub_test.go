package teammethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestRemoveClub(t *testing.T) {
	c1 := mt.Club{
		ID:   uuid.New(),
		Name: "c1",
	}

	t1 := mt.Team{
		ID:     uuid.New(),
		Name:   "t1",
		ClubID: map[uuid.UUID]string{c1.ID: c1.Name},
	}

	c2 := mt.Club{
		ID:   uuid.New(),
		Name: "c2",
	}

	expectedLen1 := 0
	expectedError2 := "team t1 is not in club c2"

	t.Run("Remove team from player", func(t *testing.T) {
		err := t1.RemoveClub(&c1)
		err2 := t1.RemoveClub(&c2)
		if err != nil || len(t1.ClubID) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
