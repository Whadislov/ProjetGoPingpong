package myapp

import (
	"fmt"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func CreateDeletePage(db *mt.Database, w fyne.Window, a fyne.App) {
	var rebuildUI func()

	// Rebuild UI on modifications
	rebuildUI = func() {

		pLabel := widget.NewLabel("Players")
		tLabel := widget.NewLabel("Teams")
		cLabel := widget.NewLabel("Clubs")

		returnToFonctionalityPageButton := widget.NewButton("Return to functionalities", func() {
			fonctionalityPage := FunctionalityPage(db, w, a)
			w.SetContent(fonctionalityPage)
		})

		// "Sort" maps
		sortedPlayers := SortMap(db.Players)
		sortedTeams := SortMap(db.Teams)
		sortedClubs := SortMap(db.Clubs)

		// Players
		acp := widget.NewAccordion()

		for _, sortedPlayer := range sortedPlayers {
			// i is PlayerID
			i := sortedPlayer.Key
			p := db.Players[i]
			item := widget.NewAccordionItem(
				fmt.Sprintf(p.Name),
				container.NewVBox(
					PlayerInfos(p),
					widget.NewButton("Delete", func() {
						ShowConfirmationDialog(w, fmt.Sprintf("Delete player %v?", p.Name), func() {
							err := mf.DeletePlayer(p, db)
							if err != nil {
								dialog.ShowError(err, w)
							} else {
								successMsg := fmt.Sprintf("%v has been successfully deleted\n", p.Name)
								fmt.Println(successMsg)
								dialog.ShowInformation("Succes", successMsg, w)
								// Reload UI
								rebuildUI()
							}
						})
					}),
				),
			)
			acp.Append(item)
		}

		// Teams
		act := widget.NewAccordion()

		for _, sortedTeam := range sortedTeams {
			// i is TeamID
			i := sortedTeam.Key
			t := db.Teams[i]
			item := widget.NewAccordionItem(
				fmt.Sprintf(t.Name),
				container.NewVBox(
					TeamInfos(t),
					widget.NewButton("Delete", func() {
						ShowConfirmationDialog(w, fmt.Sprintf("Delete team %v?", t.Name), func() {
							err := mf.DeleteTeam(t, db)
							if err != nil {
								dialog.ShowError(err, w)
							} else {
								successMsg := fmt.Sprintf("%v has been successfully deleted\n", t.Name)
								fmt.Println(successMsg)
								dialog.ShowInformation("Succes", successMsg, w)
								// Reload UI
								rebuildUI()
							}
						})
					}),
				),
			)
			act.Append(item)
		}

		// Clubs
		acc := widget.NewAccordion()

		for _, sortedClub := range sortedClubs {
			// i is ClubID
			i := sortedClub.Key
			c := db.Clubs[i]
			item := widget.NewAccordionItem(
				fmt.Sprintf(c.Name),
				container.NewVBox(
					ClubInfos(c),
					widget.NewButton("Delete", func() {
						ShowConfirmationDialog(w, fmt.Sprintf("Delete club %v?", c.Name), func() {
							err := mf.DeleteClub(c, db)
							if err != nil {
								dialog.ShowError(err, w)
							} else {
								successMsg := fmt.Sprintf("%v has been successfully deleted\n", c.Name)
								fmt.Println(successMsg)
								dialog.ShowInformation("Succes", successMsg, w)
								// Reload UI
								rebuildUI()
							}
						})
					}),
				),
			)
			acc.Append(item)
		}

		playerVBox := container.NewVBox(
			pLabel,
			acp)

		teamVBox := container.NewVBox(
			tLabel,
			act)

		clubVBox := container.NewVBox(
			cLabel,
			acc)

		w.SetContent(container.NewVBox(
			returnToFonctionalityPageButton,
			container.NewHBox(
				playerVBox,
				teamVBox,
				clubVBox),
		),
		)
	}

	// Build UI for the first time
	rebuildUI()

}
