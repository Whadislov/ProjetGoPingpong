package myapp

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// Tracks if the database has changed
var HasChanged bool

// MainPage creates the main page
func MainPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	// Database page
	databasePage := DatabasePage(db, w, a)

	// Functionality page
	functionalityPage := FunctionalityPage(db, w, a)

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
			Quit(db, w, a, HasChanged)
		}),
	)

	// Check for unsaved changes before quitting
	w.SetCloseIntercept(func() {
		Quit(db, w, a, HasChanged)
	})

	return mainPage

}
