package my_types_test

import (
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestFindTeam(t *testing.T) {
	type testCase struct {
		club          my_types.Club
		team          my_types.Team
		expectedError error
	}
	team1 := my_types.Team{
		Name: "Mannschaft 1",
	}
	team0 := my_types.Team{
		Name: "Mannschaft 0",
	}

	tests := []testCase{
		{
			club: my_types.Club{
				Name:     "TSG Heilbronn",
				TeamList: []*my_types.Team{&team1},
			},
			team:          team1,
			expectedError: nil,
		},
		{
			club: my_types.Club{
				Name:     "TSG Heilbronn",
				TeamList: []*my_types.Team{&team1},
			},
			team:          team0,
			expectedError: fmt.Errorf("Mannschaft 0 not found in TSG Heilbronn"),
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		t.Run(fmt.Sprintf("Finding team %s in club %s", test.team.Name, test.club.Name), func(t *testing.T) {
			err := test.club.FindTeam(&test.team)

			if (err == nil && test.expectedError != nil) || (err != nil && test.expectedError == nil) ||
				(err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error()) {
				//t.Errorf("Expected error: %v, got: %v", test.expectedError, err)
				//}
				t.Errorf(`-------------------------
				Inputs:     %v
				Expecting:  (%v)
				Actual:     (%v)
				Fail`, test.club, test.expectedError, err)
				failCount++
			} else {
				fmt.Printf(`-------------------------
				Inputs:     %v
				Expecting:  (%v)
				Actual:     (%v)
				Pass
				`, test.club, test.expectedError, err)
				passCount++
			}
		})
		fmt.Println("---------------------------------")
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}
