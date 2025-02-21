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

// currentSelectionPageCfromP sets up the page for selecting players and clubs.
func currentSelectionPageCfromP(playerContent *fyne.Container, clubContent *fyne.Container, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToRemovePageButton := widget.NewButton(T("return_to_remove_menu"), func() {
		RemovePage(db, w, a)
	})

	if clubContent == nil {
		content := container.NewVBox(
			playerContent,
			returnToRemovePageButton)
		return content
	} else {
		content := container.NewVBox(
			container.NewGridWithColumns(
				2,
				playerContent,
				clubContent,
			),
			returnToRemovePageButton)
		return content
	}
}

// SelectionPageCfromP sets up the initial selection page for players.
func SelectionPageCfromP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	playersInClub := 0
	for _, club := range db.Clubs {
		if len(club.PlayerIDs) > 0 {
			playersInClub++
		}
	}

	if playersInClub > 0 {
		pageTitle := setTitle(T("remove_select_a_player"), 32)
		playerSelectionPageButton := widget.NewButton(T("select_a_player"), func() { w.SetContent(selectPlayerPageCfromP(db, w, a)) })
		return container.NewVBox(
			pageTitle,
			playerSelectionPageButton)
	} else {
		pageTitle := setTitle(T("remove"), 32)
		return container.NewVBox(
			pageTitle,
			widget.NewLabel(T("currently_0_player_available")))
	}
}

// selectPlayerPageCfromP sets up the page for selecting a player from the database.
func selectPlayerPageCfromP(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle(T("remove_select_a_player"), 32)

	returnToPlayerSelectionPageCfromPButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(
			currentSelectionPageCfromP(
				SelectionPageCfromP(db, w, a), nil, db, w, a,
			),
		)
	})

	pLabel := widget.NewLabel(T("players_with_racket_emoji"))
	playerButtons := []fyne.CanvasObject{}

	// "Sort the map of players" for a better button display
	sortedPlayers := SortMap(db.Players)

	for _, p := range sortedPlayers {
		player := p.Value

		// If the player's club list is empty, we don't want a button for this player
		if len(player.ClubIDs) == 0 {
			continue
		} else {
			playerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() { w.SetContent(selectedPlayerPageCfromP(player, db, w, a)) })
			playerButtons = append(playerButtons, playerButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToPlayerSelectionPageCfromPButton,
		pLabel,
		container.NewVBox(playerButtons...),
	)

	return content
}

// selectedPlayerPageCfromP sets up the page for a selected player and allows club selection.
func selectedPlayerPageCfromP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageTitle := setTitle("Remove: select a club", 32)

	pLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+" %v 🏓", fmt.Sprintf("%v %v", player.Firstname, player.Lastname)))
	cLabel := widget.NewLabel(T("club_current_selection_house_emoji"))

	// User can click on the selected player to return to the list of player
	selectedPlayerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
		w.SetContent(selectPlayerPageCfromP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	if len(player.ClubIDs) == 0 {
		dialog.ShowInformation(T("information"), fmt.Sprintf("%v has no club", fmt.Sprintf("%v %v", player.Firstname, player.Lastname)), w)
		return selectPlayerPageCfromP(db, w, a)
	}

	// Now select a club
	selectClubButton := widget.NewButton(T("select_a_club"), func() {
		w.SetContent(selectClubPageCfromP(player, db, w, a))
	})

	clubContent := container.NewVBox(
		cLabel,
		selectClubButton,
	)

	// Now display the whole page, with availability to choose a club
	content := container.NewVBox(
		pageTitle,
		currentSelectionPageCfromP(
			playerContent,
			clubContent,
			db, w, a,
		),
	)

	return content
}

// selectClubPageCfromP sets up the page for selecting a club for a given player.
func selectClubPageCfromP(player *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Remove: select a club", 32)

	returnToClubSelectionPageCfromPButton := widget.NewButton(T("return_to_club_selection"), func() {
		w.SetContent(selectedPlayerPageCfromP(player, db, w, a))
	})

	cLabel := widget.NewLabel(T("clubs_house_emoji"))
	clubButtons := []fyne.CanvasObject{}
	selectedClub := make(map[uuid.UUID]*mt.Club)

	// Nothing to remove
	if len(db.Players) == 0 {
		okButton := widget.NewButton("Ok", func() {
			RemovePage(db, w, a)
		})

		label := widget.NewLabel("There is currently 0 club available.")
		content := container.NewVBox(
			pageTitle,
			label,
			okButton,
		)

		w.SetContent(content)
	}

	// "Sort the map of clubs" for a better button display
	sortedClubs := SortMap(db.Clubs)

	for _, c := range sortedClubs {
		club := c.Value
		// Check if the club from the database is already in the player's club map. If yes we want a button of this club
		if _, ok := player.ClubIDs[club.ID]; ok {
			clubButton := widget.NewButton(club.Name, func() {
				selectedClub[club.ID] = club
				w.SetContent(selectedClubPageCfromP(player, selectedClub, db, w, a))
			})
			clubButtons = append(clubButtons, clubButton)
		}
	}
	content := container.NewVBox(
		pageTitle,
		returnToClubSelectionPageCfromPButton,
		cLabel,
		container.NewVBox(clubButtons...),
	)

	return content
}

