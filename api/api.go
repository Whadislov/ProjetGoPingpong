package api

import (
	"fmt"
	"log"
	"net/http"
)

func RunApi(config *Config) {
	http.Handle("/api", CorsMiddleware(http.HandlerFunc(IsApiReady)))
	http.Handle("/api/load-database", CorsMiddleware(authMiddleware(http.HandlerFunc(loadUserDatabaseHandler))))
	http.Handle("/api/save-database", CorsMiddleware(authMiddleware(http.HandlerFunc(saveDatabaseHandler))))
	http.Handle("/api/login", CorsMiddleware(http.HandlerFunc(LoginHandler)))
	http.Handle("/api/signup", CorsMiddleware(http.HandlerFunc(SignUpHandler)))

	address := fmt.Sprintf("%s:%s", config.ServerAddress, config.ServerPort)
	log.Printf("API started on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
