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

	menu1Item1 := fyne.NewMenuItem("Statistics", func() {})
	newMenu1 := fyne.NewMenu("Main menu", menu1Item1)

	menu2Item1 := fyne.NewMenuItem("Players", func() { PlayerMenu(w, db.Players) })
	menu2Item2 := fyne.NewMenuItem("Teams", func() { TeamMenu(w, db.Teams) })
	menu2Item3 := fyne.NewMenuItem("Clubs", func() { ClubMenu(w, db.Clubs) })
	newMenu2 := fyne.NewMenu("Database", menu2Item1, menu2Item2, menu2Item3)

	menu3Item1 := fyne.NewMenuItem("Create ", func() { CreateMenu(w, db) })
	menu3Item2 := fyne.NewMenuItem("Add ... to ...", func() {})
	menu3Item3 := fyne.NewMenuItem("Remove ... from ...", func() {})
	menu3Item4 := fyne.NewMenuItem("Delete", func() { CreateDeleteMenu(w, db) })
	newMenu3 := fyne.NewMenu("Functions", menu3Item1, menu3Item2, menu3Item3, menu3Item4)

	menu := fyne.NewMainMenu(newMenu1, newMenu2, newMenu3)

	MainMenu(w)
	w.SetMainMenu(menu)
	w.ShowAndRun()
}
