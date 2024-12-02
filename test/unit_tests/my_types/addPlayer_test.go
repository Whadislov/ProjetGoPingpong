package my_types_test

import (
	"testing"
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func TestAddPlayer(t *testing.T) {

	var club my_types.Club
	club.Name = "TSG Heilbronn"
	club.PlayerList = []*my_types.Player{}

	var julien my_types.Player
	julien.Name = "Julien"
	expectedPlayerList  := []*my_types.Player{&julien}


	t.Run(fmt.Sprintf("Add player to player list of club %s", club.Name), func(t *testing.T) {
		club.AddPlayer(&julien)
	for i := range club.PlayerList {
		if club.PlayerList[i] != &julien {
			t.Errorf("Player list of %s is currently %v and is expected to be %v", club.Name, club.PlayerList[i], expectedPlayerList[i])
		} else {
			fmt.Printf("Player list of %s is currently %v and is expected to be %v", club.Name, club.PlayerList[i], expectedPlayerList[i])
		}
	}
	})
}