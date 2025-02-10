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
			0: {ID: 0, Firstname: "Julien", Lastname: "M", Age: 27, Ranking: 1632, Material: []string{"Forehand", "Backhand", "Blade"}},
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
		ServerPort:    "7000",
	}
}

func TestRunApi(t *testing.T) {
	config := setupMockConfig()

	go func() {
		api.RunApi(config)
	}()

	// Let the server have time to start
	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://" + config.ServerAddress + ":" + config.ServerPort + "/players")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var players map[int]*mt.Player
	err = json.NewDecoder(resp.Body).Decode(&players)
	if err != nil {
		t.Fatalf("Failed to decode players body: %v", err)
	}

	if len(players) != 1 || players[0].Firstname != "Julien" {
		t.Fatalf("Unexpected players data: %+v", players)
	}

	resp, err = http.Get("http://" + config.ServerAddress + ":" + config.ServerPort + "/teams")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	var teams map[int]*mt.Team
	err = json.NewDecoder(resp.Body).Decode(&teams)
	if err != nil {
		t.Fatalf("Failed to decode teams body: %v", err)
	}

	if len(teams) != 1 || teams[0].Name != "Mannschaft 2" {
		t.Fatalf("Unexpected teams data: %+v", players)
	}

	resp, err = http.Get("http://" + config.ServerAddress + ":" + config.ServerPort + "/clubs")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	var clubs map[int]*mt.Club
	err = json.NewDecoder(resp.Body).Decode(&clubs)
	if err != nil {
		t.Fatalf("Failed to decode clubs body: %v", err)
	}

	if len(clubs) != 1 || clubs[0].Name != "TSG Heilbronn" {
		t.Fatalf("Unexpected clubs data: %+v", clubs)
	}
}
