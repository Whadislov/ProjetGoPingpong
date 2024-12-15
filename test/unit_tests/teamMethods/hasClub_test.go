package teammethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestHasClub(t *testing.T) {
	t1 := mt.Team{
		ID:     0,
		Name:   "t1",
		ClubID: map[int]string{0: "c1"},
	}

	t2 := mt.Team{
		ID:     1,
		Name:   "t2",
		ClubID: map[int]string{},
	}

	expectedBool1 := true
	expectedBool2 := false

	t.Run("Has team a club", func(t *testing.T) {
		bool1 := t1.HasClub()
		bool2 := t2.HasClub()
		if bool1 != expectedBool1 {
			t.Errorf("Expected bool %v, got %v", expectedBool1, bool1)
		}
		if bool2 != expectedBool2 {
			t.Errorf("Expected bool %v, got %v", expectedBool2, bool2)
		}
	})
}