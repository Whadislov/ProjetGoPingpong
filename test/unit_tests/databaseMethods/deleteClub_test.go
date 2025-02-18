package databasemethods_test

import (
	"fmt"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestDeleteClub(t *testing.T) {
	cID := uuid.New()
	c2ID := uuid.New()

	d := mt.Database{
		Clubs: map[uuid.UUID]*mt.Club{cID: {
			ID:   cID,
			Name: "c",
		},
		},
		Teams:   map[uuid.UUID]*mt.Team{},
		Players: map[uuid.UUID]*mt.Player{},
	}

	expectedLen := 0
	expectedError := fmt.Sprintf("clubID %v does not exist", c2ID)

	t.Run("Delete club from database", func(t *testing.T) {
		err := d.DeleteClub(cID)
		err2 := d.DeleteClub(c2ID)
		lenToVerify := len(d.Clubs)
		if lenToVerify != expectedLen {
			t.Errorf("Expected len of Clubs %v, got %v", expectedLen, lenToVerify)
		}
		if err != nil {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError, err2)
		}
	})
}
