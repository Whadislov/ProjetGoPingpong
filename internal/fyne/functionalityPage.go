package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"image/color"
)

func FunctionalityPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	backToMainMenuButton := widget.NewButton("Back to main page", func() {
		mainPage := MainPage(db, w, a)
		w.SetContent(mainPage)
		w.Show()
	})

	createMenuButton := widget.NewButton("Create", func() { CreatePage(db, w, a) })
	createAddMenuButton := widget.NewButton("Add ... to ...", func() {})
	createRemoveMenuButton := widget.NewButton("Remove ... from ...", func() {})
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
		backToMainMenuButton,
	)

	return functionalityPage
}
