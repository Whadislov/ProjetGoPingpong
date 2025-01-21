package main

import (
	"github.com/Whadislov/ProjetGoPingPong/api"
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
)

func main() {
	config := &api.Config{
		ServerAddress: "localhost",
		ServerPort:    "7000",
	}
	appStartOption := "browser"
	mdb.AppStartOption(appStartOption)
	api.RunApi(config)

}
