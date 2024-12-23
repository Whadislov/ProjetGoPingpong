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

// currentSelectionPageTtoP sets up the page for selecting players and teams.
func currentSelectionPageTtoP(playerContent *fyne.Container, teamContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToAddPageButton := widget.NewButton("Return to the add menu", func() {
		AddPage(db, w, a)
	})

	if teamContent == nil {
		content := container.NewVBox(playerContent,
			returnToAddPageButton)
		return content
	} else {
		content := container.NewVBox(
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

	playerSelectionPageTtoPButton := widget.NewButton("Select a player", func() { w.SetContent(selectPlayerPageTtoP(db, w, a)) })
	content := container.NewVBox(playerSelectionPageTtoPButton)

	return content
}

// selectPlayerPageTtoP sets up the page for selecting a player from the database.
func selectPlayerPageTtoP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToPlayerSelectionPageTtoPButton := widget.NewButton("Cancel", func() {
		w.SetContent(
			currentSelectionPageTtoP(
				SelectionPageTtoP(db, w, a), nil, db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := SortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		playerButton := widget.NewButton(player.Name, func() { w.SetContent(selectedPlayerPageTtoP(player, db, w, a)) })
		playerButtons = append(playerButtons, playerButton)
	}
	content := container.NewVBox(
		returnToPlayerSelectionPageTtoPButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPageTtoP sets up the page for a selected player and allows team selection.
func selectedPlayerPageTtoP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	tLabel := widget.NewLabel("Team current selection 🤝")

	// User can click on the selected player to return to the list of player
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPageTtoP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	// Now select a team
	selectTeamButton := widget.NewButton("Select a team", func() {
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

	returnToTeamSelectionPageTtoPButton := widget.NewButton("Return to team selection", func() {
		w.SetContent(selectedPlayerPageTtoP(player, db, w, a))
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}
	selectedTeams := make(map[int]*mt.Team)

	// We should create a team first
	if len(db.Teams) == 0 {
		yesButton := widget.NewButton("Yes", func() {
			CreatePage(db, w, a)
		})
		noButton := widget.NewButton("No", func() {
			w.SetContent(FunctionalityPage(db, w, a))
		})

		buttons := container.NewHBox(
			yesButton,
			noButton,
		)

		label := widget.NewLabel("There is currently 0 team available. Do you want to create a new team first ?")
		content := container.NewVBox(
			label,
			buttons,
		)

		w.SetContent(content)
	}

	// "Sort the map of teams" for a better button display
	sortedTeams := SortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value
		// Check if the team from the database is already in the player's team map. If not we want a button of this team
		if _, ok := player.TeamIDs[team.ID]; !ok {
			teamButton := widget.NewButton(team.Name, func() {
				selectedTeams[team.ID] = team
				w.SetContent(selectedTeamPageTtoP(player, selectedTeams, db, w, a))
			})
			teamButtons = append(teamButtons, teamButton)
		}
	}
	content := container.NewVBox(
		returnToTeamSelectionPageTtoPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// createTeamButtonsTtoP creates buttons for each selected team.
func createTeamButtonsTtoP(player *mt.Player, team *mt.Team, db *mt.Database, selectedTeams map[int]*mt.Team, selectedTeamButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
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
func addAnotherTeamPageTtoP(player *mt.Player, alreadySelectedTeams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToTeamSelectionPageTtoPButton := widget.NewButton("Cancel", func() {
		w.SetContent(selectedTeamPageTtoP(player, alreadySelectedTeams, db, w, a))
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}

	// "Sort the map of teams" for a better button display
	sortedTeams := SortMap(db.Teams)

	for _, t := range sortedTeams {
		team := t.Value
		// Check if the team from the database is already in the player's team map. If not we want a button of this team
		if _, ok := player.TeamIDs[team.ID]; !ok {
			if _, ok := alreadySelectedTeams[team.ID]; !ok {
				// Check if the team from player's team map is already in selected teams. If not we want a button of this team
				teamButton := widget.NewButton(team.Name, func() {
					alreadySelectedTeams[team.ID] = team
					w.SetContent(selectedTeamPageTtoP(player, alreadySelectedTeams, db, w, a))
				})
				teamButtons = append(teamButtons, teamButton)
			}
		}
	}

	if len(teamButtons) == 0 {
		dialog.ShowInformation("Information", "There is no more team to add", w)
		w.SetContent(selectedTeamPageTtoP(player, alreadySelectedTeams, db, w, a))
	}

	content := container.NewVBox(
		returnToTeamSelectionPageTtoPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPageTtoP sets up the page for confirming the selected teams for a player.
func selectedTeamPageTtoP(player *mt.Player, selectedTeams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToAddRemovePageButton := widget.NewButton("Return to the add menu", func() {
		AddPage(db, w, a)
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
			err = mf.AddPlayerToTeam(player, team)
			teamNames = append(teamNames, team.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf("Player %v now plays in team(s) %v", player.Name, strHelper(teamNames))
		fmt.Println(successMsg)
		dialog.ShowInformation("Succes", successMsg, w)

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
	addAnotherTeamButton := widget.NewButton("Add another team", func() {
		w.SetContent(addAnotherTeamPageTtoP(player, selectedTeams, db, w, a))
	})

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(player.Name, func() {
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
