package myfrontend

import (
	"net/http"
)

// IsApiReady posts a message if the API is started
func IsApiReady() bool {
	resp, err := http.Get(config)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return true
}
