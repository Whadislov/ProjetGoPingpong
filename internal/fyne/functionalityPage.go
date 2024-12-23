package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"image/color"
)

// FunctionalityPage creates the functionality page
func FunctionalityPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToMainMenuButton := widget.NewButton("Return main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
	})

	createMenuButton := widget.NewButton("Create", func() { CreatePage(db, w, a) })
	createAddMenuButton := widget.NewButton("Add ... to ...", func() { AddPage(db, w, a) })
	createRemoveMenuButton := widget.NewButton("Remove ... from ...", func() { RemovePage(db, w, a) })
	createDeleteMenuButton := widget.NewButton("Delete", func() { CreateDeletePage(db, w, a) })

	functionalityText := canvas.NewText("Functionalities", color.RGBA{R: 0, G: 0, B: 0, A: 255})
	functionalityText.Alignment = fyne.TextAlignCenter
	functionalityText.TextSize = 16

	functionalityPage := container.NewVBox(
		functionalityText,
		createMenuButton,
		createAddMenuButton,
		createRemoveMenuButton,
		createDeleteMenuButton,
		returnToMainMenuButton,
	)

	return functionalityPage
}
