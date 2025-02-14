package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// FunctionalityPage creates the functionality page
func FunctionalityPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Functionalities", 32)

	returnToMainMenuButton := widget.NewButton("Return main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
	})

	createMenuButton := widget.NewButton("Create", func() { CreatePage(db, w, a) })
	createAddMenuButton := widget.NewButton("Add", func() { AddPage(db, w, a) })
	createRemoveMenuButton := widget.NewButton("Remove", func() { RemovePage(db, w, a) })
	createDeleteMenuButton := widget.NewButton("Delete", func() { DeletePage(db, w, a) })
	createAddInfoToPlayerButton := widget.NewButton("Edit player information", func() { AddInfoToPlayerPage(db, w, a) })

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
