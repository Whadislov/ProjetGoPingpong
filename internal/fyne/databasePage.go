package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"image/color"
)

func DatabasePage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	backToMainMenuButton := widget.NewButton("Back to main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
		w.Show()
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
		backToMainMenuButton,
	)

	return databasePage
}
