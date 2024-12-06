package my_database

import (
	"encoding/json"
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"os"
)

func LoadDb(filename string) (*mt.Database, error) {

	// Check if file exists
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Return empty database if the file does not exist
			return &mt.Database{}, nil
		}
		return nil, fmt.Errorf("error while opening database : %w", err)
	}
	defer file.Close()

	db := &mt.Database{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(db)

	if err != nil {
		return nil, fmt.Errorf("error while JSON decoding : %w", err)
	}

	return db, nil
}
