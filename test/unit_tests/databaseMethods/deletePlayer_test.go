package databasemethods_test

import (
	"fmt"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestDeletePlayer(t *testing.T) {
	pID := uuid.New()
	p2ID := uuid.New()

	d := mt.Database{
		Clubs: map[uuid.UUID]*mt.Club{},
		Teams: map[uuid.UUID]*mt.Team{},
		Players: map[uuid.UUID]*mt.Player{pID: {
			ID:        pID,
			Firstname: "c",
		}},
	}

	expectedLen := 0
	expectedError := fmt.Sprintf("playerID %v does not exist", p2ID)

	t.Run("Delete player from database", func(t *testing.T) {
		err := d.DeletePlayer(pID)
		err2 := d.DeletePlayer(p2ID)
		lenToVerify := len(d.Players)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Players %v, got %v", expectedLen, lenToVerify)
		}
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err2)
		}
	})
}
