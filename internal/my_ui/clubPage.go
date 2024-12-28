package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// ClubInfos returns a container that has an accordion to show the players in the club, and a second accodion to show the teams in the club.
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

	content := container.NewVBox(
		teamAc,
		playerAc,
	)
	return content
}

// ClubPage sets up the page for displaying players and teams of a club.
func ClubPage(sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) {
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
		databasePage := DatabasePage(sqlDB, db, w, a)
		w.SetContent(databasePage)
	})

	w.SetContent(container.NewVBox(
		returnToDatabasePageButton,
		label,
		ac),
	)
}
