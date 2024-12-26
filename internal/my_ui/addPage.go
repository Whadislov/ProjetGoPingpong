package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// AddPage sets up the main page for adding players to teams and vice versa.
func AddPage(sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) {

	ReturnToFonctionalityPageButton := widget.NewButton("Return to the functionalities", func() {
		fonctionalityPage := FunctionalityPage(sqlDB, db, w, a)
		w.SetContent(fonctionalityPage)
	})

	addTtoPButton := widget.NewButton("Add team(s) to a player", func() {
		w.SetContent(
			currentSelectionPageTtoP(
				SelectionPageTtoP(sqlDB, db, w, a),
				nil,
				sqlDB, db, w, a,
			),
		)
	})

	addPtoTButton := widget.NewButton("Add player(s) to a team", func() {
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPagePtoT(sqlDB, db, w, a),
				nil,
				sqlDB, db, w, a,
			),
		)

	})

	addCtoPButton := widget.NewButton("Add club(s) to a player", func() {
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPageCtoP(sqlDB, db, w, a),
				nil,
				sqlDB, db, w, a,
			),
		)

	})

	addPage := container.NewVBox(
		addTtoPButton,
		addPtoTButton,
		addCtoPButton,
		ReturnToFonctionalityPageButton,
	)

	w.SetContent(addPage)
}
