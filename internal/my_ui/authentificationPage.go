package myapp

import (
	"fmt"
	"log"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// AuthentificationPage returns a page that contains a log in button and a sign up button
func AuthentificationPage(w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("TT Companion", 32)

	authLabel := widget.NewLabel("Please authenticate")
	authLabel.Alignment = fyne.TextAlignCenter

	var db *mt.Database
	var err error
	db, err = mdb.LoadUsersOnly()

	if err != nil {
		panic(err)
	}
	logInButton := widget.NewButton("Log in", func() {
		w.SetContent(logInPage(db, w, a))
	})

	signUpButton := widget.NewButton("Sign up", func() {
		cancelButtonInSignUpPage := widget.NewButton("Cancel", func() {
			w.SetContent(AuthentificationPage(w, a))
		})

		content := container.NewVBox(signUpPage(db, w, a), cancelButtonInSignUpPage)
		w.SetContent(content)
	})

	if len(db.Users) > 0 {
		authentificationPage := container.NewVBox(
			pageTitle,
			authLabel,
			logInButton,
			signUpButton,
		)
		return authentificationPage
	} else {

		quitButtonInSignUpPage := widget.NewButton("Quit", func() {
			a.Quit()
			log.Println("Application exited.")
		})
		content := container.NewVBox(
			signUpPage(db, w, a),
			quitButtonInSignUpPage,
		)
		// No user in the database, go directly to sign up page
		w.SetContent(content)
		return signUpPage(db, w, a)
	}
}

// signUpPage returns a page that contains a create username and a create password field. Adds a new user in the database
func signUpPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Create your account", 32)

	emailLabel := widget.NewLabel("✉️ Email")
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("abc@def.com")

	usernameLabel := widget.NewLabel("👤 Username")
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel("🔒 Password")
	passwordEntry := widget.NewPasswordEntry()

	confirmPasswordLabel := widget.NewLabel("Confirm Password")
	confirmPasswordEntry := widget.NewPasswordEntry()

	validationButton := widget.NewButton("Create", func() {
		log.Println("Creating new User")
		newUser, err := mf.NewUser(usernameEntry.Text, emailEntry.Text, passwordEntry.Text, confirmPasswordEntry.Text, db)
		if err != nil {
			switch err.Error() {
			case "email cannot be empty":
				log.Println("email is empty")
				dialog.ShowError(err, w)
				emailEntry.SetPlaceHolder("abc@def.com")
			case "email is already used":
				log.Println("email is already used")
				dialog.ShowError(err, w)
				emailEntry.SetPlaceHolder("abc@def.com")
			case "email must be valid. Example: abc@def.com":
				log.Println("email is not valid")
				dialog.ShowError(err, w)
				emailEntry.SetPlaceHolder("abc@def.com")
			case "username cannot be empty":
				log.Println("username is empty")
				dialog.ShowError(err, w)
				usernameEntry.SetPlaceHolder("")
			case "username must be valid (only letters and figures are allowed, spaces are not allowed)":
				log.Println("username is not valid")
				dialog.ShowError(err, w)
				usernameEntry.SetPlaceHolder("")
			case "username is already taken":
				log.Println("username is already taken")
				dialog.ShowError(err, w)
				usernameEntry.SetPlaceHolder("")
			case "password cannot be empty":
				log.Println("password cannot be empty")
				dialog.ShowError(err, w)
				passwordEntry.SetText("")
				confirmPasswordEntry.SetText("")
			case "password must be valid (spaces are not allowed)":
				log.Println("password is not valid")
				dialog.ShowError(err, w)
				passwordEntry.SetText("")
				confirmPasswordEntry.SetText("")
			case "passwords do not match":
				log.Println("passwords do not match")
				dialog.ShowError(err, w)
				passwordEntry.SetText("")
				confirmPasswordEntry.SetText("")
			default:
				log.Println("unknown error")
				dialog.ShowError(err, w)
				w.SetContent(signUpPage(db, w, a))
			}
		} else {
			var err error
			dialog.ShowInformation("Success", "Your user account has been created !", w)
			mdb.SetUserIDOfSession(newUser.ID)

			// Save the new user in the database
			mdb.SaveDB(db)
			log.Println("Sign up is successfull")

			// Now load the whole database
			db, err = mdb.LoadDB()

			if err != nil {
				dialog.ShowError(err, w)
			} else {
				userOfSession = db.Users[newUser.ID]
				w.SetContent(MainPage(db, w, a))
				w.SetMainMenu(MainMenu(db, w, a))
			}
		}
	})

	signUpPage := container.NewVBox(
		pageTitle,
		emailLabel,
		emailEntry,
		usernameLabel,
		usernameEntry,
		passwordLabel,
		passwordEntry,
		confirmPasswordLabel,
		confirmPasswordEntry,
		validationButton,
	)

	return signUpPage
}

// logInPage returns a page that contains a enter username and a enter password field. Checks if the user is in the database. If yes, sets the variable user_id for the rest of the program
func logInPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {
	pageTitle := setTitle("Login", 32)

	usernameLabel := widget.NewLabel("👤 Username")
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel("🔒 Password")
	passwordEntry := widget.NewPasswordEntry()

	validationButton := widget.NewButton("Connect", func() {
		// Verify if username and password match
		log.Println("Verifying username and password")
		for _, user := range db.Users {
			if usernameEntry.Text == user.Name {
				if passwordEntry.Text == user.PasswordHash {
					log.Println("Login is successfull")
					mdb.SetUserIDOfSession(user.ID)
					// Now load the corresponding database of the user
					var err error

					mdb.SetUserIDOfSession(user.ID)
					db, err = mdb.LoadDB()

					if err != nil {
						dialog.ShowError(err, w)
					} else {
						userOfSession = db.Users[user.ID]
						w.SetContent(MainPage(db, w, a))
						w.SetMainMenu(MainMenu(db, w, a))
						return
					}
				}
			}
		}
		// Username and password mismatch
		log.Println("Login failed : Username and password mismatch")
		dialog.ShowError(fmt.Errorf("username and password mismatch"), w)
		// Reset entries
		ReinitWidgetEntryText(usernameEntry, "")
		passwordEntry.SetText("")

	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPage(w, a))
	})

	logInPage := container.NewVBox(
		pageTitle,
		usernameLabel,
		usernameEntry,
		passwordLabel,
		passwordEntry,
		validationButton,
		cancelButton,
	)

	return logInPage
}
