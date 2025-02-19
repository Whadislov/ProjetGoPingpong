package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// Tracks if the database has changed
var HasChanged bool

// MainPage creates the main page
func MainPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	// Initialize
	var mainPage *fyne.Container

	// Database page
	databasePage := DatabasePage(db, w, a)

	// Functionality page
	functionalityPage := FunctionalityPage(db, w, a)

	mainText := setTitle("TT Companion", 32)

	showDBButton := widget.NewButton("Your database", func() {
		w.SetContent(databasePage)
	})

	showFuncButton := widget.NewButton("Functionalities", func() {
		w.SetContent(functionalityPage)
	})

	// Options button
	themeButton := widget.NewButton("Options", func() {

		if darkTheme.IsActivated {
			a.Settings().SetTheme(&lightTheme)
			lightTheme.IsActivated = true
			darkTheme.IsActivated = false
			w.SetContent(MainPage(db, w, a))
		} else {
			a.Settings().SetTheme(&darkTheme)
			lightTheme.IsActivated = false
			darkTheme.IsActivated = true
			w.SetContent(MainPage(db, w, a))
		}
	})

	quitButton := widget.NewButton("Quit", func() {
		Quit(db, w, a, HasChanged)
	})

	if appStartOption == "local" {
		mainPage = container.NewVBox(
			mainText,
			showDBButton,
			showFuncButton,
			themeButton,
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
