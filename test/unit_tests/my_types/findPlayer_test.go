package my_types_test

import (
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestFindPlayer(t *testing.T) {
	type testCase struct {
		club          my_types.Club
		player        my_types.Player
		expectedError error
	}
	player1 := my_types.Player{
		Name: "Julien",
	}
	player2 := my_types.Player{
		Name: "Lasse",
	}
	player3 := my_types.Player{
		Name: "Dominik",
	}

	tests := []testCase{
		{
			club: my_types.Club{
				Name: "TSG Heilbronn",
				PlayerList: []*my_types.Player{
					&player1,
					&player2,
				},
			},
			player:        player1,
			expectedError: nil,
		},
		{
			club: my_types.Club{
				Name: "TSG Heilbronn",
				TeamList: []*my_types.Team{
					{Name: "Julien"},
					{Name: "Lasse"},
				},
			},
			player:        player3,
			expectedError: fmt.Errorf("Dominik not found in TSG Heilbronn"),
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		t.Run(fmt.Sprintf("Finding player %s in club %s", test.player.Name, test.club.Name), func(t *testing.T) {
			err := test.club.FindPlayer(&test.player)

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
