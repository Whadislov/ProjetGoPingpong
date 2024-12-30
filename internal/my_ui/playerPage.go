package myapp

import (
	"slices"
	"strconv"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PlayerInfos returns a container that displays all the infos of a player.
func PlayerInfos(player *mt.Player) *fyne.Container {
	nameStr := "Name: " + player.Name
	ageStr := "Age: " + strconv.Itoa(player.Age)
	materialStr := "Material: " + strHelper(player.Material)
	rankingStr := "Ranking: " + strconv.Itoa(player.Ranking)
	teams := []string{}
	clubs := []string{}
	for _, team := range player.TeamIDs {
		teams = append(teams, team)
	}
	// Sort teams alphabetically
	slices.Sort(teams)
	for _, club := range player.ClubIDs {
		clubs = append(clubs, club)
	}
	// Sort clubs alphabetically
	slices.Sort(clubs)
	teamsStr := "Teams: " + strHelper(teams)
	clubsStr := "Clubs: " + strHelper(clubs)
	item := container.NewVBox(
		widget.NewLabel(nameStr),
		widget.NewLabel(ageStr),
		widget.NewLabel(materialStr),
		widget.NewLabel(rankingStr),
		widget.NewLabel(teamsStr),
		widget.NewLabel(clubsStr),
	)
	return item
}

// PlayerPage sets up the page for displaying player info.
func PlayerPage(db *mt.Database, w fyne.Window, a fyne.App) {

	label := widget.NewLabel("Players")
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedPlayers := SortMap(db.Players)

	for _, player := range sortedPlayers {
		item := widget.NewAccordionItem(
			player.Value.Name,
			PlayerInfos(player.Value),
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
