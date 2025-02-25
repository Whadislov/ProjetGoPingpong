package myfunctions_test

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	d := mt.Database{
		Players: map[uuid.UUID]*mt.Player{},
	}

	expectedLen := 1
	expectedError1 := "player name cannot be empty"
	expectedError2 := "player name can only contain letters"
	expectedError3 := "string is too long"

	t.Run("Delete Player", func(t *testing.T) {
		_, err1 := mf.NewPlayer("", "", &d)
		_, err2 := mf.NewPlayer("firstname2", "lastname2", &d)
		_, err3 := mf.NewPlayer("firstnameeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "lastname", &d)
		_, err4 := mf.NewPlayer("firstname", "lastname", &d)

		if err1 == nil {
			t.Errorf("Expected error %v, got %v", expectedError1, err1)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
		if err3 == nil {
			t.Errorf("Expected error %v, got %v", expectedError3, err3)
		}
		if err4 != nil || len(d.Players) != expectedLen {
			t.Errorf("Expected error %v, got %v", nil, err4)
		}
	})
}
