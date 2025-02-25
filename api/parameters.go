package api

import (
	"os"
)

var jwtSecret []byte

func SetJWTSecretKey(jwtSecretString string) {
	jwtSecret = []byte(os.Getenv(jwtSecretString))
}
