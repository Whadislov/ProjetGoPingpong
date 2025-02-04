package api

import (
	"fmt"
	"log"
	"net/http"
)

func RunApi(config *Config) {
	http.Handle("/api", corsMiddleware(http.HandlerFunc(isApiReady)))
	http.Handle("/api/load-database", corsMiddleware(authMiddleware(http.HandlerFunc(loadUserDatabaseHandler))))
	http.Handle("/api/save-database", corsMiddleware(http.HandlerFunc(saveDatabaseHandler)))
	http.Handle("/api/login", corsMiddleware(http.HandlerFunc(loginHandler)))
	http.Handle("/api/signup", corsMiddleware(http.HandlerFunc(signUpHandler)))

	address := fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort)
	log.Printf("API started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
