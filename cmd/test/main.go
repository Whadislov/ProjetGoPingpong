package main

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {

	golangDB, err := mf.LoadDB()
	if err != nil {
		panic(err)
	}
	mu.Display(golangDB)

}
