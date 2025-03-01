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

// currentSelectionPageTtoP sets up the page for selecting players and teams.
func currentSelectionPageTtoP(playerContent *fyne.Container, teamContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_team"), 32)

	returnToAddPageButton := widget.NewButton(T("return_to_the_add_menu"), func() {
		AddPage(db, w, a)
	})

	if teamContent == nil {
		content := container.NewVBox(playerContent,
			returnToAddPageButton)
		return content
	} else {
		content := container.NewVBox(
			pageTitle,
			container.NewGridWithColumns(
				2,
				playerContent,
				teamContent,
			),
			returnToAddPageButton)
		return content
	}
}

// SelectionPageTtoP sets up the initial selection page for players.
func SelectionPageTtoP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_player"), 32)

	playerSelectionPageTtoPButton := widget.NewButton(T("select_a_player"), func() { w.SetContent(selectPlayerPageTtoP(db, w, a)) })
	content := container.NewVBox(
		pageTitle,
		playerSelectionPageTtoPButton)

	return content
}

// selectPlayerPageTtoP sets up the page for selecting a player from the database.
func selectPlayerPageTtoP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("add_select_a_player"), 32)

	returnToPlayerSelectionPageTtoPButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(
			currentSelectionPageTtoP(
				SelectionPageTtoP(db, w, a), nil, db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel(T("players_with_racket_emoji"))
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := sortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		playerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() { w.SetContent(selectedPlayerPageTtoP(player, db, w, a)) })
		playerButtons = append(playerButtons, playerButton)
	}
	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageTtoPButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPageTtoP sets up the page for a selected player and allows team selection.
