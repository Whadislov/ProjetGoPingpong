package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
	"testing"
)

func TestGetTeam(t *testing.T) {
	tID := uuid.New()
	t2ID := uuid.New()

	d := mt.Database{
		Teams: map[uuid.UUID]*mt.Team{tID: {
			ID:   tID,
			Name: "t",
		},
		},
	}
	expectedTeam := mt.Team{
		ID:   tID,
		Name: "t",
	}

	expectedError := "teamID 1 does not exist"

	t.Run("Get team from team ID", func(t *testing.T) {
		team, err := d.GetTeam(tID)
		_, err2 := d.GetTeam(t2ID)
		if team == nil {
			t.Errorf("Expected team %v, got %v", expectedTeam, team)
		}
		if err != nil {
			t.Errorf("Expected err %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected err %v, got %v", expectedError, err2)
		}

	})

}
