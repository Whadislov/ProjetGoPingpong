package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestDeletePlayer(t *testing.T) {

	p1 := mt.Player{
		ID:   0,
		Name: "p1",
	}

	p2 := mt.Player{
		ID:   1,
		Name: "p2",
	}

	d := mt.Database{
		Players: map[int]*mt.Player{0: &p1},
	}

	expectedLen1 := 0
	expectedError2 := "error when deleting player p2: playerID 1 does not exist"

	t.Run("Delete Player", func(t *testing.T) {
		err := mf.DeletePlayer(&p1, &d)
		err2 := mf.DeletePlayer(&p2, &d)
		if err != nil || len(d.Players) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
