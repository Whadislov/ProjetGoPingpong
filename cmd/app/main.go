package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"github.com/Whadislov/ProjetGoPingPong/api"
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {

	// Start app locally or on a browser ?
	appStartOption := "browser"

	// Start app on a browser
	if appStartOption == "browser" {
		// Load env variables
		err := godotenv.Load("credentials.env")
		if err != nil {
			log.Fatal("Cannot load variables from .env")
		}

		mdb.SetPsqlInfo(os.Getenv("WEB_DB_LINK"))
		mdb.SetDBName(os.Getenv("DB_NAME"))

		var wg sync.WaitGroup

		// API server (8001)
		apiReady := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			config := api.Config{
				ServerAddress: "localhost",
				ServerPort:    "8001",
			}
			// Start API
			go api.RunApi(&config)
			// Wait for the API to be ready
			time.Sleep(1 * time.Second)
			close(apiReady)
		}()

		// Web App (8000)
		wg.Add(1)

		go func() {
			<-apiReady
			defer wg.Done()
			log.Println("Starting app server on: 8000")
			err := http.ListenAndServe(":8000", http.FileServer(http.Dir("./wasm")))
			if err != nil {
				log.Fatalf("App server error: %v", err)
			}
		}()

		wg.Wait()

	}

	// Start app locally
	if appStartOption == "local" {
		// Load env variables
		err := godotenv.Load("credentials.env")
		if err != nil {
			log.Fatal("Cannot load variables from .env")
		}

		mdb.SetPsqlInfo(os.Getenv("LOCAL_DB_LINK"))
		mdb.SetDBName(os.Getenv("DB_NAME"))
		mu.Display(appStartOption)
	}
}
