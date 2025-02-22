package myapp

import (
	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// currentSelectionPagePfromT sets up the page for selecting teams and players.
func currentSelectionPagePfromT(teamContent *fyne.Container, playerContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToRemovePageButton := widget.NewButton(T("return_to_remove_menu"), func() {
		RemovePage(db, w, a)
	})

	if playerContent == nil {
		content := container.NewVBox(
			teamContent,
			returnToRemovePageButton)
		return content
	} else {
		content := container.NewVBox(
			container.NewGridWithColumns(
				2,
				teamContent,
				playerContent,
			),
			returnToRemovePageButton)
		return content
	}
}

// selectionPagePfromT sets up the initial selection page for teams.
func SelectionPagePfromT(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	playersInTeam := 0
	for _, team := range db.Teams {
		if len(team.PlayerIDs) > 0 {
			playersInTeam++
		}
	}

	if playersInTeam > 0 {
		pageTitle := setTitle(T("remove_select_a_team"), 32)
		teamSelectionPageButton := widget.NewButton(T("select_a_team"), func() { w.SetContent(selectTeamPagePfromT(db, w, a)) })
		return container.NewVBox(
			pageTitle,
			teamSelectionPageButton)
	} else {
		pageTitle := setTitle(T("remove"), 32)
		return container.NewVBox(
			pageTitle,
			widget.NewLabel(T("currently_0_player_available")))
	}
}

// selectTeamPagePfromT sets up the page for selecting a team from the database.
func selectTeamPagePfromT(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("remove_select_a_team"), 32)

	returnToTeamSelectionPageButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(
			currentSelectionPagePfromT(
				SelectionPagePfromT(db, w, a), nil, db, w, a,
			),
		)
	})

	tLabel := widget.NewLabel(T("team_with_hands_emoji"))
	teamButtons := []fyne.CanvasObject{}

	// "Sort the map of teams" for a better button display
	sortedTeams := sortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value

		// If the team is empty, we don't want a button of this team
		if len(team.PlayerIDs) == 0 {
			continue
		} else {
			teamButton := widget.NewButton(team.Name, func() { w.SetContent(selectedTeamPagePfromT(team, db, w, a)) })
			teamButtons = append(teamButtons, teamButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToTeamSelectionPageButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPagePfromT sets up the page for a selected team and allows player selection.
func selectedTeamPagePfromT(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("remove_select_a_player"), 32)

	tLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+"  %v 🤝", team.Name))
	pLabel := widget.NewLabel(T("player_current_selection_emoji"))

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePfromT(db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectedTeamButton,
	)

	if len(team.PlayerIDs) == 0 {
		dialog.ShowInformation(T("information"), fmt.Sprintf("%v "+T("is_empty"), team.Name), w)
		return selectTeamPagePfromT(db, w, a)
	}

	// Now select a player
	selectPlayerButton := widget.NewButton(T("select_a_player"), func() {
		w.SetContent(selectPlayerPagePfromT(team, db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectPlayerButton,
	)

	// Now display the whole page, with availability to choose a player
	content := container.NewVBox(
		pageTitle,
		currentSelectionPagePfromT(
			teamContent,
			playerContent,
			db, w, a,
		),
	)

	return content
}

// selectPlayerPagePfromT sets up the page for selecting a player for a given team.
func selectPlayerPagePfromT(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("remove_select_a_team"), 32)

	returnToPlayerSelectionPageButton := widget.NewButton(T("return_to_player_selection"), func() {
		w.SetContent(selectedTeamPagePfromT(team, db, w, a))
	})

	pLabel := widget.NewLabel(T("players_with_racket_emoji"))
	playerButtons := []fyne.CanvasObject{}
	selectedPlayers := make(map[uuid.UUID]*mt.Player)

	// Nothing to remove
	if len(db.Players) == 0 {
		okButton := widget.NewButton("Ok", func() {
			CreatePage(db, w, a)
		})

		label := widget.NewLabel(T("currently_0_player_available"))
		content := container.NewVBox(
			pageTitle,
			label,
			okButton,
		)

		w.SetContent(content)
	}

	// "Sort the map of players" for a better button display
	sortedPlayers := sortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		// Check if the player from the database is already in the team's player map. If yes we want a button of this player
		if _, ok := team.PlayerIDs[player.ID]; ok {
			playerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
				selectedPlayers[player.ID] = player
				w.SetContent(selectedPlayerPagePfromT(team, selectedPlayers, db, w, a))
			})
			playerButtons = append(playerButtons, playerButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// createPlayerButtonsPfromT creates buttons for each selected player.
func createPlayerButtonsPfromT(team *mt.Team, player *mt.Player, db *mt.Database, selectedPlayers map[uuid.UUID]*mt.Player, selectedPlayerButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
	// User can click on the selected player to remove the player from the selected player list
	selectedPlayerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
		delete(selectedPlayers, player.ID)

		// If there is 0 selected player, we should return to the player selection page
		if len(selectedPlayers) == 0 {
			w.SetContent(selectPlayerPagePfromT(team, db, w, a))
		} else {
			w.SetContent(selectedPlayerPagePfromT(team, selectedPlayers, db, w, a))
		}
	})

	selectedPlayerButtons = append(selectedPlayerButtons, selectedPlayerButton)
	return selectedPlayerButtons

}

