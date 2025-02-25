package databasemethods_test

import (
	"fmt"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestDeleteTeam(t *testing.T) {
	tID := uuid.New()
	t2ID := uuid.New()

	d := mt.Database{
		Clubs: map[uuid.UUID]*mt.Club{},
		Teams: map[uuid.UUID]*mt.Team{tID: {
			ID:   tID,
			Name: "c",
		}},
		Players: map[uuid.UUID]*mt.Player{},
	}

	expectedLen := 0
	expectedError := fmt.Sprintf("teamID %v does not exist", t2ID)

	t.Run("Delete team from database", func(t *testing.T) {
		err := d.DeleteTeam(tID)
		err2 := d.DeleteTeam(t2ID)
		lenToVerify := len(d.Teams)
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
