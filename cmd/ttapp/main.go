package main

import (
	//"sync"
	//"time"

	//"github.com/Whadislov/ProjetGoPingPong/api"
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	//mf "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	// Let know when the API is ready
	//var wg sync.WaitGroup
	//wg.Add(1)

	// Start app locally or on a browser ?
	//appStartOption := "browser"
	appStartOption := "browser"

	// Load the configuration file
	/*
		config, err := api.LoadConfig("config.json")
		if err != nil {
			log.Printf("Error loading config: %v", err)
		}
	*/

	// Start app on a browser
	if appStartOption == "browser" {
		mdb.AppStartOption(appStartOption)
		mu.AppStartOption(appStartOption)
		// Load the database (deserialize)
		/*
			config := &api.Config{
				ServerAddress: "localhost",
				ServerPort:    "7000",
			}

			go func() {

				go api.RunApi(config)
				time.Sleep(1 * time.Second)
				if mf.IsApiReady() {
					wg.Done()
				} else {
					panic("Issue with API")
				}

			}()

			// Wait for the API to be ready
			wg.Wait()
		*/
		/*
			golangDB, err := mf.LoadDB()
			if err != nil {
				panic(err)
			}
		*/

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

		mu.Display(&golangDB)
	}

	// Start app locally
	if appStartOption == "local" {
		mdb.AppStartOption(appStartOption)
		mu.AppStartOption(appStartOption)
		// Load the database (deserialize)
		golangDB, err := mdb.LoadDB()
		if err != nil {
			panic(err)
		}
		mu.Display(golangDB)
	}
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
