package my_functions

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// NewPlayer creates a new player with the given name and adds it to the database.
// Returns the created player and an error if the player name is empty or if there is an issue with the operation.
func NewPlayer(playerName string, db *mt.Database) (*mt.Player, error) {
	if playerName == "" {
		return nil, fmt.Errorf("player name cannot be empty")
	}

	for _, r := range playerName {
		if r < 'A' || r > 'z' {
			return nil, fmt.Errorf("player name can only contain letters")
		}
	}

	playerName = strings.ToLower(playerName)
	firstRune, size := utf8.DecodeRuneInString(playerName)
	if firstRune != utf8.RuneError {
		playerName = string(unicode.ToUpper(firstRune)) + playerName[size:]
	}

	p := &mt.Player{
		ID:       len(db.Players),
		Name:     playerName,
		Age:      -1,
		Ranking:  -1,
		Material: DefaultPlayerMaterial(),
		TeamIDs:  make(map[int]string),
		ClubIDs:  make(map[int]string),
	}

	db.AddPlayer(p)
	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}
