package myfunctions_test

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestNewTeam(t *testing.T) {
	d := mt.Database{
		Teams: map[int]*mt.Team{},
	}

	expectedLen := 1
	expectedError1 := "Team name cannot be empty"

	t.Run("Delete Team", func(t *testing.T) {
		_, err1 := mf.NewTeam("", &d)
		_, err2 := mf.NewTeam("p2", &d)

		if err1 == nil {
			t.Errorf("Expected error %v, got %v", expectedError1, err1)
		}
		if err2 != nil || len(d.Teams) != expectedLen {
			t.Errorf("Expected error %v, got %v", nil, err2)
		}
	})
}
