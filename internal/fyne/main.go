package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func Display(db *mt.Database) {
	// New app
	a := app.New()
	w := a.NewWindow("Main menu items")
	w.Resize(fyne.NewSize(400, 400)) // 400x400 window size

	// Menu
	menuItem1 := fyne.NewMenuItem("Players", func() { PlayerMenu(w, db.Players) })
	menuItem2 := fyne.NewMenuItem("Teams", func() { TeamMenu(w, db.Teams) })
	menuItem3 := fyne.NewMenuItem("Clubs", func() { ClubMenu(w, db.Clubs) })
	MainMenu(w)

	newMenu1 := fyne.NewMenu("Show database", menuItem1, menuItem2, menuItem3)
	newMenu2 := fyne.NewMenu("Functions")

	menu := fyne.NewMainMenu(newMenu1, newMenu2)
	w.SetMainMenu(menu)
	w.ShowAndRun()
}
