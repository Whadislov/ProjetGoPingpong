package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// FunctionalityPage creates the functionality page
func FunctionalityPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("functionalities"), 32)

	returnToMainMenuButton := widget.NewButton("Return main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
	})

	createMenuButton := widget.NewButton(T("create"), func() { CreatePage(db, w, a) })
	createAddMenuButton := widget.NewButton(T("add"), func() { AddPage(db, w, a) })
	createRemoveMenuButton := widget.NewButton(T("remove"), func() { RemovePage(db, w, a) })
	createDeleteMenuButton := widget.NewButton(T("delete"), func() { DeletePage(db, w, a) })
	createAddInfoToPlayerButton := widget.NewButton(T("edit_player_information"), func() { AddInfoToPlayerPage(db, w, a) })

	functionalityPage := container.NewVBox(
		pageTitle,
		createMenuButton,
		createAddMenuButton,
		createRemoveMenuButton,
		createDeleteMenuButton,
		createAddInfoToPlayerButton,
		returnToMainMenuButton,
	)

	return functionalityPage
}
