package myapp

import (
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
)

// Display launches the UI
func Display(appStartOption string) {

	mdb.AppStartOption(appStartOption)
	AppStartOption(appStartOption)
	StarterPage()
}
