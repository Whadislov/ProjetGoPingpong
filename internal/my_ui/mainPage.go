package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
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
	themeColor := a.Settings().Theme().Color("foreground", a.Settings().ThemeVariant())
	mainText := canvas.NewText("TT Companion", themeColor)
	mainText.Alignment = fyne.TextAlignCenter
	mainText.TextSize = 32

	showDBButton := widget.NewButton("Your database", func() {
		w.SetContent(databasePage)
	})

	showFuncButton := widget.NewButton("Functionalities", func() {
		w.SetContent(functionalityPage)
	})

	quitButton := widget.NewButton("Quit", func() {
		Quit(db, w, a, HasChanged)
	})

	// Initialize
	var mainPage *fyne.Container

	if appStartOption == "local" {
		mainPage = container.NewVBox(
			mainText,
			showDBButton,
			showFuncButton,
			quitButton,
		)

	} else if appStartOption == "browser" {
		// Remove the quit button
		mainPage = container.NewVBox(
			mainText,
			showDBButton,
			showFuncButton,
		)
	}

	// Check for unsaved changes before quitting
	w.SetCloseIntercept(func() {
		Quit(db, w, a, HasChanged)
	})

	return mainPage

}
