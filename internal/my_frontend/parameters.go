package myfrontend

import ()

type Config struct {
	ServerAddress string `json:"server_address"`
	ServerPort    string `json:"server_port"`
}

var apiConfig string = "http://localhost:8001/"
