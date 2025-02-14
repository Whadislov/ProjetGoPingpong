package teammethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestRemoveClub(t *testing.T) {
	t1 := mt.Team{
		ID:     0,
		Name:   "t1",
		ClubID: map[int]string{0: "c1"},
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
