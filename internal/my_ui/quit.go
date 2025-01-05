package myapp

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"

	msql "github.com/Whadislov/ProjetGoPingPong/internal/my_sqlitedb"
	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"
)

func Quit(db *mt.Database, w fyne.Window, a fyne.App, HasChanged bool) {

	// Check for unsaved changes before quitting

	if HasChanged {
		dialog.ShowConfirm("Unsaved Changes", "You have unsaved changes. Do you want to save them before quitting?", func(confirm bool) {
			if confirm {
				err := msql.SaveDB(db)
				if err != nil {
					dialog.ShowError(err, w)
				}
				a.Quit()
				log.Println("Application exited.")

			} else {
				a.Quit()
				log.Println("Application exited.")

			}
		}, w)
	} else {
		a.Quit()
		log.Println("Application exited.")
	}
}
