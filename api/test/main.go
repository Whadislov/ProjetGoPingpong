package main

import (
	"github.com/Whadislov/ProjetGoPingPong/api"
)

func main() {
	config := &api.Config{
		ServerAddress: "localhost",
		ServerPort:    "7000",
	}
	api.RunApi(config)

}