// addAnotherPlayerPagePfromT sets up the page for adding another player to the selected team.
func addAnotherPlayerPagePfromT(team *mt.Team, alreadySelectedPlayers map[uuid.UUID]*mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToPlayerSelectionPageButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(selectedPlayerPagePfromT(team, alreadySelectedPlayers, db, w, a))
	})

	pLabel := widget.NewLabel(T("players_with_racket_emoji"))
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := sortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		// Check if the player from the database is already in the team's player map. If yes we want a button of this player
		if _, ok := team.PlayerIDs[player.ID]; ok {
			if _, ok := alreadySelectedPlayers[player.ID]; !ok {
				// Check if the player from team's player map is already in selected players. If not we want a button of this player
				playerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
					alreadySelectedPlayers[player.ID] = player
					w.SetContent(selectedPlayerPagePfromT(team, alreadySelectedPlayers, db, w, a))
				})
				playerButtons = append(playerButtons, playerButton)
			}
		}
	}

	if len(playerButtons) == 0 {
		dialog.ShowInformation(T("information"), T("there_is_no_more_player_to_add"), w)
		w.SetContent(selectedPlayerPagePfromT(team, alreadySelectedPlayers, db, w, a))
	}

	content := container.NewVBox(
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPagePfromT sets up the page for confirming the selected players for a team.
func selectedPlayerPagePfromT(team *mt.Team, selectedPlayers map[uuid.UUID]*mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("remove_confirm_selection"), 32)

	returnToRemovePageButton := widget.NewButton(T("return_to_remove_menu"), func() {
		AddPage(db, w, a)
	})

	tLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+"  %v 🤝", team.Name))
	//pLabel := widget.NewLabel(T("player_current_selection_emoji") )

	// "Sort the map of selected players" for a better button display
	sortedSelectedPlayers := sortMap(selectedPlayers)

	confirmButton := widget.NewButton(T("confirm"), func() {
		var err error
		playerNames := []string{}
		for _, player := range selectedPlayers {
			// Do the link
			err = mf.RemovePlayerFromTeam(player, team)
			playerNames = append(playerNames, fmt.Sprintf("%v %v", player.Firstname, player.Lastname))
			if err != nil {
				dialog.ShowError(err, w)
			}
		}
		if len(playerNames) > 1 {
			successMsg := fmt.Sprintf("%v "+T("are_no_longer_playing_in")+" %v", strHelper(playerNames), team.Name)
			fmt.Println(successMsg)
			dialog.ShowInformation(T("success"), successMsg, w)
		} else {
			successMsg := fmt.Sprintf("%v "+T("is_no_longer_playing_in")+" %v", strHelper(playerNames), team.Name)
			fmt.Println(successMsg)
			dialog.ShowInformation(T("success"), successMsg, w)
		}

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPagePfromT(
				SelectionPagePfromT(db, w, a), nil, db, w, a,
			),
		)
	})

	// Create buttons for each selected player
	selectedPlayerButtons := []fyne.CanvasObject{}
	for _, p := range sortedSelectedPlayers {
		player := p.Value
		selectedPlayerButtons = createPlayerButtonsPfromT(team, player, db, selectedPlayers, selectedPlayerButtons, w, a)
	}

	// Add another player in the player selection
	addAnotherPlayerButton := widget.NewButton(T("add_another_player"), func() {
		w.SetContent(addAnotherPlayerPagePfromT(team, selectedPlayers, db, w, a))
	})

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePfromT(db, w, a))
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
		pageTitle,
		container.NewGridWithColumns(
			2,
			teamContent,
			playerContent,
		),
		confirmButton,
		returnToRemovePageButton,
	)

	return content
}
