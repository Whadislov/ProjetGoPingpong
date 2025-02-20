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

	themeButton := widget.NewButton("Change theme color", func() {
		returnToMainMenuButton := widget.NewButton(T("return_to_main_page"), func() {
			w.SetContent(MainPage(db, w, a))
		})

		if darkTheme.IsActivated {
			a.Settings().SetTheme(&lightTheme)
			lightTheme.IsActivated = true
			darkTheme.IsActivated = false
			w.SetContent(container.NewVBox(OptionPage(db, w, a), returnToMainMenuButton))
		} else {
			a.Settings().SetTheme(&darkTheme)
			lightTheme.IsActivated = false
			darkTheme.IsActivated = true
			w.SetContent(container.NewVBox(OptionPage(db, w, a), returnToMainMenuButton))
		}
	})

	changeLanguageButton := widget.NewButton("Change language", func() { w.SetContent(ChangeLanguagePage(db, w, a)) })

	optionPage := container.NewVBox(
		pageTitle,
		themeButton,
		changeLanguageButton,
	)

	return optionPage
}

func ChangeLanguagePage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToMainMenuButton := widget.NewButton(T("return_to_main_page"), func() {
		w.SetContent(MainPage(db, w, a))
	})

	languageSelector := widget.NewSelect([]string{"English", "Français"}, func(selected string) {
		switch selected {
		case "English":
			loadLanguage("en")
			currentSelectedLanguage = "English"
			w.SetContent(container.NewVBox(OptionPage(db, w, a), returnToMainMenuButton))
		case "Français":
			loadLanguage("fr")
			currentSelectedLanguage = "Français"
			w.SetContent(container.NewVBox(OptionPage(db, w, a), returnToMainMenuButton))
		}
	})
	languageSelector.PlaceHolder = currentSelectedLanguage
	languageSelector.Alignment = fyne.TextAlignCenter
	//languageSelector.SetSelected(currentSelectedLanguage)

	changeLanguagePage := container.NewVBox(
		OptionPage(db, w, a),
		languageSelector,
		returnToMainMenuButton,
	)

	return changeLanguagePage

}
