package api

import (
	"fmt"
	"log"
	"net/http"
)

func RunApi(config *Config) {
	http.Handle("/api", corsMiddleware(http.HandlerFunc(isApiReady)))
	http.Handle("/api/load-database", corsMiddleware(http.HandlerFunc(loadDatabaseHandler)))
	http.Handle("/api/load-users", corsMiddleware(http.HandlerFunc(loadUsersHandler)))
	http.Handle("/api/save-database", corsMiddleware(http.HandlerFunc(saveDatabaseHandler)))
	http.Handle("/api/authenticate-user", corsMiddleware(http.HandlerFunc(authenticateUserHandler)))

	address := fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort)
	log.Printf("API started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
