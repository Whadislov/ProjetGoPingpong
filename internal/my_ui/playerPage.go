package myapp

import (
	"fmt"
	"slices"
	"strconv"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PlayerInfos returns a container that displays all the infos of a player.
func PlayerInfos(player *mt.Player) *fyne.Container {
	firstnameStr := "Firstname: 	" + player.Firstname
	LastnameStr := "Lastname: 	" + player.Lastname
	ageStr := "Age: 		" + strconv.Itoa(player.Age)
	materialStr := fmt.Sprintf(
		`Forehand:	%v
Backhand:	%v
Blade:	%v`,
		player.Material[0], player.Material[1], player.Material[2])

	rankingStr := "Ranking: 	" + strconv.Itoa(player.Ranking)
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
	teamsStr := "Teams: 	" + strHelper(teams)
	clubsStr := "Clubs: 	" + strHelper(clubs)

	text := fmt.Sprintln(firstnameStr) + fmt.Sprintln(LastnameStr) + fmt.Sprintln(ageStr) + fmt.Sprintln(rankingStr) + fmt.Sprintln(teamsStr) + fmt.Sprintln(clubsStr) + fmt.Sprintln(materialStr)
	text = text[:len(text)-1] // remove the last \n
	textLabel := widget.NewLabel(text)
	item := container.NewVBox(textLabel)
	return item
}

// PlayerPage sets up the page for displaying player info.
func PlayerPage(db *mt.Database, w fyne.Window, a fyne.App) {

	pageTitle := setTitle("Players", 32)
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedPlayers := SortMap(db.Players)

	for _, player := range sortedPlayers {
		item := widget.NewAccordionItem(
			fmt.Sprintf("%v %v", player.Value.Firstname, player.Value.Lastname),
			PlayerInfos(player.Value),
		)

		ac.Append(item)
	}

	returnToDatabasePageButton := widget.NewButton("Return to database", func() {
		databasePage := DatabasePage(db, w, a)
		w.SetContent(databasePage)
	})

	w.SetContent(container.NewVBox(
		pageTitle,
		returnToDatabasePageButton,
		ac),
	)
}
