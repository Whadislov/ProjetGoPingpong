package myapp

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// AddPage sets up the main page for adding players to teams and vice versa.
func AddPage(db *mt.Database, w fyne.Window, a fyne.App) {

	ReturnToFonctionalityPageButton := widget.NewButton("Return to the functionalities", func() {
		fonctionalityPage := FunctionalityPage(db, w, a)
		w.SetContent(fonctionalityPage)
	})

	addTtoPButton := widget.NewButton("Add team(s) to a player", func() {
		w.SetContent(
			currentSelectionPageTtoP(
				SelectionPageTtoP(db, w, a),
				nil,
				db, w, a,
			),
		)
	})

	addPtoTButton := widget.NewButton("Add player(s) to a team", func() {
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPagePtoT(db, w, a),
				nil,
				db, w, a,
			),
		)

	})

	addCtoPButton := widget.NewButton("Add club(s) to a player", func() {
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPageCtoP(db, w, a),
				nil,
				db, w, a,
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
