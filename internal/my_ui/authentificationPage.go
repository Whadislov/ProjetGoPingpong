package myapp

import (
	"fmt"
	"log"

	mdb "github.com/Whadislov/ProjetGoPingPong/internal/my_db"
	mfr "github.com/Whadislov/ProjetGoPingPong/internal/my_frontend"
	mf "github.com/Whadislov/ProjetGoPingPong/internal/my_functions"

	mt "github.com/Whadislov/ProjetGoPingPong/internal/my_types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// AuthentificationPage returns a page that contains a log in button and a sign up button
func AuthentificationPage(w fyne.Window, a fyne.App) *fyne.Container {

	pageLabel := widget.NewLabel("Please authenticate")
	pageLabel.Alignment = fyne.TextAlignCenter

	var db *mt.Database
	var err error
	if appStartOption == "local" {
		db, err = mdb.LoadDB()
		if err != nil {
			panic(err)
		}
	} else if appStartOption == "browser" {
		db, err = mfr.LoadDB()
		if err != nil {
			panic(err)
		}
	}
	logInButton := widget.NewButton("Log in", func() {
		w.SetContent(logInPage(db, w, a))
	})

	signUpButton := widget.NewButton("Sign up", func() {
		cancelButton := widget.NewButton("Cancel", func() {
			w.SetContent(AuthentificationPage(w, a))
		})

		content := container.NewVBox(signUpPage(db, w, a), cancelButton)
		w.SetContent(content)
	})

	if len(db.Users) > 0 {
		identificationPage := container.NewVBox(
			pageLabel,
			logInButton,
			signUpButton,
		)
		return identificationPage
	} else {

		quitButton := widget.NewButton("Quit", func() {
			a.Quit()
			log.Println("Application exited.")
		})
		content := container.NewVBox(
			signUpPage(db, w, a),
			quitButton,
		)
		// No user in the database, go directly to sign up page
		w.SetContent(content)
		return signUpPage(db, w, a)
	}
}

// signUpPage returns a page that contains a create username and a create password field. Adds a new user in the database
func signUpPage(db *mt.Database, w fyne.Window, a fyne.App) *fyne.Container {

	pageLabel := widget.NewLabel("Create your account")
	pageLabel.Alignment = fyne.TextAlignCenter

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
		if len(db.Users) == 0 {

			_, err := mf.NewUser(usernameEntry.Text, emailEntry.Text, passwordEntry.Text, db)
			if err != nil {
				switch err.Error() {
				case "username cannot be empty":
					dialog.ShowError(err, w)
					w.SetContent(signUpPage(db, w, a))
				case "username is already taken":
					dialog.ShowError(err, w)
					usernameEntry.SetPlaceHolder("")
				case "email cannot be empty":
					dialog.ShowError(err, w)
					emailEntry.SetPlaceHolder("abc@def.com")
				case "email is already used":
					dialog.ShowError(err, w)
					emailEntry.SetPlaceHolder("abc@def.com")
				case "email must be valid. Example: abc@def.com":
					dialog.ShowError(err, w)
					emailEntry.SetPlaceHolder("abc@def.com")
				default:
					dialog.ShowError(err, w)
					w.SetContent(signUpPage(db, w, a))
				}
			} else {
				dialog.ShowInformation("Success", "Your user account has been created !", w)
				mdb.SetUsernameOfSession(usernameEntry.Text)

				// Save the new user in the database
				if appStartOption == "local" {
					mdb.SaveDB(db)
				} else if appStartOption == "browser" {
					mfr.SaveDB(db)
				}

				log.Println("Sign up is successfull")
				fmt.Println("len(db.Users): ", len(db.Users))

				w.SetContent(MainPage(db, w, a))
				w.SetMainMenu(MainMenu(db, w, a))
			}
		} else {
			// Verify if username or email are already used. Verify if passwords match
			for _, user := range db.Users {
				if emailEntry.Text == user.Email {
					dialog.ShowError(fmt.Errorf("this email is already used"), w)
					// Email already used, reset the entry
					entryEmailHolder := ""
					emailEntry.SetPlaceHolder(entryEmailHolder)
				} else if usernameEntry.Text == user.Name {
					dialog.ShowError(fmt.Errorf("this username is already taken"), w)
					// Username already used, reset the entry
					entryUsernameHolder := ""
					usernameEntry.SetPlaceHolder(entryUsernameHolder)
				} else if passwordEntry.Text != confirmPasswordEntry.Text {
					dialog.ShowError(fmt.Errorf("passwords do not match"), w)
					entryPasswordHolder := ""
					passwordEntry.SetPlaceHolder(entryPasswordHolder)
					entryConfirmPasswordHolder := ""
					confirmPasswordEntry.SetPlaceHolder(entryConfirmPasswordHolder)
				} else {
					dialog.ShowInformation("Success", "Your user account has been created !", w)
					mdb.SetUsernameOfSession(usernameEntry.Text)
					log.Println("Sign up is successfull")
					mf.NewUser(usernameEntry.Text, emailEntry.Text, passwordEntry.Text, db)

					// Save the new user in the database
					if appStartOption == "local" {
						mdb.SaveDB(db)
					} else if appStartOption == "browser" {
						mfr.SaveDB(db)
					}
					fmt.Println("len(db.Users): ", len(db.Users))

					w.SetContent(MainPage(db, w, a))
					w.SetMainMenu(MainMenu(db, w, a))
				}
			}

		}

	})

	signUpPage := container.NewVBox(
		pageLabel,
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
	pageLabel := widget.NewLabel("Logging in ...")

	usernameLabel := widget.NewLabel("👤 Username")
	usernameEntry := widget.NewEntry()

	passwordLabel := widget.NewLabel("🔒 Password")
	passwordEntry := widget.NewPasswordEntry()

	validationButton := widget.NewButton("Connect", func() {
		// Verify if username and password match
		log.Println("Verifying username and password")
		fmt.Printf("Username: %s\n", usernameEntry.Text)
		fmt.Println("len(users)", len(db.Users))
		fmt.Println("len(users)", len(db.Players))
		for _, user := range db.Users {
			if usernameEntry.Text == user.Name {
				log.Println("username exists")
				if passwordEntry.Text == user.PasswordHash {
					log.Println("Log in is successfull")
					mdb.SetUsernameOfSession(usernameEntry.Text)
					w.SetContent(MainPage(db, w, a))
					w.SetMainMenu(MainMenu(db, w, a))
				}
			}
		}
		log.Println("Username and password mismatch")
		err := fmt.Errorf("username and password mismatch")
		dialog.ShowError(err, w)
		// Reset entries
		ReinitWidgetEntryText(usernameEntry, "")
		ReinitWidgetEntryText(passwordEntry, "")

	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.SetContent(AuthentificationPage(w, a))
	})

	logInPage := container.NewVBox(
		pageLabel,
		usernameLabel,
		usernameEntry,
		passwordLabel,
		passwordEntry,
		validationButton,
		cancelButton,
	)

	return logInPage
}
