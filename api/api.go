package api

import (
	"fmt"
	"log"
	"net/http"
)

func RunApi(config *Config) {
	http.Handle("/api", corsMiddleware(http.HandlerFunc(isApiReady)))
	http.Handle("/api/load-database", corsMiddleware(http.HandlerFunc(loadDatabaseHandler)))
	http.Handle("/api/save-database", corsMiddleware(http.HandlerFunc(saveDatabaseHandler)))

	address := fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort)
	log.Printf("Server started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
