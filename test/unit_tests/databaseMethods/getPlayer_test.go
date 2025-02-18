package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	pID := uuid.New()
	p2ID := uuid.New()

	d := mt.Database{
		Players: map[uuid.UUID]*mt.Player{pID: {
			ID:        pID,
			Firstname: "p",
		},
		},
	}
	expectedPlayer := mt.Player{
		ID:        pID,
		Firstname: "p",
	}

	expectedError := "playerID 1 does not exist"

	t.Run("Get player from player ID", func(t *testing.T) {
		p, err := d.GetPlayer(pID)
		_, err2 := d.GetPlayer(p2ID)
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
