package myapp

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
	"github.com/google/uuid"
)

// currentSelectionPagePtoT sets up the page for selecting teams and players.
func currentSelectionPagePtoT(teamContent *fyne.Container, playerContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_player"), 32)

	returnToAddPageButton := widget.NewButton(T("return_to_the_add_menu"), func() {
		AddPage(db, w, a)
	})

	if playerContent == nil {
		content := container.NewVBox(teamContent,
			returnToAddPageButton)
		return content
	} else {
		content := container.NewVBox(
			pageTitle,
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
func SelectionPagePtoT(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_team"), 32)

	teamSelectionPageButton := widget.NewButton(T("select_a_team"), func() { w.SetContent(selectTeamPagePtoT(db, w, a)) })
	content := container.NewVBox(
		pageTitle,
		teamSelectionPageButton)

	return content
}

// selectTeamPagePtoT sets up the page for selecting a team from the database.
func selectTeamPagePtoT(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("add_select_a_team"), 32)

	returnToTeamSelectionPageButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPagePtoT(db, w, a), nil, db, w, a,
			),
		)
	})

	tLabel := widget.NewLabel(T("team_with_hands_emoji"))
	teamButtons := []fyne.CanvasObject{}

	// "Sort the map of teams" for a better button display
	sortedTeams := sortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value
		teamButton := widget.NewButton(team.Name, func() { w.SetContent(selectedTeamPagePtoT(team, db, w, a)) })
		teamButtons = append(teamButtons, teamButton)
	}
	content := container.NewVBox(
		pageTitle,
		returnToTeamSelectionPageButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPagePtoT sets up the page for a selected team and allows player selection.
func selectedTeamPagePtoT(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	tLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+"  %v 🤝", team.Name))
	pLabel := widget.NewLabel(T("player_current_selection_emoji"))

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPagePtoT(db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectedTeamButton,
	)

	// Now select a player
	selectPlayerButton := widget.NewButton(T("select_a_player"), func() {
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

// selectPlayerPagePtoT sets up the page for selecting a player for a given team.
func selectPlayerPagePtoT(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_player"), 32)

	returnToPlayerSelectionPageButton := widget.NewButton(T("return_to_player_selection"), func() {
		w.SetContent(selectedTeamPagePtoT(team, db, w, a))
	})

	pLabel := widget.NewLabel(T("players_with_racket_emoji"))
	playerButtons := []fyne.CanvasObject{}
	selectedPlayers := make(map[uuid.UUID]*mt.Player)

	// We should create a player first
	if len(db.Players) == 0 {
		yesButton := widget.NewButton(T("yes"), func() {
			CreatePage(db, w, a)
		})
		noButton := widget.NewButton(T("no"), func() {
			w.SetContent(FunctionalityPage(db, w, a))
		})

		buttons := container.NewHBox(
			yesButton,
			noButton,
		)

		label := widget.NewLabel(T("there_is_currently_0_player_available_do_you_want_create"))
		content := container.NewVBox(
			pageTitle,
			label,
			buttons,
		)

		w.SetContent(content)
	}

	// We do not want to have a button for all the players in the database. They have to meet some criterias : same club as the team, not already in the team
	screenedPlayers := make(map[uuid.UUID]*mt.Player)

	for _, player := range db.Players {
		for playerClubID := range player.ClubIDs {
			// Check if the player is in the same club as the team. If yes, continue checking if the player is not already in the team
			_, ok := team.ClubID[playerClubID]
			if ok {
				// Check if the player is not already in the team. If not, we want a button of this player
				_, ok := player.TeamIDs[team.ID]
				if !ok {
					screenedPlayers[player.ID] = player
				}
			}
		}
	}

	// "Sort the map of players" for a better button display
	sortedPlayers := sortMap(screenedPlayers)

	for _, p := range sortedPlayers {
		player := p.Value
		playerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
			selectedPlayers[player.ID] = player
			w.SetContent(selectedPlayerPagePtoT(team, selectedPlayers, db, w, a))
		})
		playerButtons = append(playerButtons, playerButton)

	}
	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// createPlayerButtonsPtoT creates buttons for each selected player.
func createPlayerButtonsPtoT(team *mt.Team, player *mt.Player, db *mt.Database, selectedPlayers map[uuid.UUID]*mt.Player, selectedPlayerButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
	// User can click on the selected player to remove the player from the selected player list
	selectedPlayerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
		delete(selectedPlayers, player.ID)

		// If there is 0 selected player, we should return to the player selection page
		if len(selectedPlayers) == 0 {
			w.SetContent(selectPlayerPagePtoT(team, db, w, a))
		} else {
			w.SetContent(selectedPlayerPagePtoT(team, selectedPlayers, db, w, a))
		}
	})

	selectedPlayerButtons = append(selectedPlayerButtons, selectedPlayerButton)
	return selectedPlayerButtons

}

