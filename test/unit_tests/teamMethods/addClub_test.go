package teammethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestAddClub(t *testing.T) {
	t1 := mt.Team{
		ID:     0,
		Name:   "t1",
		ClubID: map[int]string{0: "c1"},
	}

	t2 := mt.Team{
		ID:     0,
		Name:   "t2",
		ClubID: map[int]string{},
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	expectedError := "team t1 is already in a club"
	expectedLen := 1

	t.Run("Add club to player", func(t *testing.T) {
		err := t1.AddClub(&c2)
		err2 := t2.AddClub(&c2)
		if err == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
		if err2 != nil || len(t2.ClubID) != expectedLen {
			t.Errorf("Expected len of ClubID %v, got %v", expectedLen, len(t2.ClubID))
		}
	})
}
