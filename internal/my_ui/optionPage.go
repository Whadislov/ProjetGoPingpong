package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// OptionPage sets up the page in which the user can change the theme color, the language
func OptionPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("options"), 32)

	themeButton := widget.NewButton(T("change_theme_color"), func() {
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

	changeLanguageButton := widget.NewButton(T("change_language"), func() { w.SetContent(ChangeLanguagePage(db, w, a)) })

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
