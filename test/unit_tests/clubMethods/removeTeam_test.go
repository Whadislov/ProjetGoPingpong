package clubmethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestRemoveTeam(t *testing.T) {
	c1 := mt.Club{
		ID:      0,
		Name:    "c1",
		TeamIDs: map[int]string{0: "t1"},
	}

	t1 := mt.Team{
		ID:   0,
		Name: "t1",
	}

	t2 := mt.Team{
		ID:   1,
		Name: "t2",
	}

	expectedLen1 := 0
	expectedError2 := "team t1 is not in club c2"

	t.Run("Remove team from player", func(t *testing.T) {
		err := c1.RemoveTeam(&t1)
		err2 := c1.RemoveTeam(&t2)
		if err != nil || len(c1.TeamIDs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
