package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Whadislov/TTCompanion/api"
	mdb "github.com/Whadislov/TTCompanion/internal/my_db"
	mfr "github.com/Whadislov/TTCompanion/internal/my_frontend"
	mu "github.com/Whadislov/TTCompanion/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

//go:embed translation/*
var translations embed.FS

func loadConfig(filename string) (*mfr.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &mfr.Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func initWASM() {
	// Load translations
	mu.InitTranslations(translations)
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

		// API
		api.RunApi()
		// App
		config, err := loadConfig("config_app.json")
		if err != nil {
			log.Fatalf("Cannot read config file: %v", err)
		}
		log.Printf("Starting app server on: %v:%v", config.ServerAddress, config.ServerPort)
		fs := http.FileServer(http.Dir("./wasm"))
		http.Handle("/", fs)
		err = http.ListenAndServe(config.ServerAddress+":"+config.ServerPort, nil)
		if err != nil {
			log.Fatalf("App server error: %v", err)
		}

	}

	// Start app locally
	if appStartOption == "local" {
		initWASM()
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
