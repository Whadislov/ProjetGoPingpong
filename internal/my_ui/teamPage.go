package myapp

import (
	"fmt"
	"slices"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TeamInfos returns a container that displays all the infos of a team.
func TeamInfos(team *mt.Team) *fyne.Container {

	var textC string
	var textP string

	if len(team.ClubID) == 0 {
		textC = T("team_0_club")
	} else {
		clubs := []string{}
		for _, club := range team.ClubID {
			clubs = append(clubs, club)
		}
		textC = fmt.Sprintln(T("club_with_space") + strHelper(clubs))
	}

	if len(team.PlayerIDs) == 0 {
		textP = T("team_0_player")
	} else {
		// Sort players alphabetically
		players := []string{}
		for _, player := range team.PlayerIDs {
			players = append(players, player)
		}
		slices.Sort(players)

		// string that contains player names to display
		textP = T("players_with_space")
		for _, player := range players {
			textP += fmt.Sprintln(player)
			textP += "		"
		}
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
	pageTitle := setTitle(T("teams"), 32)
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedTeams := SortMap(db.Teams)

	for _, team := range sortedTeams {
		item := widget.NewAccordionItem(team.Value.Name,
			TeamInfos(team.Value),
		)
		ac.Append(item)
	}

	returnToDatabasePageButton := widget.NewButton(T("return_to_database"), func() {
		databasePage := DatabasePage(db, w, a)
		w.SetContent(databasePage)
	})

	w.SetContent(container.NewVBox(
		pageTitle,
		returnToDatabasePageButton,
		ac),
	)
}
