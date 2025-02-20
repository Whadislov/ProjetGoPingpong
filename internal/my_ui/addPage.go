package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// AddPage sets up the main page for adding players to teams and vice versa.
func AddPage(db *mt.Database, w fyne.Window, a fyne.App) {
	pageTitle := setTitle(T("add"), 32)

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
		pageTitle,
		addTtoPButton,
		addPtoTButton,
		addCtoPButton,
		ReturnToFonctionalityPageButton,
	)

	w.SetContent(addPage)
}
