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

	returnToAddRemovePageButton := widget.NewButton("Return to add... remove...", func() {
		AddPage(db, w, a)
	})

	content := container.NewVBox(
		container.NewGridWithColumns(
			2,
			playerContent,
			teamContent,
		),
		returnToAddRemovePageButton)

	return content
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
				SelectionPageTtoP(db, w, a), waitForPlayerSelectionPageTtoP(), db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	for _, player := range db.Players {
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

// waitForPlayerSelectionPageTtoP sets up a placeholder page prompting the user to select a player first.
func waitForPlayerSelectionPageTtoP() *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("First, select a player"),
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
	if len(db.Teams) < 1 {
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

	for _, team := range db.Teams {
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
func createTeamButtonsTtoP(player *mt.Player, team *mt.Team, db *mt.Database, selectedTeamButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {

	// Remove the selected team from the team list

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPageTtoP(player, db, w, a))
	})

	selectedTeamButtons = append(selectedTeamButtons, selectedTeamButton)
	return selectedTeamButtons

}

// addAnotherTeamPageTtoP sets up the page for adding another team to the selected player.
func addAnotherTeamPageTtoP(player *mt.Player, alreadySelectedTeams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToTeamSelectionPageTtoPButton := widget.NewButton("Return to team selection", func() {
		w.SetContent(selectedPlayerPageTtoP(player, db, w, a))
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}

	for _, team := range db.Teams {
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

	content := container.NewVBox(
		returnToTeamSelectionPageTtoPButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

// selectedTeamPageTtoP sets up the page for confirming the selected teams for a player.
func selectedTeamPageTtoP(player *mt.Player, teams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToAddRemovePageButton := widget.NewButton("Return to add... remove...", func() {
		AddPage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	//tLabel := widget.NewLabel("Team current selection 🤝")

	confirmButton := widget.NewButton("Confirm", func() {
		var err error
		teamNames := []string{}
		for _, team := range teams {
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
				SelectionPageTtoP(db, w, a), waitForPlayerSelectionPageTtoP(), db, w, a,
			),
		)
	})

	// Init
	selectedTeamButtons := []fyne.CanvasObject{}
	for _, team := range teams {
		selectedTeamButtons = createTeamButtonsTtoP(player, team, db, selectedTeamButtons, w, a)
	}

	fmt.Println("Team buttons: ", selectedTeamButtons)

	addAnotherTeamButton := widget.NewButton("Add another team", func() {
		// Remove the selected team from the team list
		w.SetContent(addAnotherTeamPageTtoP(player, teams, db, w, a))
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
