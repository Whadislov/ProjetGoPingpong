package myapp

import (
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
)

// Display launches the UI
func Display(appSOption string) {

	appStartOption = appSOption
	mdb.AppStartOption(appSOption)
	StarterPage()
}
