package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeleteClub(t *testing.T) {

	c1 := mt.Club{
		ID:   0,
		Name: "c1",
	}

	c2 := mt.Club{
		ID:   1,
		Name: "c2",
	}

	d := mt.Database{
		Clubs: map[int]*mt.Club{0: &c1},
	}

	expectedLen1 := 0
	expectedError2 := "error when deleting Club c2: ClubID 1 does not exist"

	t.Run("Delete Club", func(t *testing.T) {
		err := mf.DeleteClub(&c1, &d)
		err2 := mf.DeleteClub(&c2, &d)
		if err != nil || len(d.Clubs) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
