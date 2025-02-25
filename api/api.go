package api

import (
	"net/http"
)

func RunApi() {
	http.Handle("/api", CorsMiddleware(http.HandlerFunc(IsApiReady)))
	http.Handle("/api/load-database", CorsMiddleware(authMiddleware(http.HandlerFunc(loadUserDatabaseHandler))))
	http.Handle("/api/save-database", CorsMiddleware(authMiddleware(http.HandlerFunc(saveDatabaseHandler))))
	http.Handle("/api/login", CorsMiddleware(http.HandlerFunc(LoginHandler)))
	http.Handle("/api/signup", CorsMiddleware(http.HandlerFunc(SignUpHandler)))
}
