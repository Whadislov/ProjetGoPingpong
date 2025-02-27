package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"github.com/Whadislov/TTCompanion/api"
	mdb "github.com/Whadislov/TTCompanion/internal/my_db"
	mu "github.com/Whadislov/TTCompanion/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

//go:embed translation/*
var translations embed.FS

func loadConfig(filename string) (string, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	type Config struct {
		ServerAddress string `json:"server_address"`
		ServerPort    string `json:"server_port"`
	}

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return "", "", err
	}

	return config.ServerAddress, config.ServerPort, nil
}

func waitForAPI(url string, retries int, delay time.Duration) {
	for i := 0; i < retries; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Println("API is ready!")
			return
		}
		log.Println("Waiting for API to be ready...")
		time.Sleep(delay)
	}
	log.Fatal("API did not start in time")
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

		api.SetJWTSecretKey(os.Getenv("JWT_SECRET_KEY"))
		mdb.SetPsqlInfo(os.Getenv("WEB_DB_LINK"))
		mdb.SetDBName(os.Getenv("DB_NAME"))

		var wg sync.WaitGroup

		// API
		wg.Add(1)
		go func() {
			defer wg.Done()
			api.RunApi()
		}()

		// Verify that the API is ready
		apiURL := "http://localhost:8001/api/healthz"
		waitForAPI(apiURL, 10, 500*time.Millisecond)

		// App
		wg.Add(1)
		go func() {
			defer wg.Done()
			serverAddress, serverPort, err := loadConfig("config_app.json")
			if err != nil {
				log.Fatalf("Cannot read config file: %v", err)
			}
			log.Printf("Starting app server on %v:%v", serverAddress, serverPort)

			err = http.ListenAndServe(serverAddress+":"+serverPort, http.FileServer(http.Dir("./wasm")))
			if err != nil {
				log.Fatalf("App server error: %v", err)
			}
			/*
				fs := http.FileServer(http.Dir("./wasm")) // Mix
				http.Handle("/", fs)
				err = http.ListenAndServe(config.ServerAddress+":"+config.ServerPort, nil)
				if err != nil {
					log.Fatalf("App server error: %v", err)
				}
			*/
		}()

		wg.Wait()

	}

	// Start app locally
	if appStartOption == "local" {
		// Load translations
		mu.InitTranslations(translations)

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
