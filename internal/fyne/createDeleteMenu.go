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

func CreateDeleteWindow(a fyne.App, label string, onConfirm func()) fyne.Window {
	ynw := a.NewWindow("Confirmation")
	ynw.SetContent(container.NewVBox(
		widget.NewLabel(label),
		container.NewHBox(
			widget.NewButton("Yes", func() {
				onConfirm()
				ynw.Close()
			}),
			widget.NewButton("No", func() {
				ynw.Close()
			}),
		),
	))
	ynw.Resize(fyne.NewSize(300, 150))
	return ynw
}

func ShowConfirmationDialog(w fyne.Window, message string, onConfirm func()) {
	d := dialog.NewConfirm("Confirm deletion", message, func(confirm bool) {
		if confirm {
			onConfirm()
		}
	}, w)
	d.Show()
}

func CreateDeleteMenu(w fyne.Window, db *mt.Database) {
	var rebuildUI func()

	// Rebuild UI on modifications
	rebuildUI = func() {

		// Players
		acp := widget.NewAccordion()

		for _, player := range db.Players {
			p := player
			item := widget.NewAccordionItem(
				fmt.Sprintf("Delete player %v", p.Name),
				container.NewVBox(
					PlayerInfos(p),
					widget.NewButton("Delete", func() {
						ShowConfirmationDialog(w, fmt.Sprintf("Delete player %v?", p.Name), func() {
							err := mf.DeletePlayer(p, db)
							if err != nil {
								dialog.ShowError(err, w)
							} else {
								fmt.Printf("Player %v has been successfully deleted\n", p.Name)
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

		for _, team := range db.Teams {
			t := team
			item := widget.NewAccordionItem(
				fmt.Sprintf("Delete team %v", t.Name),
				container.NewVBox(
					TeamInfos(t),
					widget.NewButton("Delete", func() {
						ShowConfirmationDialog(w, fmt.Sprintf("Delete team %v?", t.Name), func() {
							err := mf.DeleteTeam(t, db)
							if err != nil {
								fmt.Printf("UI Error: %v\n", err)
								dialog.ShowError(err, w)
							} else {
								successMsg := fmt.Sprintf("Team %v has been successfully deleted\n", t.Name)
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

		for _, club := range db.Clubs {
			c := club
			item := widget.NewAccordionItem(
				fmt.Sprintf("Delete team %v", c.Name),
				container.NewVBox(
					ClubInfos(c),
					widget.NewButton("Delete", func() {
						ShowConfirmationDialog(w, fmt.Sprintf("Delete club %v?", c.Name), func() {
							err := mf.DeleteClub(c, db)
							if err != nil {
								dialog.ShowError(err, w)
							} else {
								fmt.Printf("Club %v has been successfully deleted\n", c.Name)
								// Reload UI
								rebuildUI()
							}
						})
					}),
				),
			)
			act.Append(item)
		}

		w.SetContent(container.NewHBox(
			acp,
			act,
			acc),
		)
	}

	// Build UI for the first time
	rebuildUI()

}
