package myapp

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// ClubInfos returns a container that has an accordion to show the players in the club, and a second accodion to show the teams in the club.
func ClubInfos(club *mt.Club) *fyne.Container {
	// team label list
	wt := []fyne.CanvasObject{}
	// player label list
	wp := []fyne.CanvasObject{}

	// Sort players alphabetically
	players := []string{}
	for _, player := range club.PlayerIDs {
		players = append(players, player)
	}
	slices.Sort(players)

	for _, player := range players {
		wp = append(wp, widget.NewLabel(player))
	}
	itemp := widget.NewAccordionItem("Show players",
		container.NewVBox(wp...),
	)
	playerAc := widget.NewAccordion(itemp)

	// Sort teams alphabetically
	teams := []string{}
	for _, team := range club.TeamIDs {
		teams = append(teams, team)
	}
	slices.Sort(teams)

	for _, team := range teams {
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
	})

	w.SetContent(container.NewVBox(
		returnToDatabasePageButton,
		label,
		ac),
	)
}
