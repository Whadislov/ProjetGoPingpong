package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mdb "github.com/Whadislov/TTCompanion/internal/my_db"
	"github.com/joho/godotenv"
)

func RunApi() {

	// Set env variables
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Cannot load variables from .env")
	}

	SetJWTSecretKey(os.Getenv("JWT_SECRET_KEY"))
	mdb.SetPsqlInfo(os.Getenv("WEB_DB_LINK"))
	mdb.SetDBName(os.Getenv("DB_NAME"))

	http.Handle("/api/healthz", CorsMiddleware(http.HandlerFunc(IsApiReady)))
	http.Handle("/api/load-database", CorsMiddleware(authMiddleware(http.HandlerFunc(loadUserDatabaseHandler))))
	http.Handle("/api/save-database", CorsMiddleware(authMiddleware(http.HandlerFunc(saveDatabaseHandler))))
	http.Handle("/api/login", CorsMiddleware(http.HandlerFunc(LoginHandler)))
	http.Handle("/api/signup", CorsMiddleware(http.HandlerFunc(SignUpHandler)))

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to TT Companion's API")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to TT Companion")
	})
}
