package myapp

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// currentSelectionPageCtoP sets up the page for selecting players and clubs.
func currentSelectionPageCtoP(playerContent *fyne.Container, clubContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Add: select a club", 32)

	returnToAddPageButton := widget.NewButton("Return to the add menu", func() {
		AddPage(db, w, a)
	})

	if clubContent == nil {
		content := container.NewVBox(playerContent,
			returnToAddPageButton)
		return content
	} else {
		content := container.NewVBox(
			pageTitle,
			container.NewGridWithColumns(
				2,
				playerContent,
				clubContent,
			),
			returnToAddPageButton)
		return content
	}
}

// SelectionPageCtoP sets up the initial selection page for players.
func SelectionPageCtoP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Add: select a player", 32)

	playerSelectionPageCtoPButton := widget.NewButton("Select a player", func() { w.SetContent(selectPlayerPageCtoP(db, w, a)) })
	content := container.NewVBox(
		pageTitle,
		playerSelectionPageCtoPButton)

	return content
}

// selectPlayerPageCtoP sets up the page for selecting a player from the database.
func selectPlayerPageCtoP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Add: select a player", 32)

	returnToPlayerSelectionPageCtoPButton := widget.NewButton("Cancel", func() {
		w.SetContent(
			currentSelectionPageCtoP(
				SelectionPageCtoP(db, w, a), nil, db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel("Players 🏓")
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := SortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value
		playerButton := widget.NewButton(player.Name, func() { w.SetContent(selectedPlayerPageCtoP(player, db, w, a)) })
		playerButtons = append(playerButtons, playerButton)
	}
	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageCtoPButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPageCtoP sets up the page for a selected player and allows club selection.
func selectedPlayerPageCtoP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))
	cLabel := widget.NewLabel("Club current selection 🏠")

	// User can click on the selected player to return to the list of player
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPageCtoP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	// Now select a club
	selectClubButton := widget.NewButton("Select a club", func() {
		w.SetContent(selectClubPageCtoP(player, db, w, a))
	})

	clubContent := container.NewVBox(
		cLabel,
		selectClubButton,
	)

	// Now display the whole page, with availability to choose a club
	content := currentSelectionPageCtoP(
		playerContent,
		clubContent,
		db, w, a,
	)

	return content
}

