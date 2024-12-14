package myapp

import (
	"image/color"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TeamMenu(w fyne.Window, teams map[int]*mt.Team) {

	label := canvas.NewText("List of teams :", color.Black)
	ac := widget.NewAccordion()

	for _, team := range teams {
		wp := []fyne.CanvasObject{}
		for _, player := range team.PlayerIDs {
			wp = append(wp, widget.NewLabel(player))
		}
		itemp := widget.NewAccordionItem("Show players",
			container.NewVBox(wp...),
		)
		playerAc := widget.NewAccordion(itemp)
		clubs := []string{}
		for _, club := range team.ClubID {
			clubs = append(clubs, club)
		}
		clubStr := "Club: " + strHelper(clubs)

		item := widget.NewAccordionItem(team.Name,
			container.NewVBox(
				widget.NewLabel(clubStr),
				playerAc,
			),
		)

		ac.Append(item)
	}

	w.SetContent(container.NewVBox(
		label,
		ac),
	)
}
