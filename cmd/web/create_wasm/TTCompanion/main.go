package main

import (
	"embed"
	mu "github.com/Whadislov/TTCompanion/internal/my_frontend/my_ui"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

//go:embed translation/*
var translations embed.FS

// Used to create the web assembly folder, using fyne serve --sourceDir internal/my_frontend
func main() {
	// Load translations
	mu.InitTranslations(translations)
	mu.Display("browser")
}