func selectedPlayerPageTtoP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+" %v 🏓", fmt.Sprintf("%v %v", player.Firstname, player.Lastname)))
	tLabel := widget.NewLabel(T("team_current_selection_emoji"))

	// User can click on the selected player to return to the list of player
	selectedPlayerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
		w.SetContent(selectPlayerPageTtoP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	// Now select a team
	selectTeamButton := widget.NewButton(T("select_a_team"), func() {
		w.SetContent(selectTeamPageTtoP(player, db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectTeamButton,
	)

	// Now display the whole page, with availability to choose a team
	content := currentSelectionPageTtoP(
		playerContent,
		teamContent,
		db, w, a,
	)

	return content
}

// selectTeamPageTtoP sets up the page for selecting a team for a given player.
func selectTeamPageTtoP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_team"), 32)

	returnToTeamSelectionPageTtoPButton := widget.NewButton(T("return_to_team_selection"), func() {
		w.SetContent(selectedPlayerPageTtoP(player, db, w, a))
	})

	tLabel := widget.NewLabel(T("team_with_hands_emoji"))
	teamButtons := []fyne.CanvasObject{}
	selectedTeams := make(map[uuid.UUID]*mt.Team)

	// We should create a team first
	if len(db.Teams) == 0 {
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

		label := widget.NewLabel(T("there_is_currently_0_team_available_do_you_want_create"))
		content := container.NewVBox(
			pageTitle,
			label,
			buttons,
		)

		w.SetContent(content)
	}

	// We do not want to have a button for all the teams in the database. They have to meet some criterias : same club as the player, player not already playing in
	screenedTeams := make(map[uuid.UUID]*mt.Team)
	// First get the club(s)
	for clubID := range player.ClubIDs {
		club, _ := db.GetClub(clubID)
		// Then get the team(s) from the club(s)
		for teamID := range club.TeamIDs {
			// Check if the team is not already in the player's team map.
			_, ok := player.TeamIDs[teamID]
			if !ok {
				screenedTeams[teamID] = db.Teams[teamID]
			}
		}
	}

	// "Sort the map of teams" for a better button display
	sortedTeams := sortMap(screenedTeams)

	for _, t := range sortedTeams {
		team := t.Value
		teamButton := widget.NewButton(team.Name, func() {
			selectedTeams[team.ID] = team
			w.SetContent(selectedTeamPageTtoP(player, selectedTeams, db, w, a))
		})
		teamButtons = append(teamButtons, teamButton)

	}
	content := container.NewVBox(
		pageTitle,
		returnToTeamSelectionPageTtoPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// createTeamButtonsTtoP creates buttons for each selected team.
func createTeamButtonsTtoP(player *mt.Player, team *mt.Team, db *mt.Database, selectedTeams map[uuid.UUID]*mt.Team, selectedTeamButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
	// User can click on the selected team to remove the team from the selected team list
	selectedTeamButton := widget.NewButton(team.Name, func() {
		delete(selectedTeams, team.ID)

		// If there is 0 selected team, we should return to the team selection page
		if len(selectedTeams) == 0 {
			w.SetContent(selectTeamPageTtoP(player, db, w, a))
		} else {
			w.SetContent(selectedTeamPageTtoP(player, selectedTeams, db, w, a))
		}
	})

	selectedTeamButtons = append(selectedTeamButtons, selectedTeamButton)
	return selectedTeamButtons

}

// addAnotherTeamPageTtoP sets up the page for adding another team to the selected player.
func addAnotherTeamPageTtoP(player *mt.Player, alreadySelectedTeams map[uuid.UUID]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_select_a_team"), 32)

	returnToTeamSelectionPageTtoPButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(selectedTeamPageTtoP(player, alreadySelectedTeams, db, w, a))
	})

	tLabel := widget.NewLabel(T("team_with_hands_emoji"))
	teamButtons := []fyne.CanvasObject{}

	// We do not want to have a button for all the teams in the database. They have to meet some criterias : same club as the player, player not already playing in, not already given a button
	screenedTeams := make(map[uuid.UUID]*mt.Team)
	// First get the club(s)
	for clubID := range player.ClubIDs {
		club, _ := db.GetClub(clubID)
		// Then get the team(s) from the club(s)
		for teamID := range club.TeamIDs {
			// Check if the team is not already in the player's team map. If not, continue checking if the team is not already in selected teams
			_, ok := player.TeamIDs[teamID]
			if !ok {
				// Check if the team is not already in selected teams. If not, keep this team
				if _, ok := alreadySelectedTeams[teamID]; !ok {

					screenedTeams[teamID] = db.Teams[teamID]
				}
			}
		}
	}

	// "Sort the map of teams" for a better button display
	sortedTeams := sortMap(screenedTeams)

	for _, t := range sortedTeams {
		team := t.Value

		teamButton := widget.NewButton(team.Name, func() {
			alreadySelectedTeams[team.ID] = team
			w.SetContent(selectedTeamPageTtoP(player, alreadySelectedTeams, db, w, a))
		})
		teamButtons = append(teamButtons, teamButton)

	}

	if len(teamButtons) == 0 {
		dialog.ShowInformation(T("information"), T("there_is_no_more_team_to_add"), w)
		w.SetContent(selectedTeamPageTtoP(player, alreadySelectedTeams, db, w, a))
	}

	content := container.NewVBox(
		pageTitle,
		returnToTeamSelectionPageTtoPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPageTtoP sets up the page for confirming the selected teams for a player.
func selectedTeamPageTtoP(player *mt.Player, selectedTeams map[uuid.UUID]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("add_confirm"), 32)

	returnToAddRemovePageButton := widget.NewButton(T("return_to_the_add_menu"), func() {
		AddPage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+" %v 🏓", fmt.Sprintf("%v %v", player.Firstname, player.Lastname)))

	// "Sort the map of selectedTeams" for a better button display
	sortedSelectedTeams := sortMap(selectedTeams)

	confirmButton := widget.NewButton(T("confirm"), func() {
		var err error
		teamNames := []string{}
		for _, t := range sortedSelectedTeams {
			team := t.Value
			// Do the link
			err = mf.AddPlayerToTeam(player, team)
			teamNames = append(teamNames, team.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf(T("player_now_plays_in_team"), fmt.Sprintf("%v %v", player.Firstname, player.Lastname), strHelper(teamNames))
		fmt.Println(successMsg)
		dialog.ShowInformation(T("success"), successMsg, w)

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPageTtoP(
				SelectionPageTtoP(db, w, a), nil, db, w, a,
			),
		)
	})

	// Create buttons for each selected team
	selectedTeamButtons := []fyne.CanvasObject{}
	for _, t := range sortedSelectedTeams {
		team := t.Value
		selectedTeamButtons = createTeamButtonsTtoP(player, team, db, selectedTeams, selectedTeamButtons, w, a)
	}

	// Add another team in the team selection
	addAnotherTeamButton := widget.NewButton(T("add_another_team"), func() {
		w.SetContent(addAnotherTeamPageTtoP(player, selectedTeams, db, w, a))
	})

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
		w.SetContent(selectPlayerPageTtoP(db, w, a))
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
		returnToAddRemovePageButton,
	)

	return content
}
