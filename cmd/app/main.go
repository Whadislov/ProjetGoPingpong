package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Whadislov/ProjetGoPingPong/api"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {

	// Start app locally or on a browser ?
	appStartOption := "local"

	// Start app on a browser
	if appStartOption == "browser" {
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
		mu.Display(appStartOption)
	}
}
