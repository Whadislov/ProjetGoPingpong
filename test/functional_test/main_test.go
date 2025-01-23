package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/Whadislov/ProjetGoPingPong/api"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// TestMainFunction tests the main function of the application.
func TestMainFunction(t *testing.T) {
	// Load the configuration file
	config := api.Config{
		ServerAddress: "localhost",
		ServerPort:    "8000",
	}

	// Start the API server in a separate goroutine
	go func() {
		api.RunApi(&config)
	}()

	// Allow some time for the server to start
	time.Sleep(2 * time.Second)

	// Test the /players endpoint
	resp, err := http.Get("http://" + config.ServerAddress + ":" + config.ServerPort + "/players")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Simulate the UI interaction
	go mu.Display("local")

	// Additional tests for other endpoints and UI interactions can be added here
}
