package myapp

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	mf "github.com/Whadislov/TTCompanion/internal/my_functions"
	mt "github.com/Whadislov/TTCompanion/internal/my_types"
)

// CreatePage sets up the page for creating players, teams, and clubs.
func CreatePage(db *mt.Database, w fyne.Window, a fyne.App) {

	pageTitle := setTitle(T("create"), 32)

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
		pageTitle := setTitle("Create new player: select a club", 32)

		// clubSelectionPage
		clubSelectionPageButton := widget.NewButton("Select a club", func() {
			pageTitle := setTitle("Create new player: select a club", 32)
			listOfClubs := []fyne.CanvasObject{}

			// Sort clubs for an alphabetical order button display
			sortedClubs := SortMap(db.Clubs)

			for _, c := range sortedClubs {
				club := c.Value
				clubButton := widget.NewButton(club.Name, func() {
					// After club selection
					clubLabel := widget.NewLabel(fmt.Sprintf("You are going to create a player for %v\n", club.Name))

					// We can now create the player
					firstnameEntry := widget.NewEntry()
					entryFirstnameHolder := "Firstname ..."
					firstnameEntry.SetPlaceHolder(entryFirstnameHolder)

					lastnameEntry := widget.NewEntry()
					entryLastnameHolder := "Lastname ..."
					lastnameEntry.SetPlaceHolder(entryLastnameHolder)

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

					validatationButton := widget.NewButton(T("create"), func() {
						age := -1
						ranking := -1

						// Check player name
						if firstnameEntry.Text == "" {
							dialog.ShowError(fmt.Errorf("firstname must not be empty"), w)
							return
						} else if !IsLettersOnly(firstnameEntry.Text) {
							dialog.ShowError(fmt.Errorf("firstname must be letters only"), w)
							return
						}
						if lastnameEntry.Text == "" {
							dialog.ShowError(fmt.Errorf("lastname must not be empty"), w)
							return
						} else if !IsLettersOnly(lastnameEntry.Text) {
							dialog.ShowError(fmt.Errorf("lastname must be letters only"), w)
							return
						}
						// Set player age
						if ageEntry.Text != "" {
							a, errAge := strconv.Atoi(ageEntry.Text)
							if errAge != nil {
								// Check if the age is a number
								if !IsNumbersOnly(ageEntry.Text) {
									dialog.ShowError(fmt.Errorf("age must be a number"), w)
									return
								} else {
									dialog.ShowError(errAge, w)
									return
								}
							} else {
								age = a
							}
						}

						// Set player ranking
						if rankingEntry.Text != "" {
							r, errRanking := strconv.Atoi(rankingEntry.Text)
							if errRanking != nil {
								// Check if the ranking is a number
								if !IsNumbersOnly(rankingEntry.Text) {
									dialog.ShowError(fmt.Errorf("ranking must be a number"), w)
									return
								} else {
									dialog.ShowError(errRanking, w)
									return
								}
							} else {
								ranking = r
							}
						}

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

						// Create the player
						firstname := firstnameEntry.Text
						lastname := lastnameEntry.Text

						p, errName := mf.NewPlayer(firstname, lastname, db)
						if errName != nil {
							dialog.ShowError(errName, w)
							return
						} else {
							p.SetPlayerAge(age)
							p.SetPlayerRanking(ranking)
							p.SetPlayerMaterial(forehandEntry.Text, backhandEntry.Text, bladeEntry.Text)
						}

						// Link the player to the club
						err := mf.AddPlayerToClub(p, club)
						if err != nil {
							dialog.ShowError(err, w)
							return
						} else {
							// Player creation + link to club success
							successMsg := fmt.Sprintf("Player %v %v has been successfully created\n", firstname, lastname)
							fmt.Println(successMsg)
							dialog.ShowInformation("Succes", successMsg, w)

							// Set the flag to true to indicate that the database has changed
							HasChanged = true

							// Reinit the entry text
							ReinitWidgetEntryText(firstnameEntry, entryFirstnameHolder)
							ReinitWidgetEntryText(lastnameEntry, entryLastnameHolder)
							ReinitWidgetEntryText(ageEntry, entryAgeHolder)
							ReinitWidgetEntryText(rankingEntry, entryRankingHolder)
							ReinitWidgetEntryText(forehandEntry, entryForehandHolder)
							ReinitWidgetEntryText(backhandEntry, entryBackhandHolder)
							ReinitWidgetEntryText(bladeEntry, entryBladeHolder)
						}

					})
					// Create a player in this club page
					pageTitle := setTitle("Create new player: add player information", 32)

					w.SetContent(container.NewVBox(
						pageTitle,
						clubLabel,
						firstnameEntry,
						lastnameEntry,
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
				pageTitle,
				container.NewVBox(listOfClubs...),
			))
		})
		// Club selection page
		w.SetContent(container.NewVBox(
			pageTitle,
			clubSelectionPageButton,
			ReturnToCreatePageButton,
		))

	})

	// Team
	teamButton := widget.NewButton("Create a new team", func() {

		// Club Selection
		pageTitle := setTitle("Create new team: select a club", 32)

		// clubSelectionPage
		clubSelectionPageButton := widget.NewButton("Select a club", func() {
			pageTitle := setTitle("Create new team: select a club", 32)
			listOfClubs := []fyne.CanvasObject{}

			for _, club := range db.Clubs {
				clubButton := widget.NewButton(club.Name, func() {
					// After club selection
					clubLabel := widget.NewLabel(fmt.Sprintf("You are going to create a team for %v\n", club.Name))

					// We can now create the team
					nameEntry := widget.NewEntry()
					entryHolder := "Enter your team name here..."
					nameEntry.SetPlaceHolder(entryHolder)

					validatationButton := widget.NewButton(T("create"), func() {
						name := nameEntry.Text

						// If team name already exists, do not create the team
						for _, value := range db.Teams {
							if value.Name == name {
								err := fmt.Errorf("team %v already exists in %v", name, club.Name)
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
					pageTitle := setTitle("Create new team", 32)
					w.SetContent(container.NewVBox(
						pageTitle,
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
				pageTitle,
				container.NewVBox(listOfClubs...),
			))
		})
		// Club selection page
		w.SetContent(container.NewVBox(
			pageTitle,
			clubSelectionPageButton,
			ReturnToCreatePageButton,
		))

	})
	// Club
	clubButton := widget.NewButton("Create a new club", func() {
		nameEntry := widget.NewEntry()
		entryHolder := "Enter your club name here..."
		nameEntry.SetPlaceHolder(entryHolder)

		validatationButton := widget.NewButton(T("create"), func() {
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
		pageTitle := setTitle("Create new club", 32)
		w.SetContent(container.NewVBox(
			pageTitle,
			nameEntry,
			validatationButton,
			ReturnToCreatePageButton,
		))
	})

	// Page definition

	// If there is no club, a club must first be created

	if len(db.Clubs) < 1 {
		label := widget.NewLabel("You currently have 0 club available, please create a club first.")

		createPage := container.NewVBox(
			pageTitle,
			label,
			clubButton,
			ReturnToFonctionalityPageButton,
		)
		// Create menu page
		w.SetContent(createPage)
	} else {
		createPage := container.NewVBox(
			pageTitle,
			playerButton,
			teamButton,
			clubButton,
			ReturnToFonctionalityPageButton,
		)
		// Create menu page
		w.SetContent(createPage)
	}
}