// addAnotherPlayerPagePtoT sets up the page for adding another player to the selected team.
func addAnotherPlayerPagePtoT(team *mt.Team, alreadySelectedPlayers map[uuid.UUID]*mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_player"), 32)

	returnToPlayerSelectionPageButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, db, w, a))
	})

	pLabel := widget.NewLabel(T("players_with_racket_emoji"))
	playerButtons := []fyne.CanvasObject{}

	// We do not want to have a button for all the players in the database. They have to meet some criterias : same club as the team, not already in the team, not already given a button
	screenedPlayers := make(map[uuid.UUID]*mt.Player)

	for _, player := range db.Players {
		for playerClubID := range player.ClubIDs {
			// Check if the player is in the same club as the team. If yes, continue checking if the player is not already in the team
			_, ok := team.ClubID[playerClubID]
			if ok {
				// Check if the player is not already in the team. If not, continue checking if the player is not already in selected players
				_, ok := player.TeamIDs[team.ID]
				if !ok {
					// Check if the player is not already in selected players. If not, keep this player
					_, ok := alreadySelectedPlayers[player.ID]
					if !ok {
						screenedPlayers[player.ID] = player
					}
				}
			}
		}
	}

	// "Sort the map of players" for a better button display
	sortedPlayers := sortMap(screenedPlayers)

	for _, p := range sortedPlayers {
		player := p.Value

		playerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
			alreadySelectedPlayers[player.ID] = player
			w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, db, w, a))
		})
		playerButtons = append(playerButtons, playerButton)

	}

	if len(playerButtons) == 0 {
		dialog.ShowInformation(T("information"), T("there_is_no_more_player_to_add"), w)
		w.SetContent(selectedPlayerPagePtoT(team, alreadySelectedPlayers, db, w, a))
	}

	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPagePtoT sets up the page for confirming the selected players for a team.
func selectedPlayerPagePtoT(team *mt.Team, selectedPlayers map[uuid.UUID]*mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_confirm"), 32)

	returnToAddRemovePageButton := widget.NewButton(T("return_to_the_add_menu"), func() {
		AddPage(db, w, a)
	})

	tLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+"  %v 🤝", team.Name))

	// "Sort the map of selected players" for a better button display
	sortedSelectedPlayers := sortMap(selectedPlayers)

	confirmButton := widget.NewButton(T("confirm"), func() {
		var err error
		playerNames := []string{}
		for _, player := range selectedPlayers {
			// Do the link
			err = mf.AddPlayerToTeam(player, team)
			playerNames = append(playerNames, fmt.Sprintf("%v %v", player.Firstname, player.Lastname))
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf(T("team_now_has_player"), team.Name, strHelper(playerNames))
		fmt.Println(successMsg)
		dialog.ShowInformation(T("success"), successMsg, w)

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPagePtoT(
				SelectionPagePtoT(db, w, a), nil, db, w, a,
			),
		)
	})

	// Create buttons for each selected player
	selectedPlayerButtons := []fyne.CanvasObject{}
	for _, p := range sortedSelectedPlayers {
		player := p.Value
		selectedPlayerButtons = createPlayerButtonsPtoT(team, player, db, selectedPlayers, selectedPlayerButtons, w, a)
	}

	// Add another player in the player selection
	addAnotherPlayerButton := widget.NewButton(T("add_another_player"), func() {
		w.SetContent(addAnotherPlayerPagePtoT(team, selectedPlayers, db, w, a))
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
		pageTitle,
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
