package myapp

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// DatabasePage sets up the page for showing players, teams, and clubs.
func DatabasePage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToMainMenuButton := widget.NewButton("Return to main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
	})

	playerButton := widget.NewButton("Show players", func() { PlayerPage(db, w, a) })
	teamButton := widget.NewButton("Show teams", func() { TeamPage(db, w, a) })
	clubButton := widget.NewButton("Show clubs", func() { ClubPage(db, w, a) })

	databaseText := canvas.NewText("Database", color.RGBA{R: 0, G: 0, B: 0, A: 255})
	databaseText.Alignment = fyne.TextAlignCenter
	databaseText.TextSize = 16

	databasePage := container.NewVBox(
		databaseText,
		playerButton,
		teamButton,
		clubButton,
		returnToMainMenuButton,
	)

	return databasePage
}