// selectClubPageCtoP sets up the page for selecting a club for a given player.
func selectClubPageCtoP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Add: select a club", 32)

	returnToClubSelectionPageCtoPButton := widget.NewButton("Return to club selection", func() {
		w.SetContent(selectedPlayerPageCtoP(player, db, w, a))
	})

	cLabel := widget.NewLabel("Clubs 🏠")
	clubButtons := []fyne.CanvasObject{}
	selectedClub := make(map[int]*mt.Club)

	// We should create a club first
	if len(db.Clubs) == 0 {
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

		label := widget.NewLabel("There is currently 0 club available. Do you want to create a new club first ?")
		content := container.NewVBox(
			pageTitle,
			label,
			buttons,
		)

		w.SetContent(content)
	}

	// "Sort the map of clubs" for a better button display
	sortedClubs := SortMap(db.Clubs)

	for _, c := range sortedClubs {
		club := c.Value
		// Check if the club from the database is already in the player's club map. If not we want a button of this club
		if _, ok := player.ClubIDs[club.ID]; !ok {
			clubButton := widget.NewButton(club.Name, func() {
				selectedClub[club.ID] = club
				w.SetContent(selectedClubPageCtoP(player, selectedClub, db, w, a))
			})
			clubButtons = append(clubButtons, clubButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToClubSelectionPageCtoPButton,
		cLabel,
		container.NewVBox(clubButtons...),
	)

	return content
}

// createclubButtonsCtoP creates buttons for each selected club.
func createclubButtonsCtoP(player *mt.Player, club *mt.Club, db *mt.Database, selectedClub map[int]*mt.Club, selectedclubButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
	// User can click on the selected club to remove the club from the selected club list
	selectedclubButton := widget.NewButton(club.Name, func() {
		delete(selectedClub, club.ID)

		// If there is 0 selected club, we should return to the club selection page
		if len(selectedClub) == 0 {
			w.SetContent(selectClubPageCtoP(player, db, w, a))
		} else {
			w.SetContent(selectedClubPageCtoP(player, selectedClub, db, w, a))
		}
	})

	selectedclubButtons = append(selectedclubButtons, selectedclubButton)
	return selectedclubButtons

}

// addAnotherclubPageCtoP sets up the page for adding another club to the selected player.
func addAnotherclubPageCtoP(player *mt.Player, alreadyselectedClub map[int]*mt.Club, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Add: select a club", 32)

	returnToClubSelectionPageCtoPButton := widget.NewButton("Cancel", func() {
		w.SetContent(selectedClubPageCtoP(player, alreadyselectedClub, db, w, a))
	})

	cLabel := widget.NewLabel("Clubs 🏠")
	clubButtons := []fyne.CanvasObject{}

	// "Sort the map of clubs" for a better button display
	sortedClubs := SortMap(db.Clubs)

	for _, c := range sortedClubs {
		club := c.Value
		// Check if the club from the database is already in the player's club map. If not we want a button of this club
		if _, ok := player.ClubIDs[club.ID]; !ok {
			if _, ok := alreadyselectedClub[club.ID]; !ok {
				// Check if the club from player's club map is already in selected clubs. If not we want a button of this club
				clubButton := widget.NewButton(club.Name, func() {
					alreadyselectedClub[club.ID] = club
					w.SetContent(selectedClubPageCtoP(player, alreadyselectedClub, db, w, a))
				})
				clubButtons = append(clubButtons, clubButton)
			}
		}
	}

	if len(clubButtons) == 0 {
		dialog.ShowInformation("Information", "There is no more club to add", w)
		w.SetContent(selectedClubPageCtoP(player, alreadyselectedClub, db, w, a))
	}

	content := container.NewVBox(
		pageTitle,
		returnToClubSelectionPageCtoPButton,
		cLabel,
		container.NewVBox(clubButtons...),
	)

	return content
}

// selectedClubPageCtoP sets up the page for confirming the selected clubs for a player.
func selectedClubPageCtoP(player *mt.Player, selectedClub map[int]*mt.Club, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Add: confirm", 32)

	returnToAddRemovePageButton := widget.NewButton("Return to the add menu", func() {
		AddPage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf("You have selected %v 🏓", player.Name))

	// "Sort the map of selectedClub" for a better button display
	sortedselectedClub := SortMap(selectedClub)

	confirmButton := widget.NewButton("Confirm", func() {
		var err error
		clubNames := []string{}
		for _, c := range sortedselectedClub {
			club := c.Value
			// Do the link
			err = mf.AddPlayerToClub(player, club)
			clubNames = append(clubNames, club.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf("Player %v now plays in club(s) %v", player.Name, strHelper(clubNames))
		fmt.Println(successMsg)
		dialog.ShowInformation("Succes", successMsg, w)

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPageCtoP(
				SelectionPageCtoP(db, w, a), nil, db, w, a,
			),
		)
	})

	// Create buttons for each selected club
	selectedclubButtons := []fyne.CanvasObject{}
	for _, c := range sortedselectedClub {
		club := c.Value
		selectedclubButtons = createclubButtonsCtoP(player, club, db, selectedClub, selectedclubButtons, w, a)
	}

	// Add another club in the club selection
	addAnotherclubButton := widget.NewButton("Add another club", func() {
		w.SetContent(addAnotherclubPageCtoP(player, selectedClub, db, w, a))
	})

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(player.Name, func() {
		w.SetContent(selectPlayerPageCtoP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	clubContent := container.NewVBox(
		addAnotherclubButton,
		container.NewVBox(selectedclubButtons...),
	)

	// Now display the whole finished page, with chosen clubs
	content := container.NewVBox(
		pageTitle,
		container.NewGridWithColumns(
			2,
			playerContent,
			clubContent,
		),
		confirmButton,
		returnToAddRemovePageButton,
	)

	return content
}
