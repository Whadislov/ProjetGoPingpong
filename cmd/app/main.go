package main

import (
	"encoding/json"
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

func loadConfig(filename string) (*api.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &api.Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

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
		api.SetJWTSecretKey(os.Getenv("JWT_SECRET_KEY"))

		var wg sync.WaitGroup

		// API server (8001)
		apiReady := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			configApi, err := loadConfig("config_api.json")
			if err != nil {
				log.Fatalf("Cannot read config file: %v", err)
			}
			// Start API
			go api.RunApi(configApi)
			// Wait for the API to be ready
			time.Sleep(1 * time.Second)
			close(apiReady)
		}()

		// Web App (8000)
		wg.Add(1)

		go func() {
			<-apiReady
			defer wg.Done()
			configApp, err := loadConfig("config_app.json")
			if err != nil {
				log.Fatalf("Cannot read config file: %v", err)
			}
			log.Printf("Starting app server on: %v:%v", configApp.ServerAddress, configApp.ServerPort)
			err = http.ListenAndServe(configApp.ServerAddress+":"+configApp.ServerPort, http.FileServer(http.Dir("./wasm")))
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
