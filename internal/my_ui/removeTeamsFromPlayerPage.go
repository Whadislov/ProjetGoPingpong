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

// currentSelectionPageTfromP sets up the page for selecting players and teams.
func currentSelectionPageTfromP(playerContent *fyne.Container, teamContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Remove: select a team", 32)

	returnToRemovePageButton := widget.NewButton("Return to the remove menu", func() {
		RemovePage(db, w, a)
	})

	if teamContent == nil {
		content := container.NewVBox(
			pageTitle,
			playerContent,
			returnToRemovePageButton)
		return content
	} else {
		content := container.NewVBox(
			pageTitle,
			container.NewGridWithColumns(
				2,
				playerContent,
				teamContent,
			),
			returnToRemovePageButton)
		return content
	}
}

// SelectionPageTfromP sets up the initial selection page for players.
func SelectionPageTfromP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	playersInTeam := 0
	for _, team := range db.Teams {
		if len(team.PlayerIDs) > 0 {
			playersInTeam++
		}
	}

	if playersInTeam > 0 {
		playerSelectionPageTfromPButton := widget.NewButton("Select a player", func() { w.SetContent(selectPlayerPageTfromP(db, w, a)) })
		return container.NewVBox(playerSelectionPageTfromPButton)
	} else {
		return container.NewVBox(widget.NewLabel("There is currently 0 player in any team"))
	}
}

