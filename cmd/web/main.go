package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Whadislov/TTCompanion/api"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func loadConfig(filename string) (string, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	type Config struct {
		ServerPrefix  string `json:"server_prefix"`
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
	for range retries {
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

	var wg sync.WaitGroup

	// API
	wg.Add(1)
	go func() {
		defer wg.Done()
		api.RunApi()
	}()

	// Verify that the API is ready
	apiAddress, apiPort, err := loadConfig("config_api.json")
	if err != nil {
		log.Fatalf("Cannot read config file: %v", err)
	}
	apiURL := "http://" + apiAddress + ":" + apiPort
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
	}()

	wg.Wait()

}
