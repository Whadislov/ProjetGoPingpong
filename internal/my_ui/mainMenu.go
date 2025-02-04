package myapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"log"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mfr "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

// MainMenu creates the menus
func MainMenu(db *mt.Database, w fyne.Window, a fyne.App) *fyne.MainMenu {

	mainPage := MainPage(db, w, a)

	menu1Item1 := fyne.NewMenuItem("Main page", func() { w.SetContent(mainPage) })
	menu1Item2 := fyne.NewMenuItem("My profile", func() { UserPage(userOfSession, db, w, a) })
	menu1Item3 := fyne.NewMenuItem("Save changes", func() {
		if !HasChanged {
			dialog.ShowInformation("Information", "There is nothing new to save", w)
		} else {
			if appStartOption == "local" {
				err := mdb.SaveDB(db)
				if err != nil {
					dialog.ShowError(err, w)
				} else {
					HasChanged = false
					// Reload the database after saving (refresh the IDs)
					db, err = mdb.LoadDB()
					if err != nil {
						dialog.ShowError(err, w)
					} else {
						dialog.ShowInformation("Information", "Changes saved", w)
					}
				}

			} else if appStartOption == "browser" {
				err := mfr.SaveDB(db)
				if err != nil {
					dialog.ShowError(err, w)
				} else {
					HasChanged = false
					// Reload the database after saving (refresh the IDs)
					db, err = mfr.LoadDB(jsonWebToken)
					if err != nil {
						dialog.ShowError(err, w)
					} else {
						dialog.ShowInformation("Information", "Changes saved", w)
					}
				}
			}
		}
	})
	menu1Item4 := fyne.NewMenuItem("Log out", func() {
		if HasChanged {
			dialog.ShowConfirm("Unsaved Changes", "You have unsaved changes. Do you want to save them before logging out?", func(confirm bool) {
				if confirm {
					// User wants to save the changes
					var err error
					if appStartOption == "local" {
						err = mdb.SaveDB(db)
					} else if appStartOption == "browser" {
						err = mfr.SaveDB(db)
					}
					if err != nil {
						dialog.ShowError(err, w)
					} else {
						HasChanged = false
						log.Println("User logged out and saved his changes.")
						w.SetContent(AuthentificationPage(w, a))
					}
				} else {
					// User does not want to save the changes
					log.Println("User logged out and saved nothing.")
					w.SetContent(AuthentificationPage(w, a))
				}
			}, w)
		} else {
			// Nothing has changed
			log.Println("User logged out.")
			w.SetContent(AuthentificationPage(w, a))
		}

	})

	newMenu1 := fyne.NewMenu("Main menu", menu1Item1, menu1Item2, menu1Item3, menu1Item4)

	menu2Item1 := fyne.NewMenuItem("Players", func() { PlayerPage(db, w, a) })
	menu2Item2 := fyne.NewMenuItem("Teams", func() { TeamPage(db, w, a) })
	menu2Item3 := fyne.NewMenuItem("Clubs", func() { ClubPage(db, w, a) })
	newMenu2 := fyne.NewMenu("Database", menu2Item1, menu2Item2, menu2Item3)

	menu3Item1 := fyne.NewMenuItem("Create ", func() { CreatePage(db, w, a) })
	menu3Item2 := fyne.NewMenuItem("Add ... to ...", func() { AddPage(db, w, a) })
	menu3Item3 := fyne.NewMenuItem("Remove ... from ...", func() { RemovePage(db, w, a) })
	menu3Item4 := fyne.NewMenuItem("Delete", func() { DeletePage(db, w, a) })
	menu3Item5 := fyne.NewMenuItem("Modify player information", func() { AddInfoToPlayerPage(db, w, a) })
	newMenu3 := fyne.NewMenu("Functions", menu3Item1, menu3Item2, menu3Item3, menu3Item4, menu3Item5)

	menu := fyne.NewMainMenu(newMenu1, newMenu2, newMenu3)

	return menu

}
