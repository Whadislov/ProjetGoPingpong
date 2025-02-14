package api

import (
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	ServerPort    string `json:"server_port"`
}

var jwtSecret []byte

func SetJWTSecretKey(jwtSecretString string) {
	jwtSecret = []byte(os.Getenv(jwtSecretString))
}
