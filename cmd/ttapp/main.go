package main

import (
	"fmt"
	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {

	// Load the database (deserialize)
	golangDB, err := msql.LoadDB()
	if err != nil {
		fmt.Println("Error while loading golang database:", err)
		return
	}
	mu.Display(golangDB)

}
