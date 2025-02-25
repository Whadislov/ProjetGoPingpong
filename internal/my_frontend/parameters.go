package myfrontend

import ()

type Config struct {
	ServerAddress string `json:"server_address"`
	ServerPort    string `json:"server_port"`
}

var config string

func SetConfig(c *Config) {
	config = c.ServerAddress + ":" + c.ServerPort + "/"
}
