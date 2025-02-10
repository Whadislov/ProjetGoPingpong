package myapp

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func AddInfoToPlayerPage(db *mt.Database, w fyne.Window, a fyne.App) {
	pageTitle := setTitle("Edit player information", 32)

	returnToFonctionalityPageButton := widget.NewButton("Return to functionalities", func() { w.SetContent(FunctionalityPage(db, w, a)) })

	// When the player is not yet selected, display nothing
	rightContent := container.NewVBox()

	sortedPlayers := SortMap(db.Players)
	playerButtons := []fyne.CanvasObject{}

	for _, p := range sortedPlayers {
		player := p.Value
		playerButton := widget.NewButton(player.Firstname+player.Lastname, func() { AddInfoToSelectedPlayerPage(player, db, w, a) })
		playerButtons = append(playerButtons, playerButton)
	}

	content := container.NewVBox(
		pageTitle,
		returnToFonctionalityPageButton,
		container.NewGridWithColumns(
			2,
			container.NewVBox(playerButtons...),
			rightContent,
		),
	)

	w.SetContent(content)
}

func AddInfoToSelectedPlayerPage(p *mt.Player, db *mt.Database, w fyne.Window, a fyne.App) {
	pageTitle := setTitle("Edit player information", 32)

	returnToFonctionalityPageButton := widget.NewButton("Return to functionalities", func() { w.SetContent(FunctionalityPage(db, w, a)) })

	cancelButton := widget.NewButton("Cancel", func() { AddInfoToPlayerPage(db, w, a) })

	playerLabel := widget.NewLabel(fmt.Sprintf("You have selected %v %v.", p.Firstname, p.Lastname))

	// Here are optional informations that can be added to the player
	ageEntry := widget.NewEntry()
	entryAgeHolder := strconv.Itoa(p.Age)
	ageEntry.SetPlaceHolder(entryAgeHolder)

	rankingEntry := widget.NewEntry()
	entryRankingHolder := strconv.Itoa(p.Ranking)
	rankingEntry.SetPlaceHolder(entryRankingHolder)

	forehandEntry := widget.NewEntry()
	entryForehandHolder := p.Material[0]
	forehandEntry.SetPlaceHolder(entryForehandHolder)

	backhandEntry := widget.NewEntry()
	entryBackhandHolder := p.Material[1]
	backhandEntry.SetPlaceHolder(entryBackhandHolder)

	bladeEntry := widget.NewEntry()
	entryBladeHolder := p.Material[2]
	bladeEntry.SetPlaceHolder(entryBladeHolder)

	var isAgeModified bool
	var isRankingModified bool
	var isForehandModified bool
	var isBackhandModified bool
	var isBladeModified bool

	confirmButton := widget.NewButton("Confirm", func() {
		// Check player age
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
				if a != -1 {
					isAgeModified = true
					p.SetPlayerAge(a)
				}
			}
		}

		// Check player ranking
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
				if r != -1 {
					isRankingModified = true
					p.SetPlayerRanking(r)
				}
			}
		}

		// Check player material
		if forehandEntry.Text != "" {
			isForehandModified = true
			p.SetPlayerMaterial(forehandEntry.Text, p.Material[1], p.Material[2])
		}
		if backhandEntry.Text != "" {
			isBackhandModified = true
			p.SetPlayerMaterial(p.Material[0], backhandEntry.Text, p.Material[2])
		}
		if bladeEntry.Text != "" {
			isBladeModified = true
			p.SetPlayerMaterial(p.Material[0], p.Material[1], bladeEntry.Text)
		}

		// Has something changed ?
		if isAgeModified || isRankingModified || isForehandModified || isBackhandModified || isBladeModified {
			HasChanged = true
			dialog.ShowInformation("Succes", fmt.Sprintf("%v %v has been modified", p.Firstname, p.Lastname), w)
			AddInfoToPlayerPage(db, w, a)
		} else {
			dialog.ShowInformation("Information", fmt.Sprintf("%v %v has not been modified", p.Firstname, p.Lastname), w)
		}

	})

	sortedPlayers := SortMap(db.Players)
	playerButtons := []fyne.CanvasObject{}

	for _, p := range sortedPlayers {
		player := p.Value
		playerButton := widget.NewButton(player.Firstname+player.Lastname, func() { AddInfoToSelectedPlayerPage(player, db, w, a) })
		playerButtons = append(playerButtons, playerButton)
	}

	leftContent := container.NewVBox(playerButtons...)

	formLayout := container.New(layout.NewFormLayout(),
		widget.NewLabel("Age:"), ageEntry,
		widget.NewLabel("Ranking:"), rankingEntry,
		widget.NewLabel("Forehand:"), forehandEntry,
		widget.NewLabel("Backhand:"), backhandEntry,
		widget.NewLabel("Blade:"), bladeEntry,
	)

	rightContent := container.NewVBox(
		playerLabel,
		formLayout,
		container.NewGridWithColumns(2, confirmButton, cancelButton),
	)

	content := container.NewVBox(
		pageTitle,
		returnToFonctionalityPageButton,
		container.NewGridWithColumns(
			2,
			leftContent,
			rightContent,
		),
	)
	w.SetContent(content)
}
