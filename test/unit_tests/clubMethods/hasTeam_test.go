package clubmethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestHasTeam(t *testing.T) {
	c1 := mt.Club{
		ID:      uuid.New(),
		Name:    "c1",
		TeamIDs: map[uuid.UUID]string{uuid.New(): "t1"},
	}

	c2 := mt.Club{
		ID:      uuid.New(),
		Name:    "c2",
		TeamIDs: map[uuid.UUID]string{},
	}

	expectedBool1 := true
	expectedBool2 := false

	t.Run("Has club a team", func(t *testing.T) {
		bool1 := c1.HasTeam()
		bool2 := c2.HasTeam()
		if bool1 != expectedBool1 {
			t.Errorf("Expected bool %v, got %v", expectedBool1, bool1)
		}
		if bool2 != expectedBool2 {
			t.Errorf("Expected bool %v, got %v", expectedBool2, bool2)
		}
	})
}
