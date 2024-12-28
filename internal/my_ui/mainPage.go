package myapp

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Tracks if the database has changed
var HasChanged bool

// MainPage creates the main page
func MainPage(sqlDB *msql.Database, golangDB *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	// Database page
	databasePage := DatabasePage(sqlDB, golangDB, w, a)

	// Functionality page
	functionalityPage := FunctionalityPage(sqlDB, golangDB, w, a)

	// Main page design
	mainText := canvas.NewText("TTapp 🏓", color.RGBA{R: 0, G: 0, B: 0, A: 255})
	mainText.Alignment = fyne.TextAlignCenter
	mainText.TextSize = 32

	mainPage := container.NewVBox(
		mainText,
		widget.NewButton("Show database", func() {
			w.SetContent(databasePage)
		}),
		widget.NewButton("Show functionalities", func() {
			w.SetContent(functionalityPage)
		}),
		widget.NewButton("Quit", func() {
			Quit(sqlDB, golangDB, w, a, HasChanged)
		}),
	)

	// Check for unsaved changes before quitting
	w.SetCloseIntercept(func() {
		Quit(sqlDB, golangDB, w, a, HasChanged)
	})

	return mainPage

}
