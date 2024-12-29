package main

import (
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/api"
	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	// Load the database (deserialize)
	golangDB, err := msql.LoadDB()
	if err != nil {
		fmt.Println("Error while loading golang database:", err)
		return
	}

	fmt.Println("Api is running")
	api.RunApi(golangDB)
}