// createClubButtonsCfromP creates buttons for each selected club.
func createClubButtonsCfromP(player *mt.Player, club *mt.Club, db *mt.Database, selectedClub map[uuid.UUID]*mt.Club, selectedclubButtons []fyne.CanvasObject, w fyne.Window, a fyne.App) []fyne.CanvasObject {
	// User can click on the selected club to remove the club from the selected club list
	selectedclubButton := widget.NewButton(club.Name, func() {
		delete(selectedClub, club.ID)

		// If there is 0 selected club, we should return to the club selection page
		if len(selectedClub) == 0 {
			w.SetContent(selectClubPageCfromP(player, db, w, a))
		} else {
			w.SetContent(selectedClubPageCfromP(player, selectedClub, db, w, a))
		}
	})

	selectedclubButtons = append(selectedclubButtons, selectedclubButton)
	return selectedclubButtons

}

// addAnotherclubPageCfromP sets up the page for adding another club to the selected player.
func addAnotherclubPageCfromP(player *mt.Player, alreadyselectedClub map[uuid.UUID]*mt.Club, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	returnToClubSelectionPageCfromPButton := widget.NewButton(T("cancel"), func() {
		w.SetContent(selectedClubPageCfromP(player, alreadyselectedClub, db, w, a))
	})

	cLabel := widget.NewLabel(T("clubs_house_emoji"))
	clubButtons := []fyne.CanvasObject{}

	// "Sort the map of clubs" for a better button display
	sortedClubs := SortMap(db.Clubs)

	for _, c := range sortedClubs {
		club := c.Value
		// Check if the club from the database is already in the player's club map. If yes we want a button of this club
		if _, ok := player.ClubIDs[club.ID]; ok {
			if _, ok := alreadyselectedClub[club.ID]; !ok {
				// Check if the club from player's club map is already in selected clubs. If not we want a button of this club
				clubButton := widget.NewButton(club.Name, func() {
					alreadyselectedClub[club.ID] = club
					w.SetContent(selectedClubPageCfromP(player, alreadyselectedClub, db, w, a))
				})
				clubButtons = append(clubButtons, clubButton)
			}
		}
	}

	if len(clubButtons) == 0 {
		dialog.ShowInformation(T("information"), T("there_is_no_more_club_to_add"), w)
		w.SetContent(selectedClubPageCfromP(player, alreadyselectedClub, db, w, a))
	}

	content := container.NewVBox(
		returnToClubSelectionPageCfromPButton,
		cLabel,
		container.NewVBox(clubButtons...),
	)

	return content
}

// selectedClubPageCfromP sets up the page for confirming the selected clubs for a player.
func selectedClubPageCfromP(player *mt.Player, selectedClub map[uuid.UUID]*mt.Club, db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle(T("remove_confirm_selection"), 32)

	returnToRemovePageButton := widget.NewButton(T("return_to_remove_menu"), func() {
		RemovePage(db, w, a)
	})

	pLabel := widget.NewLabel(fmt.Sprintf(T("you_have_selected")+" %v 🏓", fmt.Sprintf("%v %v", player.Firstname, player.Lastname)))

	// "Sort the map of selectedClub" for a better button display
	sortedselectedClub := SortMap(selectedClub)

	confirmButton := widget.NewButton(T("confirm"), func() {
		var err error
		clubNames := []string{}
		for _, c := range sortedselectedClub {
			club := c.Value
			// Do the link
			err = mf.RemovePlayerFromClub(player, club)
			clubNames = append(clubNames, club.Name)
			if err != nil {
				dialog.ShowError(err, w)
			}
		}

		successMsg := fmt.Sprintf("%v is no longer playing in %v", fmt.Sprintf("%v %v", player.Firstname, player.Lastname), strHelper(clubNames))
		fmt.Println(successMsg)
		dialog.ShowInformation(T("success"), successMsg, w)

		// Set the flag to true to indicate that the database has changed
		HasChanged = true

		// Return to empty page
		w.SetContent(
			currentSelectionPageCfromP(
				SelectionPageCfromP(db, w, a), nil, db, w, a,
			),
		)
	})

	// Create buttons for each selected club
	selectedClubButtons := []fyne.CanvasObject{}
	for _, c := range sortedselectedClub {
		club := c.Value
		selectedClubButtons = createClubButtonsCfromP(player, club, db, selectedClub, selectedClubButtons, w, a)
	}

	// Add another club in the club selection
	addAnotherclubButton := widget.NewButton(T("add_another_club"), func() {
		w.SetContent(addAnotherclubPageCfromP(player, selectedClub, db, w, a))
	})

	// User can click on the selected player to return the list of players
	selectedPlayerButton := widget.NewButton(fmt.Sprintf("%v %v", player.Firstname, player.Lastname), func() {
		w.SetContent(selectPlayerPageCfromP(db, w, a))
	})

	playerContent := container.NewVBox(
		pLabel,
		selectedPlayerButton,
	)

	clubContent := container.NewVBox(
		addAnotherclubButton,
		container.NewVBox(selectedClubButtons...),
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
		returnToRemovePageButton,
	)

	return content
}
