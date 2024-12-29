package myapp

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// RemovePage sets up the main page for removing players from teams and vice versa.
func RemovePage(db *mt.Database, w fyne.Window, a fyne.App) {

	ReturnToFonctionalityPageButton := widget.NewButton("Return to the functionalities", func() {
		fonctionalityPage := FunctionalityPage(db, w, a)
		w.SetContent(fonctionalityPage)
	})

	removeTfromPButton := widget.NewButton("Remove team(s) from a player", func() {
		w.SetContent(
			currentSelectionPageTfromP(
				SelectionPageTfromP(db, w, a),
				nil,
				db, w, a,
			),
		)
	})

	removePfromTButton := widget.NewButton("Remove player(s) from a team", func() {
		w.SetContent(
			currentSelectionPagePfromT(
				SelectionPagePfromT(db, w, a),
				nil,
				db, w, a,
			),
		)

	})

	removeCfromPButton := widget.NewButton("Remove club(s) from a player", func() {
		w.SetContent(
			currentSelectionPagePfromT(
				SelectionPageCfromP(db, w, a),
				nil,
				db, w, a,
			),
		)

	})

	removePage := container.NewVBox(
		removeTfromPButton,
		removePfromTButton,
		removeCfromPButton,
		ReturnToFonctionalityPageButton,
	)

	w.SetContent(removePage)
}
