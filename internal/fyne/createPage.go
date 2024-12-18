package myapp

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func CreatePage(db *mt.Database, w fyne.Window, a fyne.App) {

	ReturnToFonctionalityPageButton := widget.NewButton("Return to functionalities", func() {
		fonctionalityPage := FunctionalityPage(db, w, a)
		w.SetContent(fonctionalityPage)
		w.Show()
	})

	createPlayerButton := widget.NewButton("Create a new player", func() {
		playerNameEntry := widget.NewEntry()
		playerNameEntry.SetPlaceHolder("Enter your player name here...")

		validatationButton := widget.NewButton("Create", func() {
			playerName := playerNameEntry.Text
			_, err := mf.NewPlayer(playerName, db)
			if err != nil {
				dialog.ShowError(err, w)
				return
			} else {
				successMsg := fmt.Sprintf("Player %v has been successfully created\n", playerName)
				fmt.Println(successMsg)
				dialog.ShowInformation("Succes", successMsg, w)
			}
		})

		w.SetContent(container.NewVBox(
			playerNameEntry,
			validatationButton,
		))
	},
	)
	createTeamButton := widget.NewButton("Create a new team", func() {
		teamNameEntry := widget.NewEntry()
		teamNameEntry.SetPlaceHolder("Enter your team name here...")
		w.SetContent(teamNameEntry)
		teamName := teamNameEntry.SelectedText()
		_, err := mf.NewTeam(teamName, db)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			successMsg := fmt.Sprintf("Team %v has been successfully created\n", teamName)
			fmt.Println(successMsg)
			dialog.ShowInformation("Succes", successMsg, w)
		}
	},
	)
	createClubButton := widget.NewButton("Create a new club", func() {
		clubNameEntry := widget.NewEntry()
		clubNameEntry.SetPlaceHolder("Enter your club name here...")
		w.SetContent(clubNameEntry)
		clubName := clubNameEntry.SelectedText()
		_, err := mf.NewClub(clubName, db)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			successMsg := fmt.Sprintf("Club %v has been successfully created\n", clubName)
			fmt.Println(successMsg)
			dialog.ShowInformation("Succes", successMsg, w)
		}
	},
	)

	w.SetContent(
		container.NewVBox(
			ReturnToFonctionalityPageButton,
			createPlayerButton,
			createTeamButton,
			createClubButton,
		),
	)
}
