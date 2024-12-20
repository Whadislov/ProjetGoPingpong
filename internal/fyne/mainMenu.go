package myapp

import (
	"fyne.io/fyne/v2"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func MainMenu(db *mt.Database, w fyne.Window, a fyne.App) *fyne.MainMenu {

	mainPage := MainPage(db, w, a)

	menu1Item1 := fyne.NewMenuItem("Show Main page", func() {
		w.SetContent(mainPage)
	})
	newMenu1 := fyne.NewMenu("Main menu", menu1Item1)

	menu2Item1 := fyne.NewMenuItem("Players", func() { PlayerPage(db, w, a) })
	menu2Item2 := fyne.NewMenuItem("Teams", func() { TeamPage(db, w, a) })
	menu2Item3 := fyne.NewMenuItem("Clubs", func() { ClubPage(db, w, a) })
	newMenu2 := fyne.NewMenu("Database", menu2Item1, menu2Item2, menu2Item3)

	menu3Item1 := fyne.NewMenuItem("Create ", func() { CreatePage(db, w, a) })
	menu3Item2 := fyne.NewMenuItem("Add ... to ...", func() { AddPage(db, w, a) })
	menu3Item3 := fyne.NewMenuItem("Remove ... from ...", func() {})
	menu3Item4 := fyne.NewMenuItem("Delete", func() { CreateDeletePage(db, w, a) })
	newMenu3 := fyne.NewMenu("Functions", menu3Item1, menu3Item2, menu3Item3, menu3Item4)

	menu := fyne.NewMainMenu(newMenu1, newMenu2, newMenu3)

	return menu

}
