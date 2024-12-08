package clubmethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestAddTeam(t *testing.T) {
	c1 := mt.Club{
		ID:      0,
		Name:    "c1",
		TeamIDs: map[int]string{0: "t1"},
	}

	t1 := mt.Team{
		ID:     0,
		Name:   "t1",
		ClubID: map[int]string{},
	}

	t2 := mt.Team{
		ID:   1,
		Name: "t2",
	}

	expectedError := "team t1 is already in club c1"
	expectedLen := 2

	t.Run("Add team to club", func(t *testing.T) {
		err := c1.AddTeam(&t1)
		err2 := c1.AddTeam(&t2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(c1.TeamIDs) != expectedLen {
			t.Errorf("Expected len of ClubID %v, got %v", expectedLen, len(c1.TeamIDs))
		}
	})
}
