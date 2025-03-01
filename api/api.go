package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to TT Companion's API")
	})

	// Get server address and port
	config, err := loadConfig("config_api.json")
	if err != nil {
		log.Fatalf("Cannot read config file: %v", err)
	}

	log.Printf("API started on %s:%s", config.ServerAddress, config.ServerPort)

	errLS := http.ListenAndServe(config.ServerAddress+":"+config.ServerPort, nil)
	if err != nil {
		log.Fatalf("Api error: %v", errLS)
	}
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	c := &Config{}
	err = decoder.Decode(c)
	if err != nil {
		return nil, err
	}

	// Remove http:// if present
	c.ServerAddress = cleanAddress(c.ServerAddress)

	return c, nil
}

func cleanAddress(address string) string {
	if strings.HasPrefix(address, "http://") {
		return address[7:]
	}
	return address
}
