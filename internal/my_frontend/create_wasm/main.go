package main

import (
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Used to create the web assembly folder, using fyne serve --sourceDir internal/my_frontend
func main() {

	appStartOption := "browser"
	mdb.AppStartOption(appStartOption)
	mu.AppStartOption(appStartOption)

	golangDB, err := mf.LoadDB()
	if err != nil {
		panic(err)
	}
	mu.Display(golangDB)

}
