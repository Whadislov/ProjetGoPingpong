package myapp

import (
	"fmt"
	"slices"
	"strconv"

	mt "github.com/Whadislov/TTCompanion/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PlayerInfos returns a container that displays all the infos of a player.
func PlayerInfos(player *mt.Player) *fyne.Container {

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

	/*
		// To much space between lines
		firstnameLabel1 := widget.NewLabel(T("firstname") + ":")
		firstnameLabel2 := widget.NewLabel(player.Firstname)
		firstnameContent := container.NewGridWithColumns(2, firstnameLabel1, firstnameLabel2)

		lastnameLabel1 := widget.NewLabel(T("lastname") + ":")
		lastnameLabel2 := widget.NewLabel(player.Lastname)
		lastnameContent := container.NewGridWithColumns(2, lastnameLabel1, lastnameLabel2)

		ageLabel1 := widget.NewLabel(T("age") + ":")
		ageLabel2 := widget.NewLabel(strconv.Itoa(player.Age))
		ageContent := container.NewGridWithColumns(2, ageLabel1, ageLabel2)

		forehandLabel1 := widget.NewLabel(T("forehand") + ":")
		forehandLabel2 := widget.NewLabel(player.Material[0])
		forehandContent := container.NewGridWithColumns(2, forehandLabel1, forehandLabel2)

		backhandLabel1 := widget.NewLabel(T("backhand") + ":")
		backhandLabel2 := widget.NewLabel(player.Material[1])
		backhandContent := container.NewGridWithColumns(2, backhandLabel1, backhandLabel2)

		bladeLabel1 := widget.NewLabel(T("blade") + ":")
		bladeLabel2 := widget.NewLabel(player.Material[2])
		bladeContent := container.NewGridWithColumns(2, bladeLabel1, bladeLabel2)

		rankingLabel1 := widget.NewLabel(T("ranking") + ":")
		rankingLabel2 := widget.NewLabel(strconv.Itoa(player.Ranking))
		rankingContent := container.NewGridWithColumns(2, rankingLabel1, rankingLabel2)

		teamsLabel1 := widget.NewLabel(T("teams") + ":")
		teamsLabel2 := widget.NewLabel(strHelper(teams))
		teamsContent := container.NewGridWithColumns(2, teamsLabel1, teamsLabel2)

		clubsLabel1 := widget.NewLabel(T("clubs") + ":")
		clubsLabel2 := widget.NewLabel(strHelper(clubs))
		clubsContent := container.NewGridWithColumns(2, clubsLabel1, clubsLabel2)
	*/

	leftText := T("firstname") + ":\n" + T("lastname") + ":\n" + T("age") + ":\n" + T("forehand") + ":\n" + T("backhand") + ":\n" + T("blade") + ":\n" + T("ranking") + ":\n" + T("teams") + ":\n" + T("clubs") + ":"
	rightText := player.Firstname + "\n" + player.Lastname + "\n" + strconv.Itoa(player.Age) + "\n" + player.Material[0] + "\n" + player.Material[1] + "\n" + player.Material[2] + "\n" + strconv.Itoa(player.Ranking) + "\n" + strHelper(teams) + "\n" + strHelper(clubs)

	playerInfosContent := container.NewGridWithColumns(2, widget.NewLabel(leftText), widget.NewLabel(rightText))
	return playerInfosContent
}

// PlayerPage sets up the page for displaying player info.
func PlayerPage(db *mt.Database, w fyne.Window, a fyne.App) {

	pageTitle := setTitle(T("players"), 32)
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedPlayers := sortMap(db.Players)

	for _, player := range sortedPlayers {
		item := widget.NewAccordionItem(
			fmt.Sprintf("%v %v", player.Value.Firstname, player.Value.Lastname),
			PlayerInfos(player.Value),
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
