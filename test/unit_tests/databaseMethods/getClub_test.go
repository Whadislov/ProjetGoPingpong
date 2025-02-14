package databasemethods_test

import (
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"testing"
)

func TestGetClub(t *testing.T) {
	d := mt.Database{
		Clubs: map[int]*mt.Club{0: {
			ID:   0,
			Name: "c",
		},
		},
	}
	expectedClub := mt.Club{
		ID:   0,
		Name: "c",
	}

	expectedError := "clubID 1 does not exist"

	t.Run("Get club from club ID", func(t *testing.T) {
		c, err := d.GetClub(0)
		_, err2 := d.GetClub(1)
		if c == nil {
			t.Errorf("Expected club %v, got %v", expectedClub, c)
		}
		if err != nil {
			t.Errorf("Expected err %v, got %v", nil, err)
		}
		if err2 == nil {
			t.Errorf("Expected err %v, got %v", expectedError, err2)
		}

	})

}
