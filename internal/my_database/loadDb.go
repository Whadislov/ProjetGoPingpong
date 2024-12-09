package my_database

import (
	"encoding/json"
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"os"
)

func LoadDb(filename string) (*mt.Database, error) {
	emptyDb := mt.Database{
		Clubs:   map[int]*mt.Club{},
		Teams:   map[int]*mt.Team{},
		Players: map[int]*mt.Player{},
	}

	// Check if file exists
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Return empty database if the file does not exist and create the file
			os.Create(filename)
			return &emptyDb, nil
		}
		return nil, fmt.Errorf("error during database opening: %w", err)
	}
	defer file.Close()

	// Check if database is empty
	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	if stat.Size() == 0 {
		return &emptyDb, nil
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&emptyDb)

	if err != nil {
		return nil, fmt.Errorf("error during JSON decoding: %w", err)
	}

	return &emptyDb, nil
}
