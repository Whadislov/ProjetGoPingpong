package api

import (
	"log"
	"net/http"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func RunApi(db *mt.Database) {
	http.HandleFunc("/players", GetPlayers(db))
	http.HandleFunc("/teams", GetTeams(db))
	http.HandleFunc("/clubs", GetClubs(db))
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
