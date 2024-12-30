package myapp

import (
	"slices"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TeamInfos returns a container that displays all the infos of a team.
func TeamInfos(team *mt.Team) *fyne.Container {
	wp := []fyne.CanvasObject{}
	// Sort players alphabetically
	players := []string{}
	for _, player := range team.PlayerIDs {
		players = append(players, player)
	}
	slices.Sort(players)

	for _, player := range players {
		wp = append(wp, widget.NewLabel(player))
	}

	itemp := widget.NewAccordionItem("Show players",
		container.NewVBox(wp...),
	)
	playerAc := widget.NewAccordion(itemp)
	clubs := []string{}
	for _, club := range team.ClubID {
		clubs = append(clubs, club)
	}
	clubStr := "Club: " + strHelper(clubs)

	item := container.NewVBox(
		widget.NewLabel(clubStr),
		playerAc,
	)

	return item
}

// TeamPage sets up the page for displaying team info.
func TeamPage(db *mt.Database, w fyne.Window, a fyne.App) {
	label := widget.NewLabel("Teams")
	ac := widget.NewAccordion()

	// "Sort the map"
	sortedTeams := SortMap(db.Teams)

	for _, team := range sortedTeams {
		item := widget.NewAccordionItem(team.Value.Name,
			TeamInfos(team.Value),
		)
		ac.Append(item)
	}

	returnToDatabasePageButton := widget.NewButton("Return to database", func() {
		databasePage := DatabasePage(db, w, a)
		w.SetContent(databasePage)
	})

	w.SetContent(container.NewVBox(
		returnToDatabasePageButton,
		label,
		ac),
	)
}
