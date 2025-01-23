package myapp

import (
	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	//mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// IdentificationPage returns a page that contains a log in button and a sign up button
func IdentificationPage(w fyne.Window, a fyne.App) *fyne.Container {

	pageLabel := widget.NewLabel("Please identify yourself")

	logInButton := widget.NewButton("Log in", func() {
		w.SetContent(logInPage(w, a))
	})

	signUpButton := widget.NewButton("Sign up", func() {
		w.SetContent(signUpPage(w, a))
	})

	identificationPage := container.NewVBox(
		pageLabel,
		logInButton,
		signUpButton,
	)

	return identificationPage
}

// signUpPage returns a page that contains a create username and a create password field. Adds a new user in the database
func signUpPage(w fyne.Window, a fyne.App) *fyne.Container {

	pageLabel := widget.NewLabel("Signing up ...")

	usernameEntry := widget.NewEntry()
	entryUsernameHolder := "Username ..."
	usernameEntry.SetPlaceHolder(entryUsernameHolder)

	passwordEntry := widget.NewEntry()
	entryPasswordHolder := "Password ..."
	passwordEntry.SetPlaceHolder(entryPasswordHolder)

	validationButton := widget.NewButton("Confirm", func() {
		dialog.NewInformation("Success", "Your user account has been successfuly created !", w)

		mdb.SetUsernameOfSession(passwordEntry.Text)

		if appStartOption == "local" {
			db, err := mdb.LoadDB()
			if err != nil {
				panic(err)
			}
			w.SetContent(MainPage(db, w, a))
		} else if appStartOption == "browser" {
			db, err := mf.LoadDB()
			if err != nil {
				panic(err)
			}
			w.SetContent(MainPage(db, w, a))
		}

	})

	signUpPage := container.NewVBox(
		pageLabel,
		usernameEntry,
		passwordEntry,
		validationButton,
	)

	return signUpPage
}

// logInPage returns a page that contains a enter username and a enter password field. Checks if the user is in the database. If yes, sets the variable user_id for the rest of the program
func logInPage(w fyne.Window, a fyne.App) *fyne.Container {
	pageLabel := widget.NewLabel("Signing up ...")

	usernameEntry := widget.NewEntry()
	entryUsernameHolder := "Username ..."
	usernameEntry.SetPlaceHolder(entryUsernameHolder)

	passwordEntry := widget.NewEntry()
	entryPasswordHolder := "Password ..."
	passwordEntry.SetPlaceHolder(entryPasswordHolder)

	validationButton := widget.NewButton("Connect", func() {
		// go to main page
	})

	logInPage := container.NewVBox(
		pageLabel,
		usernameEntry,
		passwordEntry,
		validationButton,
	)

	return logInPage
}
