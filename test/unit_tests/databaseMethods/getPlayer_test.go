package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	d := mt.Database{
		Players: map[int]*mt.Player{0: {
			ID:   0,
			Name: "p",
		},
		},
	}
	expectedPlayer := mt.Player{
		ID:   0,
		Name: "p",
	}

	expectedError := "playerID 1 does not exist"

	t.Run("Get player from player ID", func(t *testing.T) {
		p, err := d.GetPlayer(0)
		_, err2 := d.GetPlayer(1)
		if p == nil {
			t.Errorf("Expected player %v, got %v", expectedPlayer, p)
		}
		if err != nil {
			t.Errorf("Expected err %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected err %v, got %v", expectedError, err2)
		}

	})

}
