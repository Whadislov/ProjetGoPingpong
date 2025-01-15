package main

import (
	//"log"

	"github.com/Whadislov/ProjetGoPingPong/api"
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {

	// Load the configuration file
	/*
		config, err := api.LoadConfig("config.json")
		if err != nil {
			log.Printf("Error loading config: %v", err)
		}
	*/

	config := &api.Config{
		ServerAddress: "localhost",
		ServerPort:    "7000",
	}

	// Load the database (deserialize)
	golangDB, err := mdb.LoadDB()
	if err != nil {
		panic(err)
	}

	go api.RunApi(golangDB, config)
	mu.Display(golangDB)
}

/*
// Debug for fyne serve --port 8000 --sourceDir cmd/TTapp

import (
	"log"

	"github.com/Whadislov/ProjetGoPingPong/api"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)


func main() {

	// Load the configuration file
	config, err := api.LoadConfig("config.json")
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	golangDB := mt.Database{
		Players: map[int]*mt.Player{
			0: {ID: 0, Name: "Julien", Age: 27, Ranking: 1632, Material: []string{"Forehand", "Backhand", "Blade"}},
		},
		Teams: map[int]*mt.Team{
			0: {ID: 0, Name: "Mannschaft 2"},
		},
		Clubs: map[int]*mt.Club{
			0: {ID: 0, Name: "TSG Heilbronn"},
		},
	}

	go api.RunApi(&golangDB, config)
	mu.Display(&golangDB)
}
*/
