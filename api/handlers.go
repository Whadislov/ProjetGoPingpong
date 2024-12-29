package api

import (
	"encoding/json"
	"net/http"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// GetPlayers handles GET requests to retrieve all players
func GetPlayers(db *mt.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(db.Players)
	}
}

// GetTeams handles GET requests to retrieve all teams
func GetTeams(db *mt.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(db.Teams)
	}
}

// GetClubs handles GET requests to retrieve all clubs
func GetClubs(db *mt.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(db.Clubs)
	}
}
