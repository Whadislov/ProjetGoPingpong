package myapp

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// DatabasePage sets up the page for showing players, teams, and clubs.
func DatabasePage(sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToMainMenuButton := widget.NewButton("Return to main page", func() {
		mainPage := MainPage(sqlDB, db, w, a)
		w.SetContent(mainPage)
	})

	playerButton := widget.NewButton("Show players", func() { PlayerPage(sqlDB, db, w, a) })
	teamButton := widget.NewButton("Show teams", func() { TeamPage(sqlDB, db, w, a) })
	clubButton := widget.NewButton("Show clubs", func() { ClubPage(sqlDB, db, w, a) })

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
