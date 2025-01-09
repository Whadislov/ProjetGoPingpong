package integrationtests

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/Whadislov/ProjetGoPingPong/api"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Mock database
func setupMockDB() *mt.Database {
	return &mt.Database{
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
}

// Mock config
func setupMockConfig() *api.Config {
	return &api.Config{
		ServerAddress: "localhost",
		ServerPort:    "8000",
	}
}

func TestRunApi(t *testing.T) {
	db := setupMockDB()
	config := setupMockConfig()

	// Start the API server in a separate goroutine
	go func() {
		api.RunApi(db, config)
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

	// Decode the players body
	var players map[int]*mt.Player
	err = json.NewDecoder(resp.Body).Decode(&players)
	if err != nil {
		t.Fatalf("Failed to decode players body: %v", err)
	}

	// Check the players data
	if len(players) != 1 || players[0].Name != "Julien" {
		t.Fatalf("Unexpected players data: %+v", players)
	}

	// Test the /teams endpoint
	resp, err = http.Get("http://" + config.ServerAddress + ":" + config.ServerPort + "/teams")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	// Decode the teams body
	var teams map[int]*mt.Team
	err = json.NewDecoder(resp.Body).Decode(&teams)
	if err != nil {
		t.Fatalf("Failed to decode teams body: %v", err)
	}

	// Check the teams data
	if len(teams) != 1 || teams[0].Name != "Mannschaft 2" {
		t.Fatalf("Unexpected teams data: %+v", players)
	}

	// Test the /clubs endpoint
	resp, err = http.Get("http://" + config.ServerAddress + ":" + config.ServerPort + "/clubs")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	// Decode the clubs body
	var clubs map[int]*mt.Club
	err = json.NewDecoder(resp.Body).Decode(&clubs)
	if err != nil {
		t.Fatalf("Failed to decode clubs body: %v", err)
	}

	// Check the clubs data
	if len(clubs) != 1 || clubs[0].Name != "TSG Heilbronn" {
		t.Fatalf("Unexpected clubs data: %+v", clubs)
	}
}