// selectPlayerPageTfromP sets up the page for selecting a player from the database.
func selectPlayerPageTfromP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Remove: select a player", 32)

	returnToPlayerSelectionPageTfromPButton := widget.NewButton("Cancel", func() {
		w.SetContent(
			currentSelectionPageTfromP(
				SelectionPageTfromP(db, w, a), nil, db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := SortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value

		// If the player's team list is empty, we don't want a button for this player
		if len(player.TeamIDs) == 0 {
			continue
		} else {
			playerButton := widget.NewButton(player.Name, func() { w.SetContent(selectedPlayerPageTfromP(player, db, w, a)) })
			playerButtons = append(playerButtons, playerButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageTfromPButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPageTfromP sets up the page for a selected player and allows team selection.
func selectedPlayerPageTfromP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	tLabel := widget.NewLabel("Team current selection 🤝")

	// User can click on the selected player to return to the list of player
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPageTfromP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	if len(player.TeamIDs) == 0 {
		dialog.ShowInformation("Information", fmt.Sprintf("%v has no team", player.Name), w)
		return selectPlayerPageTfromP(db, w, a)
	}

	// Now select a team
	selectTeamButton := widget.NewButton("Select a team", func() {
		w.SetContent(selectTeamPageTfromP(player, db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectTeamButton,
	)

	// Now display the whole page, with availability to choose a team
	content := currentSelectionPageTfromP(
		playerContent,
		teamContent,
		db, w, a,
	)

	return content
}

// selectTeamPageTfromP sets up the page for selecting a team for a given player.
func selectTeamPageTfromP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Remove: select a team", 32)

	returnToTeamSelectionPageTfromPButton := widget.NewButton("Return to team selection", func() {
		w.SetContent(selectedPlayerPageTfromP(player, db, w, a))
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}
	selectedTeams := make(map[int]*mt.Team)

	// Nothing to remove
	if len(db.Players) == 0 {
		okButton := widget.NewButton("Ok", func() {
			RemovePage(db, w, a)
		})

		label := widget.NewLabel("There is currently 0 team available.")
		content := container.NewVBox(
			pageTitle,
			label,
			okButton,
		)

		w.SetContent(content)
	}

	// "Sort the map of teams" for a better button display
	sortedTeams := SortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value
		// Check if the team from the database is already in the player's team map. If yes we want a button of this team
		if _, ok := player.TeamIDs[team.ID]; ok {
			teamButton := widget.NewButton(team.Name, func() {
				selectedTeams[team.ID] = team
				w.SetContent(selectedTeamPageTfromP(player, selectedTeams, db, w, a))
			})
			teamButtons = append(teamButtons, teamButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToTeamSelectionPageTfromPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// createTeamButtonsTfromP creates buttons for each selected team.
func createTeamButtonsTfromP(player *mt.Player, team *mt.Team, db *mt.Database, selectedTeams map[int]*mt.Team, selectedTeamButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {

	// User can click on the selected team to remove the team from the selected team list
	selectedTeamButton := widget.NewButton(team.Name, func() {
		delete(selectedTeams, team.ID)

		// If there is 0 selected team, we should return to the team selection page
		if len(selectedTeams) == 0 {
			w.SetContent(selectTeamPageTfromP(player, db, w, a))
		} else {
			w.SetContent(selectedTeamPageTfromP(player, selectedTeams, db, w, a))
		}
	})

	selectedTeamButtons = append(selectedTeamButtons, selectedTeamButton)
	return selectedTeamButtons

}

// addAnotherTeamPageTfromP sets up the page for adding another team to the selected player.
func addAnotherTeamPageTfromP(player *mt.Player, alreadySelectedTeams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToTeamSelectionPageTfromPButton := widget.NewButton("Cancel", func() {
		w.SetContent(selectedTeamPageTfromP(player, alreadySelectedTeams, db, w, a))
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}

	// "Sort the map of teams" for a better button display
	sortedTeams := SortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value
		// Check if the team from the database is already in the player's team map. If yes we want a button of this team
		if _, ok := player.TeamIDs[team.ID]; ok {
			if _, ok := alreadySelectedTeams[team.ID]; !ok {
				// Check if the team from player's team map is already in selected teams. If not we want a button of this team
				teamButton := widget.NewButton(team.Name, func() {
					alreadySelectedTeams[team.ID] = team
					w.SetContent(selectedTeamPageTfromP(player, alreadySelectedTeams, db, w, a))
				})
				teamButtons = append(teamButtons, teamButton)
			}
		}
	}

	if len(teamButtons) == 0 {
		dialog.ShowInformation("Information", "There is no more team to add", w)
		w.SetContent(selectedTeamPageTfromP(player, alreadySelectedTeams, db, w, a))
	}

	content := container.NewVBox(
		returnToTeamSelectionPageTfromPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPageTfromP sets up the page for confirming the selected teams for a player.
func selectedTeamPageTfromP(player *mt.Player, selectedTeams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Remove: confirm selection", 32)

	returnToRemovePageButton := widget.NewButton("Return to the remove menu", func() {
		RemovePage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	//tLabel := widget.NewLabel("Team current selection 🤝")

	// "Sort the map of selectedTeams" for a better button display
	sortedSelectedTeams := SortMap(selectedTeams)

	confirmButton := widget.NewButton("Confirm", func() {
		var err error
		teamNames := []string{}
		for _, t := range sortedSelectedTeams {
			team := t.Value
			// Do the link
			err = mf.RemovePlayerFromTeam(player, team)
			teamNames = append(teamNames, team.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf("Player %v no longer plays in team(s) %v", player.Name, strHelper(teamNames))
		fmt.Println(successMsg)
		dialog.ShowInformation("Succes", successMsg, w)

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPageTfromP(
				SelectionPageTfromP(db, w, a), nil, db, w, a,
			),
		)
	})

	// Create buttons for each selected team
	selectedTeamButtons := []fyne.CanvasObject{}
	for _, t := range sortedSelectedTeams {
		team := t.Value
		selectedTeamButtons = createTeamButtonsTfromP(player, team, db, selectedTeams, selectedTeamButtons, w, a)
	}

	// Add another team in the team selection
	addAnotherTeamButton := widget.NewButton("Add another team", func() {
		w.SetContent(addAnotherTeamPageTfromP(player, selectedTeams, db, w, a))
	})

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPageTfromP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	teamContent := container.NewVBox(
		addAnotherTeamButton,
		container.NewVBox(selectedTeamButtons...),
	)

	// Now display the whole finished page, with chosen teams
	content := container.NewVBox(
		pageTitle,
		container.NewGridWithColumns(
			2,
			playerContent,
			teamContent,
		),
		confirmButton,
		returnToRemovePageButton,
	)

	return content
}
