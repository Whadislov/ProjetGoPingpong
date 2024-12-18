package myapp

import (
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ClubInfos(club *mt.Club) *fyne.Container {
	// team list
	wt := []fyne.CanvasObject{}
	// player list
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

func ClubPage(db *mt.Database, w fyne.Window, a fyne.App) {
	label := widget.NewLabel("Clubs")
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedClubs := SortMap(db.Clubs)

	for _, club := range sortedClubs {
		item := widget.NewAccordionItem(club.Value.Name,
			ClubInfos(club.Value),
		)
		ac.Append(item)
	}

	returnToDatabasePageButton := widget.NewButton("Return to database", func() {
		databasePage := DatabasePage(db, w, a)
		w.SetContent(databasePage)
		w.Show()
	})

	w.SetContent(container.NewVBox(
		returnToDatabasePageButton,
		label,
		ac),
	)
}
