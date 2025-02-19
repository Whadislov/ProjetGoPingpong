package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// DatabasePage sets up the page for showing players, teams, and clubs.
func OptionPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Options", 32)

	returnToMainMenuButton := widget.NewButton("Return to main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
	})

	themeButton := widget.NewButton("Change theme color", func() {

		if darkTheme.IsActivated {
			a.Settings().SetTheme(&lightTheme)
			lightTheme.IsActivated = true
			darkTheme.IsActivated = false
			w.SetContent(OptionPage(db, w, a))
		} else {
			a.Settings().SetTheme(&darkTheme)
			lightTheme.IsActivated = false
			darkTheme.IsActivated = true
			w.SetContent(OptionPage(db, w, a))
		}
	})

	changeLanguageButton := widget.NewButton("Change language", func() {
	})

	optionPage := container.NewVBox(
		pageTitle,
		themeButton,
		changeLanguageButton,
		returnToMainMenuButton,
	)

	return optionPage
}
