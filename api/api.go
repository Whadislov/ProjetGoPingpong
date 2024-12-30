package api

import (
	"fmt"
	"log"
	"net/http"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func RunApi(db *mt.Database, config *Config) {
	http.HandleFunc("/players", GetPlayers(db))
	http.HandleFunc("/teams", GetTeams(db))
	http.HandleFunc("/clubs", GetClubs(db))

	address := fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort)
	log.Printf("Server started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
