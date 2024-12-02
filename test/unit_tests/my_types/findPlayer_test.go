package my_types_test

import (
	"testing"
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)
func TestFindPlayer(t *testing.T) {
	type testCase struct {
		club 			my_types.Club
		playerName      string
		expectedIndex 	int
		expectedError 	error
	}
	tests := []testCase{
		{
			club: my_types.Club{
				Name: "TSG Heilbronn",
				PlayerList: []*my_types.Player{
					{Name: "Julien"},
					{Name: "Lasse"},
				},
			},
			playerName:		"Julien",
			expectedIndex: 	0,
			expectedError: 	nil,
		},
		{
			club: my_types.Club{
				Name: "TSG Heilbronn",
				TeamList: []*my_types.Team{
					{Name: "Julien"},
					{Name: "Lasse"},
				},
			},
			playerName:		"Dominik",
			expectedIndex: 	-1,
			expectedError: 	fmt.Errorf("Dominik not found in the club"),
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		t.Run(fmt.Sprintf("Finding player %s in club %s", test.playerName, test.club.Name), func(t *testing.T) {
			index, err := test.club.FindPlayer(test.playerName)
			
		
			// Verify index
			if index != test.expectedIndex {
				//t.Errorf("Expected index: %d, got: %d", test.expectedIndex, index)
				t.Errorf(`-------------------------
				Inputs:     %v
				Expecting:  (%d)
				Actual:     (%d)
				Fail`, test.club, test.expectedIndex, index)
				failCount++
			} else {
				fmt.Printf(`-------------------------
				Inputs:     %v
				Expecting:  (%d)
				Actual:     (%d)
				Pass
				`, test.club, test.expectedIndex, index)
				passCount++
			}

			// Verify Error
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
