package myapp

import (
	"image/color"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ClubInfos(club *mt.Club) *fyne.Container {
	wt := []fyne.CanvasObject{}
	wp := []fyne.CanvasObject{}
	for _, player := range club.PlayerIDs {
		wp = append(wp, widget.NewLabel(player))
	}
	itemp := widget.NewAccordionItem("Show players",
		container.NewVBox(wp...),
	)
	playerAc := widget.NewAccordion(itemp)
	for _, team := range club.TeamIDs {
		wt = append(wt, widget.NewLabel(team))
	}
	itemt := widget.NewAccordionItem("Show teams",
		container.NewVBox(wt...),
	)
	teamAc := widget.NewAccordion(itemt)

	item := container.NewVBox(
		teamAc,
		playerAc,
	)
	return item
}

func ClubMenu(w fyne.Window, clubs map[int]*mt.Club) {
	label := canvas.NewText("Clubs: ", color.Black)
	ac := widget.NewAccordion()

	for _, club := range clubs {
		item := widget.NewAccordionItem(club.Name,
			ClubInfos(club),
		)
		ac.Append(item)
	}

	w.SetContent(container.NewVBox(
		label,
		ac),
	)
}
