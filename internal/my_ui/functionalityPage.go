package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
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
	createDeleteMenuButton := widget.NewButton("Delete", func() { DeletePage(db, w, a) })
	createAddInfoToPlayerButton := widget.NewButton("Modify player information", func() { AddInfoToPlayerPage(db, w, a) })

	themeColor := a.Settings().Theme().Color("foreground", a.Settings().ThemeVariant())
	functionalityText := canvas.NewText("Functionalities", themeColor)
	functionalityText.Alignment = fyne.TextAlignCenter
	functionalityText.TextSize = 32

	functionalityPage := container.NewVBox(
		functionalityText,
		createMenuButton,
		createAddMenuButton,
		createRemoveMenuButton,
		createDeleteMenuButton,
		createAddInfoToPlayerButton,
		returnToMainMenuButton,
	)

	return functionalityPage
}
