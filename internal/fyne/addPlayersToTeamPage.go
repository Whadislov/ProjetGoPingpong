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

// currentSelectionPagePtoT sets up the page for selecting teams and players.
func currentSelectionPagePtoT(teamContent *fyne.Container, playerContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToAddRemovePageButton := widget.NewButton("Return to add... remove...", func() {
		AddPage(db, w, a)
	})

	content := container.NewVBox(
		container.NewGridWithColumns(
			2,
			teamContent,
			playerContent,
		),
		returnToAddRemovePageButton)

	return content
}

// selectionPagePtoT sets up the initial selection page for teams.
func SelectionPagePtoT(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	teamSelectionPageButton := widget.NewButton("Select a team", func() { w.SetContent(selectTeamPagePtoT(db, w, a)) })
	content := container.NewVBox(teamSelectionPageButton)

	return content
}

// selectTeamPagePtoT sets up the page for selecting a team from the database.
func selectTeamPagePtoT(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToTeamSelectionPageButton := widget.NewButton("Cancel", func() {
		w.SetContent(
			currentSelectionPagePtoT(
				waitForTeamSelectionPagePtoT(), SelectionPagePtoT(db, w, a), db, w, a,
			),
		)
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}

	for _, team := range db.Teams {
		teamButton := widget.NewButton(team.Name, func() { w.SetContent(selectedTeamPagePtoT(team, db, w, a)) })
		teamButtons = append(teamButtons, teamButton)
	}
	content := container.NewVBox(
		returnToTeamSelectionPageButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPagePtoT sets up the page for a selected team and allows player selection.
func selectedTeamPagePtoT(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	tLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🤝", team.Name))
	pLabel := widget.NewLabel("Player current selection 🏓")

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePtoT(db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectedTeamButton,
	)

	// Now select a player
	selectPlayerButton := widget.NewButton("Select a player", func() {
		w.SetContent(selectPlayerPagePtoT(team, db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectPlayerButton,
	)

	// Now display the whole page, with availability to choose a player
	content := currentSelectionPagePtoT(
		teamContent,
		playerContent,
		db, w, a,
	)

	return content
}

// waitForTeamSelectionPagePtoT sets up a placeholder page prompting the user to select a team first.
func waitForTeamSelectionPagePtoT() *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("First, select a team"),
	)
	return content
}

// selectPlayerPagePtoT sets up the page for selecting a player for a given team.
func selectPlayerPagePtoT(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToPlayerSelectionPageButton := widget.NewButton("Return to player selection", func() {
		w.SetContent(selectedTeamPagePtoT(team, db, w, a))
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}
	selectedPlayers := make(map[int]*mt.Player)

	for _, player := range db.Players {
		// Check if the player from the database is already in the team's player map. If not we want a button of this player
		if _, ok := team.PlayerIDs[player.ID]; !ok {

			playerButton := widget.NewButton(player.Name, func() {
				selectedPlayers[player.ID] = player
				w.SetContent(selectedPlayerPagePtoT(team, selectedPlayers, db, w, a))
			})
			playerButtons = append(playerButtons, playerButton)
		}
	}
	content := container.NewVBox(
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// createPlayerButtonsPtoT creates buttons for each selected player.
func createPlayerButtonsPtoT(team *mt.Team, player *mt.Player, db *mt.Database, selectedPlayerButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {

	// Remove the selected player from the player list

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPagePtoT(team, db, w, a))
	})

	selectedPlayerButtons = append(selectedPlayerButtons, selectedPlayerButton)
	return selectedPlayerButtons

}

// addAnotherPlayerPagePtoT sets up the page for adding another player to the selected team.
func addAnotherPlayerPagePtoT(team *mt.Team, alreadySelectedPlayers map[int]*mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToPlayerSelectionPageButton := widget.NewButton("Return to player selection", func() {
		w.SetContent(selectedTeamPagePtoT(team, db, w, a))
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	for _, player := range db.Players {
		// Check if the player from the database is already in the team's player map. If not we want a button of this player
		if _, ok := team.PlayerIDs[player.ID]; !ok {
			if _, ok := alreadySelectedPlayers[player.ID]; !ok {
				// Check if the player from team's player map is already in selected players. If not we want a button of this player
				playerButton := widget.NewButton(player.Name, func() {
					alreadySelectedPlayers[player.ID] = player
					w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, db, w, a))
				})
				playerButtons = append(playerButtons, playerButton)
			}
		}
	}

	content := container.NewVBox(
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPagePtoT sets up the page for confirming the selected players for a team.
func selectedPlayerPagePtoT(team *mt.Team, players map[int]*mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToAddRemovePageButton := widget.NewButton("Return to add... remove...", func() {
		AddPage(db, w, a)
	})

	tLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🤝", team.Name))
	//pLabel := widget.NewLabel("Player current selection 🏓")

	confirmButton := widget.NewButton("Confirm", func() {
		var err error
		playerNames := []string{}
		for _, player := range players {
			// Do the link
			err = mf.AddPlayerToTeam(player, team)
			playerNames = append(playerNames, player.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf("Team %v now has now player(s) %v", team.Name, strHelper(playerNames))
		fmt.Println(successMsg)
		dialog.ShowInformation("Success", successMsg, w)

		// Return to empty page
		w.SetContent(
			currentSelectionPagePtoT(
				waitForTeamSelectionPagePtoT(), SelectionPagePtoT(db, w, a), db, w, a,
			),
		)
	})

	// Init
	selectedPlayerButtons := []fyne.CanvasObject{}
	for _, player := range players {
		selectedPlayerButtons = createPlayerButtonsPtoT(team, player, db, selectedPlayerButtons, w, a)
	}

	fmt.Println("Player buttons: ", selectedPlayerButtons)

	addAnotherPlayerButton := widget.NewButton("Add another player", func() {
		// Remove the selected player from the player list
		w.SetContent(addAnotherPlayerPagePtoT(team, players, db, w, a))
	})

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePtoT(db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectedTeamButton,
	)

	playerContent := container.NewVBox(
		addAnotherPlayerButton,
		container.NewVBox(selectedPlayerButtons...),
	)

	// Now display the whole finished page, with chosen players
	content := container.NewVBox(
		container.NewGridWithColumns(
			2,
			teamContent,
			playerContent,
		),
		confirmButton,
		returnToAddRemovePageButton,
	)

	return content
}
