package myfunctions_test

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestDeleteTeam(t *testing.T) {

	t1 := mt.Team{
		ID:   uuid.New(),
		Name: "t1",
	}

	t2 := mt.Team{
		ID:   uuid.New(),
		Name: "t2",
	}

	d := mt.Database{
		Teams:           map[uuid.UUID]*mt.Team{t1.ID: &t1},
		DeletedElements: map[string][]uuid.UUID{},
	}

	expectedLen1 := 0
	expectedError2 := "error when deleting Team t2: TeamID 1 does not exist"

	t.Run("Delete Team", func(t *testing.T) {
		err := mf.DeleteTeam(&t1, &d)
		err2 := mf.DeleteTeam(&t2, &d)
		if err != nil || len(d.Teams) != expectedLen1 {
			t.Errorf("Expected error %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected error %v, got %v", expectedError2, err2)
		}
	})
}
