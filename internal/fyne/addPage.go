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
	listOfPlayers := []fyne.CanvasObject{}

	for _, player := range db.Players {
		playerButton := widget.NewButton(player.Name, func() { w.SetContent(selectedPlayerPage(player, db, w, a)) })
		listOfPlayers = append(listOfPlayers, playerButton)
	}
	content := container.NewVBox(
		returnToPlayerSelectionPageButton,
		pLabel,
		container.NewVBox(listOfPlayers...),
	)

	return content
}

func selectedPlayerPage(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	tLabel := widget.NewLabel("Current selection of team(s) 🤝")

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
	listOfTeams := []fyne.CanvasObject{}

	for _, team := range db.Teams {
		// Check if the team from the database is already in the player's team map. If not we want a button of this team
		if _, ok := player.TeamIDs[team.ID]; !ok {
			teamButton := widget.NewButton(team.Name, func() { w.SetContent(selectedTeamPage(player, team, db, w, a)) })
			listOfTeams = append(listOfTeams, teamButton)
		}
	}
	content := container.NewVBox(
		returnToTeamSelectionPageButton,
		tLabel,
		container.NewVBox(listOfTeams...),
	)

	return content
}

func selectedTeamPage(player *mt.Player, team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToAddRemovePageButton := widget.NewButton("Return to add... remove...", func() {
		AddPage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	tLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🤝", team.Name))

	confirmButton := widget.NewButton("Confirm", func() {
		// Do the link
		err := mf.AddPlayerToTeam(player, team)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			successMsg := fmt.Sprintf("Player %v now plays in team %v", player.Name, team.Name)
			fmt.Println(successMsg)
			dialog.ShowInformation("Succes", successMsg, w)
		}

		// Return to empty page
		w.SetContent(
			currentSelectionPage(
				selectionPage(db, w, a), waitForPlayerSelectionPage(), db, w, a,
			),
		)
	})

	// User can click on the selected team to return the list of teams
	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPage(player, db, w, a))
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
		tLabel,
		selectedTeamButton,
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

/*

			teamButton := widget.NewButton(team.Name, func() {

				playerSelectionPageButton := widget.NewButton("Select player", func() { selectPlayerPage(db, w, a) })
				teamSelectionPageButton := widget.NewButton("Select team", func() { selectTeamForPlayerPage(player, db, w, a) })

				selectedTeamButton := widget.NewButton(team.Name, func() {
					w.SetContent(selectTeamForPlayerPage(player, db, w, a))
				})

				selectedPlayerButton := widget.NewButton(player.Name, func() {
					w.SetContent(selectPlayerPage(db, w, a))
				})

				confirmButton := widget.NewButton("Confirm", func() {
					err := mf.AddPlayerToTeam(player, team)
					if err != nil {
						dialog.ShowError(err, w)
					} else {
						successMsg := fmt.Sprintf("%v has been successfully added to %v\n", player.Name, team.Name)
						fmt.Println(successMsg)
						dialog.ShowInformation("Succes", successMsg, w)
					}
					content := container.NewVBox(
						container.NewGridWithColumns(
							2,
							container.NewVBox(
								pLabel,
								playerSelectionPageButton,
							),
							container.NewVBox(
								tLabel,
								teamSelectionPageButton,
							),
						),
						returnToAddRemovePageButton,
					)
					w.SetContent(content)

				})

				content := container.NewVBox(
					container.NewGridWithColumns(
						2,
						container.NewVBox(
							pLabel,
							playerSelectionPageButton,
							selectedPlayerButton,
						),
						container.NewVBox(
							tLabel,
							teamSelectionPageButton,
							selectedTeamButton,
						),
					),
					confirmButton,
					returnToAddRemovePageButton,
				)

				w.SetContent(content)

				//selectedTeamID = team.ID
			})

			if len(player.TeamIDs) == 0 {
				listOfTeams = append(listOfTeams, teamButton)
			} else {
				for teamID := range player.TeamIDs {
					if team.ID != teamID {
						listOfTeams = append(listOfTeams, teamButton)
					}
				}
			}

		}
		// Team buttons
		w.SetContent(
			container.NewVBox(
				tLabel,
				container.NewVBox(listOfTeams...),
			),
		)

	})

	//selectedTeam := db.Teams[selectedTeamID]
	return teamSelectionPageButton
}

*/

// Team function

/*

func selectedTeamPage(team *mt.Team, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToAddRemovePageButton := widget.NewButton("Return to add... remove...", func() {
		AddPage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🤝", team.Name))
	tLabel := widget.NewLabel("Add player 🏓")

	selectedTeamButton := widget.NewButton(team.Name, func() {
		w.SetContent(selectTeamPage(db, w, a))
	})

	playerSelectionPageButton := selectPlayer(team, db, w, a)
	teamSelectionPageButton := selectTeamPage(db, w, a)

	content := container.NewVBox(
		container.NewGridWithColumns(
			2,
			container.NewVBox(
				pLabel,
				playerSelectionPageButton,
			),
			container.NewVBox(
				tLabel,
				teamSelectionPageButton,
				selectedTeamButton,
			),
		),
		returnToAddRemovePageButton)

	w.SetContent(content)

	return content

}

func selectTeamPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	returnToTeamSelectionButton := widget.NewButton("Return to team selection", func() {
		w.SetContent(selectionPage(db, w, a))
	})

	tLabel := widget.NewLabel("Players 🏓")
	listOfTeams := []fyne.CanvasObject{}

	for _, team := range db.Teams {
		teamButton := widget.NewButton(team.Name, func() { selectedTeamPage(team, db, w, a) })
		listOfTeams = append(listOfTeams, teamButton)
	}
	content := container.NewVBox(
		returnToTeamSelectionButton,
		tLabel,
		container.NewVBox(listOfTeams...),
	)
	w.SetContent(content)

	return content
}

*/
