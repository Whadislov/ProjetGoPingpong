package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func MainMenu(w fyne.Window) {
	// Split window

	label1 := canvas.NewText("Team 1", color.Black)
	label2 := canvas.NewText("Players", color.Black)
	w1 := widget.NewIcon(theme.CancelIcon())
	btn1 := widget.NewButton("Add Player ", func() {
	})

	// Accordion

	item1 := widget.NewAccordionItem("A",
		widget.NewLabel("A for Apple"))
	item2 := widget.NewAccordionItem("B",
		widget.NewLabel("B for Ball"))
	item3 := widget.NewAccordionItem("C",
		widget.NewLabel("C for Cat"))
	ac := widget.NewAccordion(item1, item2, item3)

	itemt1 := widget.NewAccordionItem("Team 1",
		widget.NewLabel("A for Apple"))
	itemt2 := widget.NewAccordionItem("Team 2",
		widget.NewLabel("B for Ball"))
	itemt3 := widget.NewAccordionItem("Team 3",
		widget.NewLabel("C for Cat"))
	act := widget.NewAccordion(itemt1, itemt2, itemt3)

	// Setup Content
	w.SetContent(
		container.NewHSplit(
			container.NewVBox(
				label1,
				w1,
				act,
			),
			// 2nd Section
			container.NewVBox(
				label2,
				btn1,
				ac,
			),
		),
	)

}
