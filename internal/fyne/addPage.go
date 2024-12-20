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

// Current selection page

func currentSelectionPage(playerContent *fyne.Container, teamContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

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

// Player function

func selectPlayerPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToPlayerSelectionPageButton := widget.NewButton("Cancel", func() {
		w.SetContent(
			currentSelectionPage(
				selectionPage(db, w, a), waitForPlayerSelectionPage(), db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	for _, player := range db.Players {
		playerButton := widget.NewButton(player.Name, func() { w.SetContent(selectedPlayerPage(player, db, w, a)) })
		playerButtons = append(playerButtons, playerButton)
	}
	content := container.NewVBox(
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

func selectedPlayerPage(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	tLabel := widget.NewLabel("Team current selection 🤝")

	// User can click on the selected player to return the list of player
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPage(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	// Now select a team
	selectTeamButton := widget.NewButton("Select a team", func() {
		w.SetContent(selectTeamPage(player, db, w, a))
	})

	teamContent := container.NewVBox(
		tLabel,
		selectTeamButton,
	)

	// Now display the whole page, with availability to choose a team
	content := currentSelectionPage(
		playerContent,
		teamContent,
		db, w, a,
	)

	return content
}

// Empty team section
func waitForPlayerSelectionPage() *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("First, select a player"),
	)
	return content
}

// Selection page (semi-empty player section)

func selectionPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	playerSelectionPageButton := widget.NewButton("Select a player", func() { w.SetContent(selectPlayerPage(db, w, a)) })
	content := container.NewVBox(playerSelectionPageButton)

	return content
}

func selectTeamPage(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToTeamSelectionPageButton := widget.NewButton("Return to team selection", func() {
		w.SetContent(selectedPlayerPage(player, db, w, a))
	})

	tLabel := widget.NewLabel("Teams 🤝")
	teamButtons := []fyne.CanvasObject{}
	selectedTeams := make(map[int]*mt.Team)

	for _, team := range db.Teams {
		// Check if the team from the database is already in the player's team map. If not we want a button of this team
		if _, ok := player.TeamIDs[team.ID]; !ok {

			teamButton := widget.NewButton(team.Name, func() {
				selectedTeams[team.ID] = team
				w.SetContent(selectedTeamPage(player, selectedTeams, db, w, a))
			})
			teamButtons = append(teamButtons, teamButton)
		}
	}
	content := container.NewVBox(
		returnToTeamSelectionPageButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

func createTeamButtons(player *mt.Player, team *mt.Team, db *mt.Database, selectedTeamButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {

	// Remove the selected team from the team list

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPage(player, db, w, a))
	})

	selectedTeamButtons = append(selectedTeamButtons, selectedTeamButton)
	return selectedTeamButtons

}

func addAnotherTeamPage(player *mt.Player, alreadySelectedTeams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToTeamSelectionPageButton := widget.NewButton("Return to team selection", func() {
		w.SetContent(selectedPlayerPage(player, db, w, a))
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
					w.SetContent(selectedTeamPage(player, alreadySelectedTeams, db, w, a))
				})
				teamButtons = append(teamButtons, teamButton)
			}
		}
	}

	content := container.NewVBox(
		returnToTeamSelectionPageButton,
		tLabel,
		container.NewVBox(teamButtons...),
	)

	return content
}

func selectedTeamPage(player *mt.Player, teams map[int]*mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
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

		successMsg := fmt.Sprintf("Player %v now plays in team %v", player.Name, strHelper(teamNames))
		fmt.Println(successMsg)
		dialog.ShowInformation("Succes", successMsg, w)

		// Return to empty page
		w.SetContent(
			currentSelectionPage(
				selectionPage(db, w, a), waitForPlayerSelectionPage(), db, w, a,
			),
		)
	})

	// Init
	selectedTeamButtons := []fyne.CanvasObject{}
	for _, team := range teams {
		selectedTeamButtons = createTeamButtons(player, team, db, selectedTeamButtons, w, a)
	}

	fmt.Println("Team buttons: ", selectedTeamButtons)

	addAnotherTeamButton := widget.NewButton("Add another team", func() {
		// Remove the selected team from the team list
		w.SetContent(addAnotherTeamPage(player, teams, db, w, a))
	})

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPage(db, w, a))
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

func AddPage(db *mt.Database, w fyne.Window, a fyne.App) {

	ReturnToFonctionalityPageButton := widget.NewButton("Return to the functionalities", func() {
		fonctionalityPage := FunctionalityPage(db, w, a)
		w.SetContent(fonctionalityPage)
	})

	// Link page

	// Remove a player from a team
	// Add a player to a club
	// Remove a player from a club

	addTtoPButton := widget.NewButton("Add team(s) to a player", func() {
		w.SetContent(
			currentSelectionPage(
				selectionPage(db, w, a),
				waitForPlayerSelectionPage(),
				db, w, a,
			),
		)
	})

	addPtoTButton := widget.NewButton("Add player(s) to a team", func() {

	})

	addPage := container.NewVBox(
		addTtoPButton,
		addPtoTButton,
		ReturnToFonctionalityPageButton,
	)

	w.SetContent(addPage)
}
