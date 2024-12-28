package myapp

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// currentSelectionPagePtoT sets up the page for selecting teams and players.
func currentSelectionPagePtoT(teamContent *fyne.Container, playerContent *fyne.Container, sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToAddPageButton := widget.NewButton("Return to the add menu", func() {
		AddPage(sqlDB, db, w, a)
	})

	if playerContent == nil {
		content := container.NewVBox(teamContent,
			returnToAddPageButton)
		return content
	} else {
		content := container.NewVBox(
			container.NewGridWithColumns(
				2,
				teamContent,
				playerContent,
			),
			returnToAddPageButton)
		return content
	}
}

// selectionPagePtoT sets up the initial selection page for teams.
func SelectionPagePtoT(sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	teamSelectionPageButton := widget.NewButton("Select a team", func() { w.SetContent(selectTeamPagePtoT(sqlDB, db, w, a)) })
	content := container.NewVBox(teamSelectionPageButton)

	return content
}

// selectTeamPagePtoT sets up the page for selecting a team from the database.
func selectTeamPagePtoT(sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToTeamSelectionPageButton := widget.NewButton("Cancel", func() {
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPagePtoT(sqlDB, db, w, a), nil, sqlDB, db, w, a,
			),
		)
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}

	// "Sort the map of teams" for a better button display
	sortedTeams := SortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value
		teamButton := widget.NewButton(team.Name, func() { w.SetContent(selectedTeamPagePtoT(team, sqlDB, db, w, a)) })
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
func selectedTeamPagePtoT(team *mt.Team, sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	tLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🤝", team.Name))
	pLabel := widget.NewLabel("Player current selection 🏓")

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePtoT(sqlDB, db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectedTeamButton,
	)

	// Now select a player
	selectPlayerButton := widget.NewButton("Select a player", func() {
		w.SetContent(selectPlayerPagePtoT(team, sqlDB, db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectPlayerButton,
	)

	// Now display the whole page, with availability to choose a player
	content := currentSelectionPagePtoT(
		teamContent,
		playerContent,
		sqlDB, db, w, a,
	)

	return content
}

// selectPlayerPagePtoT sets up the page for selecting a player for a given team.
func selectPlayerPagePtoT(team *mt.Team, sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToPlayerSelectionPageButton := widget.NewButton("Return to player selection", func() {
		w.SetContent(selectedTeamPagePtoT(team, sqlDB, db, w, a))
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}
	selectedPlayers := make(map[int]*mt.Player)

	// We should create a player first
	if len(db.Players) == 0 {
		yesButton := widget.NewButton("Yes", func() {
			CreatePage(sqlDB, db, w, a)
		})
		noButton := widget.NewButton("No", func() {
			w.SetContent(FunctionalityPage(sqlDB, db, w, a))
		})

		buttons := container.NewHBox(
			yesButton,
			noButton,
		)

		label := widget.NewLabel("There is currently 0 player available. Do you want to create a new player first ?")
		content := container.NewVBox(
			label,
			buttons,
		)

		w.SetContent(content)
	}

	// "Sort the map of players" for a better button display
	sortedPlayers := SortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		// Check if the player from the database is already in the team's player map. If not we want a button of this player
		if _, ok := team.PlayerIDs[player.ID]; !ok {
			playerButton := widget.NewButton(player.Name, func() {
				selectedPlayers[player.ID] = player
				w.SetContent(selectedPlayerPagePtoT(team, selectedPlayers, sqlDB, db, w, a))
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
func createPlayerButtonsPtoT(team *mt.Team, player *mt.Player, sqlDB *msql.Database, db *mt.Database, selectedPlayers map[int]*mt.Player, selectedPlayerButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
	// User can click on the selected player to remove the player from the selected player list
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		delete(selectedPlayers, player.ID)

		// If there is 0 selected player, we should return to the player selection page
		if len(selectedPlayers) == 0 {
			w.SetContent(selectPlayerPagePtoT(team, sqlDB, db, w, a))
		} else {
			w.SetContent(selectedPlayerPagePtoT(team, selectedPlayers, sqlDB, db, w, a))
		}
	})

	selectedPlayerButtons = append(selectedPlayerButtons, selectedPlayerButton)
	return selectedPlayerButtons

}

// addAnotherPlayerPagePtoT sets up the page for adding another player to the selected team.
func addAnotherPlayerPagePtoT(team *mt.Team, alreadySelectedPlayers map[int]*mt.Player, sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToPlayerSelectionPageButton := widget.NewButton("Cancel", func() {
		w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, sqlDB, db, w, a))
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := SortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		// Check if the player from the database is already in the team's player map. If not we want a button of this player
		if _, ok := team.PlayerIDs[player.ID]; !ok {
			if _, ok := alreadySelectedPlayers[player.ID]; !ok {
				// Check if the player from team's player map is already in selected players. If not we want a button of this player
				playerButton := widget.NewButton(player.Name, func() {
					alreadySelectedPlayers[player.ID] = player
					w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, sqlDB, db, w, a))
				})
				playerButtons = append(playerButtons, playerButton)
			}
		}
	}

	if len(playerButtons) == 0 {
		dialog.ShowInformation("Information", "There is no more players to add", w)
		w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, sqlDB, db, w, a))
	}

	content := container.NewVBox(
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPagePtoT sets up the page for confirming the selected players for a team.
func selectedPlayerPagePtoT(team *mt.Team, selectedPlayers map[int]*mt.Player, sqlDB *msql.Database, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToAddRemovePageButton := widget.NewButton("Return to the add menu", func() {
		AddPage(sqlDB, db, w, a)
	})

	tLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🤝", team.Name))
	//pLabel := widget.NewLabel("Player current selection 🏓")

	// "Sort the map of selected players" for a better button display
	sortedSelectedPlayers := SortMap(selectedPlayers)

	confirmButton := widget.NewButton("Confirm", func() {
		var err error
		playerNames := []string{}
		for _, player := range selectedPlayers {
			// Do the link
			err = mf.AddPlayerToTeam(player, team)
			playerNames = append(playerNames, player.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf("Team %v now has player(s) %v", team.Name, strHelper(playerNames))
		fmt.Println(successMsg)
		dialog.ShowInformation("Success", successMsg, w)

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPagePtoT(sqlDB, db, w, a), nil, sqlDB, db, w, a,
			),
		)
	})

	// Create buttons for each selected player
	selectedPlayerButtons := []fyne.CanvasObject{}
	for _, p := range sortedSelectedPlayers {
		player := p.Value
		selectedPlayerButtons = createPlayerButtonsPtoT(team, player, sqlDB, db, selectedPlayers, selectedPlayerButtons, w, a)
	}

	// Add another player in the player selection
	addAnotherPlayerButton := widget.NewButton("Add another player", func() {
		w.SetContent(addAnotherPlayerPagePtoT(team, selectedPlayers, sqlDB, db, w, a))
	})

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePtoT(sqlDB, db, w, a))
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
