package my_database

import (
	"encoding/json"
	"fmt"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"os"
)

func SaveDb(filename string, db *mt.Database) error {
	// create or open the file and appends new content
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error during creation of database : %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // make JSON prettier
	err = encoder.Encode(db)
	if err != nil {
		return fmt.Errorf("error during JSON encoding : %w", err)
	}

	return nil
}
