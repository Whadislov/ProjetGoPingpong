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

	var mainPage *fyne.Container
	pageTitle := setTitle("TT Companion", 32)

	showDBButton := widget.NewButton(T("your_database"), func() {
		w.SetContent(DatabasePage(db, w, a))
	})

	showFuncButton := widget.NewButton(T("functionalities"), func() {
		w.SetContent(FunctionalityPage(db, w, a))
	})

	// Options button
	OptionButton := widget.NewButton(T("options"), func() {
		returnToMainMenuButton := widget.NewButton(T("return_to_main_page"), func() {
			w.SetContent(MainPage(db, w, a))
		})
		w.SetContent(container.NewVBox(OptionPage(db, w, a), returnToMainMenuButton))
	})

	quitButton := widget.NewButton(T("quit"), func() {
		Quit(db, w, a, HasChanged)
	})

	if appStartOption == "local" {
		mainPage = container.NewVBox(
			pageTitle,
			showDBButton,
			showFuncButton,
			OptionButton,
			quitButton,
		)

	} else if appStartOption == "browser" {
		// Remove the quit button
		mainPage = container.NewVBox(
			pageTitle,
			showDBButton,
			showFuncButton,
			OptionButton,
		)
	}

	// Check for unsaved changes before quitting
	w.SetCloseIntercept(func() {
		Quit(db, w, a, HasChanged)
	})

	return mainPage

}
