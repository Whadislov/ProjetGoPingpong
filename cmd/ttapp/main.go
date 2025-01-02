package main

import (
	"fmt"
	"log"

	"github.com/Whadislov/ProjetGoPingPong/api"
	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	// Load the configuration file
	config, err := api.LoadConfig("config.json")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	// Load the database (deserialize)
	golangDB, err := msql.LoadDB()
	if err != nil {
		fmt.Println("Error while loading golang database:", err)
		return
	}
	go api.RunApi(golangDB, config)
	mu.Display(golangDB)

}
