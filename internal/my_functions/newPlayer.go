package my_functions

import (
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// NewPlayer creates a new player with the given name and adds it to the database.
// Returns the created player and an error if the player name is empty or if there is an issue with the operation.
func NewPlayer(playerName string, db *mt.Database) (*mt.Player, error) {
	b, err := IsValidName(playerName)
	if !b {
		return nil, err
	}

	playerName = strings.ToLower(playerName)
	firstRune, size := utf8.DecodeRuneInString(playerName)
	if firstRune != utf8.RuneError {
		playerName = string(unicode.ToUpper(firstRune)) + playerName[size:]
	}

	p := &mt.Player{
		ID:       NewPlayerCount,
		Name:     playerName,
		Age:      -1,
		Ranking:  -1,
		Material: DefaultPlayerMaterial(),
		TeamIDs:  make(map[int]string),
		ClubIDs:  make(map[int]string),
	}

	// Be ready for next player
	NewPlayerCount--
	db.AddPlayer(p)
	log.Printf("Player %v sucessfully created.", playerName)
	return p, nil
}
