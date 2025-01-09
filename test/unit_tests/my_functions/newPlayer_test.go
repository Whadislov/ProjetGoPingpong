package myfunctions_test

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	d := mt.Database{
		Players: map[int]*mt.Player{},
	}

	expectedLen := 1
	expectedError1 := "player name cannot be empty"
	expectedError2 := "player name can only contain letters"

	t.Run("Delete Player", func(t *testing.T) {
		_, err1 := mf.NewPlayer("", &d)
		_, err2 := mf.NewPlayer("p2", &d)
		_, err3 := mf.NewPlayer("p", &d)

		if err1 == nil {
			t.Errorf("Expected error %v, got %v", expectedError1, err1)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
		if err3 != nil || len(d.Players) != expectedLen {
			t.Errorf("Expected error %v, got %v", nil, err2)
		}
	})
}
