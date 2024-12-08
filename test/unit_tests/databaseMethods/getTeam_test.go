package databasemethods_test

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestGetTeam(t *testing.T) {
	d := mt.Database{
		Teams: map[int]*mt.Team{0: {
			ID:   0,
			Name: "t",
		},
		},
	}
	expectedTeam := mt.Team{
		ID:   0,
		Name: "t",
	}

	expectedError := "teamID 1 does not exist"

	t.Run("Get team from team ID", func(t *testing.T) {
		team, err := d.GetTeam(0)
		_, err2 := d.GetTeam(1)
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
