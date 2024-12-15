package myapp

import (
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateMenu(w fyne.Window, db *mt.Database) {

	createPlayerButton := widget.NewButton("Create a new player", func() {
		playerNameEntry := widget.NewEntry()
		playerName := playerNameEntry.SelectedText()
		mf.NewPlayer(playerName, db)
	},
	)
	createTeamButton := widget.NewButton("Create a new team", func() {
		teamNameEntry := widget.NewEntry()
		teamName := teamNameEntry.SelectedText()
		mf.NewTeam(teamName, db)
	},
	)
	createClubButton := widget.NewButton("Create a new club", func() {
		clubNameEntry := widget.NewEntry()
		clubName := clubNameEntry.SelectedText()
		mf.NewClub(clubName, db)
	},
	)

	w.SetContent(container.NewHBox(
		createPlayerButton,
		createTeamButton,
		createClubButton),
	)
}
