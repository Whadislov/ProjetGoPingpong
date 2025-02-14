package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// DatabasePage sets up the page for showing players, teams, and clubs.
func DatabasePage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Your database", 32)

	returnToMainMenuButton := widget.NewButton("Return to main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
	})

	playerButton := widget.NewButton("Your players", func() { PlayerPage(db, w, a) })
	teamButton := widget.NewButton("Your teams", func() { TeamPage(db, w, a) })
	clubButton := widget.NewButton("Your clubs", func() { ClubPage(db, w, a) })

	databasePage := container.NewVBox(
		pageTitle,
		playerButton,
		teamButton,
		clubButton,
		returnToMainMenuButton,
	)

	return databasePage
}
