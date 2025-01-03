package myapp

import (
	"fmt"
	"slices"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TeamInfos returns a container that displays all the infos of a team.
func TeamInfos(team *mt.Team) *fyne.Container {
	clubs := []string{}
	for _, club := range team.ClubID {
		clubs = append(clubs, club)
	}
	textC := fmt.Sprintln("Club:		" + strHelper(clubs))

	// Sort players alphabetically
	players := []string{}
	for _, player := range team.PlayerIDs {
		players = append(players, player)
	}
	slices.Sort(players)

	// string that contains player names to display
	textP := "Players:	"
	for _, player := range players {
		textP += fmt.Sprintln(player)
		textP += "		"
	}
	text := textC + textP
	text = text[:len(text)-1] // remove the last \n

	item := container.NewVBox(
		widget.NewLabel(text),
	)

	return item
}

// TeamPage sets up the page for displaying team info.
func TeamPage(db *mt.Database, w fyne.Window, a fyne.App) {
	label := widget.NewLabel("Teams")
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedTeams := SortMap(db.Teams)

	for _, team := range sortedTeams {
		item := widget.NewAccordionItem(team.Value.Name,
			TeamInfos(team.Value),
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
