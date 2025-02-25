package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// OptionAuthPage sets up the option page at the start of the app, in which the user can change the theme color, the language
func OptionAuthPage(w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("options"), 32)

	themeButton := widget.NewButton(T("change_theme_color"), func() {
		returnToAuthentificationPageButton := widget.NewButton(T("return_to_authentification_page"), func() {
			w.SetContent(AuthentificationPage(w, a))
		})

		if darkTheme.IsActivated {
			a.Settings().SetTheme(&lightTheme)
			lightTheme.IsActivated = true
			darkTheme.IsActivated = false
			w.SetContent(container.NewVBox(OptionAuthPage(w, a), returnToAuthentificationPageButton))
		} else {
			a.Settings().SetTheme(&darkTheme)
			lightTheme.IsActivated = false
			darkTheme.IsActivated = true
			w.SetContent(container.NewVBox(OptionAuthPage(w, a), returnToAuthentificationPageButton))
		}
	})

	changeLanguageButton := widget.NewButton(T("change_language"), func() { w.SetContent(ChangeLanguageAuthPage(w, a)) })

	optionPage := container.NewVBox(
		pageTitle,
		themeButton,
		changeLanguageButton,
	)

	return optionPage
}

func ChangeLanguageAuthPage(w fyne.Window, a fyne.App) *fyne.Container {
	returnToAuthentificationPageButton := widget.NewButton(T("return_to_authentification_page"), func() {
		w.SetContent(AuthentificationPage(w, a))
	})

	languageSelector := widget.NewSelect([]string{"Deutsch", "English", "Français"}, func(selected string) {
		switch selected {
		case "English":
			setLanguage("en")
			currentSelectedLanguage = "English"
			// Refresh
			returnToAuthentificationPageButton = widget.NewButton(T("return_to_authentification_page"), func() {
				w.SetContent(AuthentificationPage(w, a))
			})
			w.SetContent(container.NewVBox(OptionAuthPage(w, a), returnToAuthentificationPageButton))
		case "Français":
			setLanguage("fr")
			currentSelectedLanguage = "Français"
			// Refresh
			returnToAuthentificationPageButton = widget.NewButton(T("return_to_authentification_page"), func() {
				w.SetContent(AuthentificationPage(w, a))
			})
			w.SetContent(container.NewVBox(OptionAuthPage(w, a), returnToAuthentificationPageButton))
		case "Deutsch":
			setLanguage("de")
			currentSelectedLanguage = "Deutsch"
			// Refresh
			returnToAuthentificationPageButton = widget.NewButton(T("return_to_authentification_page"), func() {
				w.SetContent(AuthentificationPage(w, a))
			})
			w.SetContent(container.NewVBox(OptionAuthPage(w, a), returnToAuthentificationPageButton))
		}
	})
	languageSelector.PlaceHolder = currentSelectedLanguage
	languageSelector.Alignment = fyne.TextAlignCenter
	//languageSelector.SetSelected(currentSelectedLanguage)

	changeLanguagePage := container.NewVBox(
		OptionAuthPage(w, a),
		languageSelector,
		returnToAuthentificationPageButton,
	)

	return changeLanguagePage

}
