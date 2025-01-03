package myapp

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// CreatePage sets up the page for creating players, teams, and clubs.
func CreatePage(db *mt.Database, w fyne.Window, a fyne.App) {

	ReturnToFonctionalityPageButton := widget.NewButton("Return to the functionalities", func() {
		fonctionalityPage := FunctionalityPage(db, w, a)
		w.SetContent(fonctionalityPage)
	})

	ReturnToCreatePageButton := widget.NewButton("Return to the creation menu", func() {
		CreatePage(db, w, a)
	})

	// Player
	playerButton := widget.NewButton("Create a new player", func() {

		// Club Selection
		selectClubLabel := widget.NewLabel("You must first select a club")

		// clubSelectionPage
		clubSelectionPageButton := widget.NewButton("Select a club", func() {
			label := widget.NewLabel("Clubs")
			listOfClubs := []fyne.CanvasObject{}

			// Sort clubs for an alphabetical order button display
			sortedClubs := SortMap(db.Clubs)

			for _, c := range sortedClubs {
				club := c.Value
				clubButton := widget.NewButton(club.Name, func() {
					// After club selection
					clubLabel := widget.NewLabel(fmt.Sprintf("You are going to create a player for %v\n", club.Name))

					// We can now create the player
					nameEntry := widget.NewEntry()
					entryNameHolder := "Name ..."
					nameEntry.SetPlaceHolder(entryNameHolder)

					// Here are optional informations that can be added to the player
					ageEntry := widget.NewEntry()
					entryAgeHolder := "Age ..."
					ageEntry.SetPlaceHolder(entryAgeHolder)

					rankingEntry := widget.NewEntry()
					entryRankingHolder := "Ranking ..."
					rankingEntry.SetPlaceHolder(entryRankingHolder)

					forehandEntry := widget.NewEntry()
					entryForehandHolder := "Forehand ..."
					forehandEntry.SetPlaceHolder(entryForehandHolder)

					backhandEntry := widget.NewEntry()
					entryBackhandHolder := "Backhand ..."
					backhandEntry.SetPlaceHolder(entryBackhandHolder)

					bladeEntry := widget.NewEntry()
					entryBladeHolder := "Blade ..."
					bladeEntry.SetPlaceHolder(entryBladeHolder)

					validatationButton := widget.NewButton("Create", func() {
						name := nameEntry.Text
						p, errName := mf.NewPlayer(name, db)
						// Set player age
						age, errAge := strconv.Atoi(ageEntry.Text)
						p.SetPlayerAge(age)
						// Set player ranking
						ranking, errRanking := strconv.Atoi(rankingEntry.Text)
						p.SetPlayerRanking(ranking)
						// Set player material
						if forehandEntry.Text == "" {
							forehandEntry.SetText("Unknown")
						}
						if backhandEntry.Text == "" {
							backhandEntry.SetText("Unknown")
						}
						if bladeEntry.Text == "" {
							bladeEntry.SetText("Unknown")
						}
						p.SetPlayerMaterial(forehandEntry.Text, backhandEntry.Text, bladeEntry.Text)

						if errName != nil {
							dialog.ShowError(errName, w)
							return
						} else if errAge != nil {
							dialog.ShowError(errAge, w)
							return
						} else if errRanking != nil {
							dialog.ShowError(errRanking, w)
							return
						} else {
							// Link the player to the club
							err := mf.AddPlayerToClub(p, club)
							if err != nil {
								dialog.ShowError(err, w)
								return
							} else {
								// Player creation + link to club success
								successMsg := fmt.Sprintf("Player %v has been successfully created\n", name)
								fmt.Println(successMsg)
								dialog.ShowInformation("Succes", successMsg, w)

								// Set the flag to true to indicate that the database has changed
								HasChanged = true

								// Reinit the entry text
								ReinitWidgetEntryText(nameEntry, entryNameHolder)
								ReinitWidgetEntryText(ageEntry, entryAgeHolder)
								ReinitWidgetEntryText(rankingEntry, entryRankingHolder)
								ReinitWidgetEntryText(forehandEntry, entryForehandHolder)
								ReinitWidgetEntryText(backhandEntry, entryBackhandHolder)
								ReinitWidgetEntryText(bladeEntry, entryBladeHolder)
							}
						}
					})
					// Create a player in this club page
					w.SetContent(container.NewVBox(
						clubLabel,
						nameEntry,
						ageEntry,
						rankingEntry,
						container.NewGridWithColumns(3, forehandEntry, backhandEntry, bladeEntry),
						validatationButton,
						ReturnToCreatePageButton,
					))
				})
				listOfClubs = append(listOfClubs, clubButton)
			}
			// Choose a club page
			w.SetContent(container.NewVBox(
				label,
				container.NewVBox(listOfClubs...),
			))
		})
		// Club selection page
		w.SetContent(container.NewVBox(
			selectClubLabel,
			clubSelectionPageButton,
			ReturnToCreatePageButton,
		))

	})

	// Team
	teamButton := widget.NewButton("Create a new team", func() {

		// Club Selection
		selectClubLabel := widget.NewLabel("You must first select a club")

		// clubSelectionPage
		clubSelectionPageButton := widget.NewButton("Select a club", func() {
			label := widget.NewLabel("Clubs")
			listOfClubs := []fyne.CanvasObject{}

			for _, club := range db.Clubs {
				clubButton := widget.NewButton(club.Name, func() {
					// After club selection
					clubLabel := widget.NewLabel(fmt.Sprintf("You are going to create a team for %v\n", club.Name))

					// We can now create the team
					nameEntry := widget.NewEntry()
					entryHolder := "Enter your team name here..."
					nameEntry.SetPlaceHolder(entryHolder)

					validatationButton := widget.NewButton("Create", func() {
						name := nameEntry.Text

						// If team name already exists, do not create the team
						for _, value := range db.Teams {
							if value.Name == name {
								err := fmt.Errorf(fmt.Sprintf("Team %v already exists in %v", name, club.Name))
								dialog.ShowError(err, w)
								// Reinit the text
								nameEntry.SetText("")
								nameEntry.SetPlaceHolder(entryHolder)
								return
							}
						}

						t, err := mf.NewTeam(name, db)

						if err != nil {
							dialog.ShowError(err, w)
							return
						} else {
							// Link the team to the club
							err := mf.AddTeamToClub(t, club)
							if err != nil {
								dialog.ShowError(err, w)
								return
							} else {
								// team creation + link to club success
								successMsg := fmt.Sprintf("Team %v has been successfully created\n", name)
								fmt.Println(successMsg)
								dialog.ShowInformation("Succes", successMsg, w)

								// Set the flag to true to indicate that the database has changed
								HasChanged = true

								// Reinit the text
								nameEntry.SetText("")
								nameEntry.SetPlaceHolder(entryHolder)
							}
						}
					})
					// Create a team in this club page
					w.SetContent(container.NewVBox(
						clubLabel,
						nameEntry,
						validatationButton,
						ReturnToCreatePageButton,
					))
				})
				listOfClubs = append(listOfClubs, clubButton)
			}
			// Choose a club page
			w.SetContent(container.NewVBox(
				label,
				container.NewVBox(listOfClubs...),
			))
		})
		// Club selection page
		w.SetContent(container.NewVBox(
			selectClubLabel,
			clubSelectionPageButton,
			ReturnToCreatePageButton,
		))

	})
	// Club
	clubButton := widget.NewButton("Create a new club", func() {
		nameEntry := widget.NewEntry()
		entryHolder := "Enter your club name here..."
		nameEntry.SetPlaceHolder(entryHolder)

		validatationButton := widget.NewButton("Create", func() {
			name := nameEntry.Text
			_, err := mf.NewClub(name, db)

			if err != nil {
				dialog.ShowError(err, w)
				return
			} else {
				successMsg := fmt.Sprintf("Club %v has been successfully created\n", name)
				fmt.Println(successMsg)
				dialog.ShowInformation("Succes", successMsg, w)

				// Set the flag to true to indicate that the database has changed
				HasChanged = true

				// Reinit the text
				nameEntry.SetText("")
				nameEntry.SetPlaceHolder(entryHolder)
			}
		})
		// Create a club page
		w.SetContent(container.NewVBox(
			nameEntry,
			validatationButton,
			ReturnToCreatePageButton,
		))
	})

	// Page definition

	// If there is no club, a club must first be created

	if len(db.Clubs) < 1 {
		label := widget.NewLabel("In order to create new players and teams, you need first to create a new club")

		createPage := container.NewVBox(
			label,
			clubButton,
			ReturnToFonctionalityPageButton,
		)
		// Create menu page
		w.SetContent(createPage)
	} else {
		createPage := container.NewVBox(
			playerButton,
			teamButton,
			clubButton,
			ReturnToFonctionalityPageButton,
		)
		// Create menu page
		w.SetContent(createPage)
	}
}
