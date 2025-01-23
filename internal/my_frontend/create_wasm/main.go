package main

import (
	mu "github.com/Whadislov/ProjetGoPingPong/internal/my_ui"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Used to create the web assembly folder, using fyne serve --sourceDir internal/my_frontend
func main() {
	mu.Display("browser")
}
